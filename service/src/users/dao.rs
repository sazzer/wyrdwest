use crate::users::*;
use crate::database::Database;
use std::sync::Arc;

// DAO for interacting with User records
pub struct UserDao {
    db: Arc<Database>,
}

impl UserDao {
    // Create a new User DAO
    pub fn new(db: Arc<Database>) -> UserDao {
        UserDao { db }
    }
}
impl UserRetriever for UserDao {
    // Get the User from the data store that has the given ID
    fn get_user_by_id(&self, user_id: UserID) -> Result<UserModel, UserRetrieverError> {
        Err(UserRetrieverError::UnknownUser { id: user_id })
    }
}

#[cfg(test)]
mod tests {
    use crate::users::*;
    use crate::database::test_database::TestDatabase;
    use std::sync::Arc;
    use super::UserDao;

    #[test]
    fn get_unknown_user_by_id() {
        let db = TestDatabase::new();
        let dao = UserDao::new(Arc::new(db));

        let user_id: UserID = "unknown".to_owned();
        let user = dao.get_user_by_id(user_id.clone());

        assert!(user.is_err(), "An error should have been returned");
        assert_eq!(UserRetrieverError::UnknownUser{id: user_id.clone()}, user.unwrap_err());
    }
}
