use ::argon2;
use ::rand::{thread_rng, Rng};

// Hashing configuration
const CONFIG: argon2::Config = argon2::Config {
    variant: argon2::Variant::Argon2id,
    version: argon2::Version::Version13,
    mem_cost: 8192,
    time_cost: 10,
    lanes: 4,
    thread_mode: argon2::ThreadMode::Parallel,
    secret: &[],
    ad: &[],
    hash_length: 32
};

// Securely hash the provided password
pub fn hash_password(password: String) -> String {
    let salt = thread_rng().gen_ascii_chars().take(32).collect::<String>();
    argon2::hash_encoded(password.as_bytes(), salt.as_bytes(), &CONFIG).unwrap()
}

// Compare the given plaintext password to the provided hash
pub fn compare_password(password: String, hash: String) -> Option<()> {
    let result = argon2::verify_encoded(&hash, password.as_bytes()).unwrap();

    if result {
        Some(())
    } else {
        None
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn rehash_is_different() {
        let input = "password".to_owned();

        let hash1 = super::hash_password(input.clone());
        let hash2 = super::hash_password(input.clone());
        assert_ne!(hash1, hash2);
    }

    #[test]
    fn compares_pass_as_expected() {
        let input = "password".to_owned();

        let hash = super::hash_password(input.clone());
        assert_eq!(Some(()), super::compare_password(input.clone(), hash));
    }

    #[test]
    fn compares_fail_as_expected() {
        let hash = super::hash_password("password".to_owned());

        let checks = vec![
            "Password",
        ];
        for check in checks {
            assert_eq!(None, super::compare_password(check.to_owned(), hash.clone()));
        }
    }
}
