#[macro_use]
extern crate log;
extern crate actix_web;

use std::collections::HashMap;
use actix_web::{middleware, web, App, HttpRequest, HttpServer};

fn index(req: HttpRequest) -> &'static str {
    "Hello world!"
}

pub fn start(settings: HashMap<String, String>) {
    info!("Hello, world!");
    info!("{:?}", settings);

    let server = HttpServer::new(|| {
        App::new()
            // enable logger
            .wrap(middleware::Logger::default())
            .wrap(middleware::cors::Cors::new())
            .service(web::resource("/").to(index))
    });

    let port = settings.get("port")
        .map(|port| format!("[::]:{}", port))
        .unwrap();

    server
        .bind(port).unwrap()
        .workers(20)
        .run().unwrap();
}
