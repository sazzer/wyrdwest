import { UserID, UserModel } from './model';

/**
 * Interface describing how to load users from the data store
 */
export interface UserRetriever {
  /**
   * Get a User with the provided ID
   * @param id The ID of the user
   * @return The User itself
   */
  readonly getUserById: (id: UserID) => Promise<UserModel>;
}
