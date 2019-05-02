import jwt from 'jsonwebtoken';
import { AccessToken } from './accessToken';

/**
 * Mechanism by which Access Tokens can be serialized as strings
 */
export class AccessTokenSerializer {
  constructor(private readonly key: string, private readonly algorithm: string) {}

  /**
   * Serialize an access token to a string
   * @param accessToken The access token to serialize
   */
  public serialize(accessToken: AccessToken): Promise<string> {
    return new Promise((resolve, reject) => {
      const payload = {
        exp: accessToken.expires.epochSecond(),
        nbf: accessToken.created.epochSecond(),
        iat: accessToken.created.epochSecond(),
        aud: accessToken.client,
        iss: accessToken.client,
        sub: accessToken.user,
        scopes: accessToken.scopes
      };

      jwt.sign(
        payload,
        this.key,
        {
          algorithm: this.algorithm,
          jwtid: accessToken.id
        },
        (err, token) => {
          if (err) {
            reject(err);
          } else {
            resolve(token);
          }
        }
      );
    });
  }
}
