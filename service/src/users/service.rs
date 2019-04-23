use crate::users::*;
use failure::{Fail};

#[derive(Debug, Fail, PartialEq)]
pub enum UserRetrieverError {
    #[fail(display = "Unknown User ID: {}", id)]
    UnknownUser {
        id: UserID,
    },
}

// Trait describing how to retrieve user records
pub trait UserRetriever {
    // Get the User from the data store that has the given ID
    fn get_user_by_id(&self, user_id: UserID) -> Result<UserModel, UserRetrieverError>;
}
