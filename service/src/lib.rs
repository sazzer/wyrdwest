mod health;

#[macro_use]
extern crate log;

use crate::health::healthchecks::Healthcheck;
use actix_web::{middleware, server};
use std::collections::HashMap;
use std::sync::Arc;
use r2d2_postgres::PostgresConnectionManager;

struct PassingHealthcheck {}

impl Healthcheck for PassingHealthcheck {
    fn check(&self) -> Result<String, String> {
        Ok("It Failed".to_string())
    }
}

fn connect_to_database(url: &str) -> r2d2::Pool<PostgresConnectionManager<postgres::NoTls>> {
    let manager = PostgresConnectionManager::new(
        url.parse().unwrap(),
        postgres::NoTls,
    );
    let pool = r2d2::Pool::new(manager).unwrap();

    info!("Connected to database");

    pool
}

// Actually start the application
pub fn start(settings: HashMap<String, String>) {
    let pool = connect_to_database(settings.get("db_uri").unwrap());

    let mut healthchecks: HashMap<String, Arc<Healthcheck>> = HashMap::new();
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
