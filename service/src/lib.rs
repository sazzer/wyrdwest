pub mod config;
mod health;
mod database;

#[macro_use]
extern crate log;

use crate::health::healthchecks::Healthcheck;
use actix_web::{middleware, server};
use std::collections::HashMap;
use std::sync::Arc;
use r2d2_postgres::PostgresConnectionManager;

// Open the connection pool tot he database
fn connect_to_database(url: String) -> Arc<database::DatabaseWrapper> {
    info!("Connecting to database: {}", url);
    let manager = PostgresConnectionManager::new(
        url.parse().unwrap(),
        postgres::NoTls,
    );
    let pool = r2d2::Pool::new(manager).unwrap();

    info!("Connected to database");

    Arc::new(database::DatabaseWrapper::new(pool))
}

// Actually start the application
pub fn start(settings: config::Config) {
    let database = connect_to_database(settings.db_uri);

    let mut healthchecks: HashMap<String, Arc<Healthcheck>> = HashMap::new();
    healthchecks.insert("database".to_string(), database);

    let server = server::new(move || {
        vec![health::http::new(healthchecks.clone()).middleware(middleware::Logger::default())]
    });

    let port = format!("[::]:{}", settings.port);

    server.bind(port).unwrap().workers(20).run();
}
