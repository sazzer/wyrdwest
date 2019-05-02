import { Instant } from 'js-joda';
import { UserID } from '../../users/model';
import { ClientID } from '../clients/model';

/**
 * Type representing the ID of an Access Token
 */
export type AccessTokenID = string;

/**
 * Type representing an Access Token
 */
export interface AccessToken {
  readonly id: AccessTokenID;
  readonly created: Instant;
  readonly expires: Instant;
  readonly client: ClientID;
  readonly user: UserID;
  readonly scopes?: readonly string[];
}
