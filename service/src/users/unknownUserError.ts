import { UserID } from './model';

/**
 * Error indicating that a requested user could not be found
 */
export class UserNotFoundError extends Error {
  constructor(public readonly userId: UserID) {
    super(`User not found: ${userId}`);
  }
}
