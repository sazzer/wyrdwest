#[macro_use]
extern crate log;

use std::collections::HashMap;

pub fn start(settings: HashMap<String, String>) {
    info!("Hello, world!");
    info!("{:?}", settings);
}
