use std::marker::{Send, Sync};

// Representation of some means to check the health of the system
pub trait Healthcheck: Sync + Send {
    // Actually perform the healthcheck, returning either a Success or a Failure message
    fn check(&self) -> Result<String, String>;
}
