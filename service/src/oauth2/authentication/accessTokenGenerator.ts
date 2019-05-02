import { Clock, Duration } from 'js-joda';
import { v4 as uuid } from 'uuid';
import { UserID } from '../../users/model';
import { ClientID } from '../clients/model';
import { AccessToken } from './accessToken';

/**
 * Means to generate an Access Token
 */
export class AccessTokenGenerator {
  constructor(private clock: Clock, private expiry: Duration) {}
  /**
   * Generate an access token
   * @param user The user to generate the access token for
   * @param client The client to generate the access token for
   * @param scopes The scopes to include in the access token
   */
  public generate(user: UserID, client: ClientID, scopes?: readonly string[]): AccessToken {
    const now = this.clock.instant();
    const expires = now.plus(this.expiry);

    return {
      id: uuid(),
      created: now,
      expires,
      client,
      user,
      scopes
    };
  }
}
