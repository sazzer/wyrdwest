import { Database } from '../../database';
import { NoRowsReturnedError } from '../../database/noRowsReturnedError';
import { UserID, UserModel } from '../model';
import { UserNotFoundError } from '../unknownUserError';
import { parseDatabaseRow } from './model';

/**
 * Get a User with the provided ID
 * @param database The database connection to use
 * @param id The ID of the user
 * @return The User itself
 */
export async function getUserById(database: Database, id: UserID): Promise<UserModel> {
  try {
    const userRow = await database.queryOne('SELECT * FROM users WHERE user_id = $1', id);
    return parseDatabaseRow(userRow);
  } catch (e) {
    if (e instanceof NoRowsReturnedError) {
      throw new UserNotFoundError(id);
    }
    throw e;
  }
}
