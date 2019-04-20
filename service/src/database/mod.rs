use r2d2_postgres::PostgresConnectionManager;
use r2d2::Pool;

mod health;
pub mod migrations;

// Wrapper around a Database Connection Pool to make it easier to work with
pub trait Database {

}

// Standard implementation of the Database Wrapper
pub struct DatabaseWrapper {
    url: String,
    pool: Pool<PostgresConnectionManager<postgres::NoTls>>
}

impl DatabaseWrapper {
    // Create a new Database Wrapper
    pub fn new(url: String) -> DatabaseWrapper {
        let manager = PostgresConnectionManager::new(
            url.parse().unwrap(),
            postgres::NoTls,
        );
        let pool = r2d2::Pool::new(manager).unwrap();

        DatabaseWrapper {
            url,
            pool
        }
    }
}

impl Database for DatabaseWrapper {

}

