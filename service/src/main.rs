extern crate wyrdwest_service;

use log4rs;

fn main() {
    let config = ::wyrdwest_service::config::Config::new();

    // Load the logging configuration to use
    log4rs::init_file("log4rs.yml", Default::default()).unwrap();

    ::wyrdwest_service::start(config)
}
