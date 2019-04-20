pub mod config;
mod health;
mod database;

#[macro_use]
extern crate log;

use crate::health::healthchecks::Healthcheck;
use crate::database::migrations::MigratableDatabase;
use actix_web::{middleware, server};
use std::collections::HashMap;
use std::sync::Arc;

// Open the connection pool tot he database
fn connect_to_database(url: String) -> Arc<database::DatabaseWrapper> {
    info!("Connecting to database: {}", url);

    Arc::new(database::DatabaseWrapper::new(url))
}

// Actually start the application
pub fn start(settings: config::Config) {
    let database = connect_to_database(settings.db_uri);
    database.migrate().unwrap();

    let mut healthchecks: HashMap<String, Arc<Healthcheck>> = HashMap::new();
    healthchecks.insert("database".to_string(), database);

    let server = server::new(move || {
        vec![health::http::new(healthchecks.clone()).middleware(middleware::Logger::default())]
    });

    let port = format!("[::]:{}", settings.port);

    server.bind(port).unwrap().workers(20).run();
}
