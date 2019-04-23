// Representation of the ID of a user
pub type UserID = String;

// Representation of the data that makes up a user
#[derive(Debug)]
pub struct UserData {
    name: String,
    email: String,
    password: String,
}

// Representation of a Model representing a User
pub type UserModel = crate::model::Model<UserID, UserData>;
