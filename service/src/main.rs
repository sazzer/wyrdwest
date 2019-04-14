extern crate wyrdwest_service;
extern crate log4rs;

fn main() {
    log4rs::init_file("log4rs.yml", Default::default()).unwrap();

    wyrdwest_service::start()
}
