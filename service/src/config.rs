use serde::Deserialize;

// Define the default port number if not provided
fn default_port() -> u16 {
    3000
}

// The configuration for the app
#[derive(Deserialize, Debug)]
pub struct Config {
    pub db_uri: String,

    #[serde(default="default_port")]
    pub port: u16,
}

impl Config {
    // Construct a new configuration, loading from environment variables
    pub fn new() -> Config {
        envy::from_env::<Config>().unwrap()
    }
}
