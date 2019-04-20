use crate::database::DatabaseWrapper;
use crate::health::healthchecks::Healthcheck;

impl Healthcheck for DatabaseWrapper {
    // Actually perform the healthcheck, returning either a Success or a Failure message
    fn check(&self) -> Result<String, String> {
        match self.pool.get() {
            Err(e) => Err(format!("Connection Error: {}", e)),
            Ok(mut client) => {
                match client.query("SELECT 1", &[]) {
                    Err(e) => Err(format!("Query Error: {}", e)),
                    Ok(_) => Ok("Database OK".to_string()),
                }
            },
        }
    }
}
