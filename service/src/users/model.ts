import { Model } from '../service';

/**
 * Type representing the data in a user
 */
export interface UserData {
  readonly name: string;
  readonly password?: string;
  readonly email?: string;
}

/**
 * Type representing the ID of a user
 */
export type UserID = string;

/**
 * Type representing the actual user resource model
 */
export type UserModel = Model<UserID, UserData>;
