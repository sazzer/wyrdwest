#[macro_use]
extern crate log;
extern crate log4rs;

fn main() {
    log4rs::init_file("log4rs.yml", Default::default()).unwrap();

    info!("Hello, world!");
}
