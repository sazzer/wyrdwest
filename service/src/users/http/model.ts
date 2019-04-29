import uriTemplate from 'url-template';
import { UserModel } from '../model';
import { SINGLE_TEMPLATE } from './urls';

/**
 * Shape of the HTTP Model for an incoming User
 */
export interface UserRequestModel {
  readonly name: string;
  readonly email?: string;
}

/**
 * Shape of the HTTP Model for an outgoing User
 */
export type UserResponseModel = UserRequestModel & {
  readonly self: string;
};

/**
 * Translate an internal User Model into an HTTP Response
 * @param user the user to translate
 */
export function translateUserToResponse(user: UserModel): UserResponseModel {
  return {
    self: uriTemplate.parse(SINGLE_TEMPLATE).expand({ userId: user.identity.id }),
    name: user.data.name,
    email: user.data.email
  };
}
