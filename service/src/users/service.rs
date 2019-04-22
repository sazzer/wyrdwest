use crate::users::user;
use failure::{Error, Fail};

#[derive(Debug, Fail)]
pub enum UserRetrieverError {
    #[fail(display = "Unknown User ID: {}", id)]
    UnknownUser {
        id: user::UserID,
    },
}

// Trait describing how to retrieve user records
pub trait UserRetriever {
    // Get the User from the data store that has the given ID
    fn get_user_by_id(&self, user_id: &user::UserID) -> Result<user::UserModel, Error>;
}
