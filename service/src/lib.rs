mod health;

#[macro_use]
extern crate log;
#[macro_use]
extern crate assert_json_diff;
#[macro_use]
extern crate serde_json;

use crate::health::healthchecks::Healthcheck;
use actix_web::{middleware, server};
use std::collections::HashMap;
use std::sync::Arc;

struct FailingHealthcheck {}

impl Healthcheck for FailingHealthcheck {
    fn check(&self) -> Result<String, String> {
        Err("It Failed".to_string())
    }
}

struct PassingHealthcheck {}

impl Healthcheck for PassingHealthcheck {
    fn check(&self) -> Result<String, String> {
        Ok("It Failed".to_string())
    }
}

// Actually start the application
pub fn start(settings: HashMap<String, String>) {
    let mut healthchecks: HashMap<String, Arc<Healthcheck>> = HashMap::new();
    healthchecks.insert("failing".to_string(), Arc::new(FailingHealthcheck {}));
    healthchecks.insert("passing".to_string(), Arc::new(PassingHealthcheck {}));

    let server = server::new(move || {
        vec![health::http::new(healthchecks.clone()).middleware(middleware::Logger::default())]
    });

    let port = settings
        .get("port")
        .map(|port| format!("[::]:{}", port))
        .unwrap();

    server.bind(port).unwrap().workers(20).run();
}
