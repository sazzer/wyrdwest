import { Database } from '../../database';
import { NoRowsReturnedError } from '../../database/noRowsReturnedError';
import { UserID, UserModel } from '../model';
import { UserNotFoundError } from '../unknownUserError';

/**
 * Get a User with the provided ID
 * @param database The database connection to use
 * @param id The ID of the user
 * @return The User itself
 */
export async function getUserById(database: Database, id: UserID): Promise<UserModel> {
  try {
    await database.queryOne('SELECT * FROM users WHERE user_id = $1', id);
  } catch (e) {
    if (e instanceof NoRowsReturnedError) {
      throw new UserNotFoundError(id);
    }
    throw e;
  }
  throw new UserNotFoundError(id);
}
