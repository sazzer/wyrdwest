use crate::database::Database;

// Test implementation of the Database Wrapper
pub struct TestDatabase {
}

impl TestDatabase {
    // Create a new Test Database
    pub fn new() -> TestDatabase {
        TestDatabase {}
    }
}

impl Database for TestDatabase {

}

