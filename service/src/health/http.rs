use std::sync::Arc;
use std::collections::HashMap;
use actix_web::{HttpRequest, HttpResponse, App, dev::Handler, http::StatusCode};
use serde::Serialize;
use crate::health::healthchecks::Healthcheck;

// The result of a healthcheck was a success
const RESULT_SUCCESS: &str = "OK";
// The result of a healthcheck was a failure
const RESULT_FAIL: &str = "FAIL";

// The payload representing a single healthcheck component
#[derive(Serialize)]
struct ComponentResult {
    // The result of this component
    result: &'static str,
    // The detail message for this component
    detail: String,
}

// The payload representing the entire system healthcheck
#[derive(Serialize)]
struct HealthcheckResult {
    // The overall result of the healthcheck
    result: &'static str,
    // The details of each component
    components: HashMap<String, ComponentResult>,
}

// Actix Handler for processing healthchecks
struct HealthcheckHandler {
    // The map of handlers that we are working with
    handlers: HashMap<String, Arc<Healthcheck>>,
}

// Implement the Actix Handler trait for performing Healthchecks
impl<S> Handler<S> for HealthcheckHandler {
    type Result = HttpResponse;

    // Handle the request to perform healthchecks on the system
    fn handle(&self, _req: &HttpRequest<S>) -> Self::Result {
        debug!("Checking system health");
        let mut result = HealthcheckResult {
            result: RESULT_SUCCESS,
            components: HashMap::new(),
        };

        let mut status = StatusCode::OK;

        for (key, value) in &self.handlers {
            let healthcheck_result = match value.check() {
                Ok(detail) => ComponentResult { result: RESULT_SUCCESS, detail: detail },
                Err(detail) => {
                    result.result = RESULT_FAIL;
                    status = StatusCode::SERVICE_UNAVAILABLE;

                    ComponentResult { result: RESULT_FAIL, detail: detail }
                },
            };

            result.components.insert(key.to_string(), healthcheck_result);
        }

        info!("Overall system health: {}", result.result);

        HttpResponse::build(status).json(result)
    }
}

// Build the Actix App that will process healthchecks
pub fn new(handlers: HashMap<String, Arc<Healthcheck>>) -> App<()> {
    let handler = HealthcheckHandler {
        handlers: handlers,
    };

    App::new()
        .prefix("/health")
        .resource("", |r| {
            r.get().h(handler);
        })
}

#[cfg(test)]
mod tests {
    use std::collections::HashMap;
    use std::sync::Arc;
    use std::str;
    use actix_web::{http, test, HttpMessage};
    use serde_json;
    use crate::health::healthchecks::Healthcheck;


    struct PassingHealthcheck {}

    impl Healthcheck for PassingHealthcheck {
        fn check(&self) -> Result<String, String> {
            Ok("Test Passed".to_string())
        }
    }

    struct FailingHealthcheck {}

    impl Healthcheck for FailingHealthcheck {
        fn check(&self) -> Result<String, String> {
            Err("Test Failed".to_string())
        }
    }

    // Actually run a test of the healthchecks handler
    fn run_test(healthchecks: HashMap<String, Arc<Healthcheck>>, status_code: http::StatusCode, expected_response: serde_json::Value) {
        let mut server = test::TestServer::with_factory(move || super::new(healthchecks.clone()));

        let request = server.client(http::Method::GET, "/health")
            .finish().unwrap();
        let response = server.execute(request.send()).unwrap();

        assert_eq!(status_code, response.status());
        assert_eq!("application/json", response.content_type());

        let bytes = server.execute(response.body()).unwrap();
        let body = serde_json::from_str(str::from_utf8(&bytes).unwrap()).unwrap();
        assert_json_eq!(body, expected_response);
    }

    #[test]
    fn test_check_no_components() {
        let healthchecks: HashMap<String, Arc<Healthcheck>> = HashMap::new();

        run_test(healthchecks, http::StatusCode::OK, json!({
            "result": "OK",
            "components": {}
        }));
    }

    #[test]
    fn test_check_passing_components() {
        let mut healthchecks: HashMap<String, Arc<Healthcheck>> = HashMap::new();
        healthchecks.insert("passing".to_string(), Arc::new(PassingHealthcheck{}));

        run_test(healthchecks, http::StatusCode::OK, json!({
            "result": "OK",
            "components": {
                "passing": {
                    "result": "OK",
                    "detail": "Test Passed"
                }
            }
        }));
    }

    #[test]
    fn test_check_failing_components() {
        let mut healthchecks: HashMap<String, Arc<Healthcheck>> = HashMap::new();
        healthchecks.insert("failing".to_string(), Arc::new(FailingHealthcheck{}));

        run_test(healthchecks, http::StatusCode::SERVICE_UNAVAILABLE, json!({
            "result": "FAIL",
            "components": {
                "failing": {
                    "result": "FAIL",
                    "detail": "Test Failed"
                }
            }
        }));
    }

    #[test]
    fn test_check_mixed_components() {
        let mut healthchecks: HashMap<String, Arc<Healthcheck>> = HashMap::new();
        healthchecks.insert("failing".to_string(), Arc::new(FailingHealthcheck{}));
        healthchecks.insert("passing".to_string(), Arc::new(PassingHealthcheck{}));

        run_test(healthchecks, http::StatusCode::SERVICE_UNAVAILABLE, json!({
            "result": "FAIL",
            "components": {
                "failing": {
                    "result": "FAIL",
                    "detail": "Test Failed"
                },
                "passing": {
                    "result": "OK",
                    "detail": "Test Passed"
                }
            }
        }));
    }
}
