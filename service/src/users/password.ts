import argon2 from 'argon2';

/**
 * Hash the given plaintext password
 * @param password the password
 * @return the hash
 */
export function hashPassword(password: string): Promise<string> {
  return argon2.hash(password);
}

/**
 * Verify that the given hash matches the given plaintext
 * @param hash the hash
 * @param password the plaintext
 * @return True if the two match. False if they differ
 */
export function checkPassword(hash: string, password: string): Promise<boolean> {
  return argon2.verify(hash, password);
}
