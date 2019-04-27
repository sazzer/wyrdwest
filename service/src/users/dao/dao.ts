import { Database } from '../../database';
import { UserID } from '../model';
import { UserRetriever } from '../retriever';
import { getUserById } from './getUserById';

/** Type representing the DAO for working with Users */
export type UserDao = UserRetriever;

/**
 * Build the DAO for working with Users
 * @param database The database connection to use
 */
export function buildUsersDao(database: Database): UserDao {
  return {
    getUserById: (id: UserID) => getUserById(database, id)
  };
}
