extern crate wyrdwest_service;

use log4rs;
use dotenv;

fn main() {
    dotenv::dotenv().ok();

    let config = ::wyrdwest_service::config::Config::new();

    // Load the logging configuration to use
    log4rs::init_file("log4rs.yml", Default::default()).unwrap();

    ::wyrdwest_service::start(config)
}
