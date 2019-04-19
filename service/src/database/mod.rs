use r2d2_postgres::PostgresConnectionManager;
use r2d2::Pool;
use crate::health::healthchecks::Healthcheck;

// Wrapper around a Database Connection Pool to make it easier to work with
pub trait Database {

}

// Standard implementation of the Database Wrapper
pub struct DatabaseWrapper {
    pool: Pool<PostgresConnectionManager<postgres::NoTls>>
}

impl DatabaseWrapper {
    // Create a new Database Wrapper
    pub fn new(pool: Pool<PostgresConnectionManager<postgres::NoTls>>) -> DatabaseWrapper {
        DatabaseWrapper {
            pool
        }
    }
}

impl Database for DatabaseWrapper {

}

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
