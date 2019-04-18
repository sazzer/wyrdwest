extern crate wyrdwest_service;
use ::wyrdwest_service::start;

use std::collections::HashMap;
use std::env;
use config::Config;
use log4rs;

fn main() {
    let mut settings = Config::default();

    // Default port value is either the environment variable "PORT" or the value "3000", as appropriate
    let port = env::var("PORT").unwrap_or("3000".to_string());
    settings.set_default("port", port).unwrap();

    // Then merge in all of the environment variables that are prefixed with WYRDWEST
    settings
        .merge(config::Environment::with_prefix("WYRDWEST"))
        .unwrap();

    // Load the logging configuration to use
    log4rs::init_file("log4rs.yml", Default::default()).unwrap();

    start(settings.try_into::<HashMap<String, String>>().unwrap())
}
