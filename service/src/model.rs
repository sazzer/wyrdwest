use std::time::Instant;

// Representation of the identity of some resource
pub struct Identity<T> {
    id: T,
    version: String,
    created: Instant,
    updated: Instant,
}

// Representation of some model object
pub struct Model<ID, DATA> {
    identity: Identity<ID>,
    data: DATA,
}
