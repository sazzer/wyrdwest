import { Instant } from 'js-joda';
import jwt from 'jsonwebtoken';
import { AccessToken } from './accessToken';

interface DecodedJWT {
  readonly jti: string;
  readonly exp: number;
  readonly nbf: number;
  readonly iat: number;
  readonly aud: string;
  readonly iss: string;
  readonly sub: string;
  readonly scopes?: readonly string[];
}

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

  /**
   * Deserialize an access token from a string
   * @param accessToken The access token to deserialize
   */
  public deserialize(token: string): Promise<AccessToken> {
    return new Promise((resolve, reject) => {
      jwt.verify(
        token,
        this.key,
        {
          ignoreExpiration: true,
          ignoreNotBefore: true
        },
        (err, parsed) => {
          if (err) {
            reject(err);
          } else if (typeof parsed === 'string') {
            reject(err);
          } else {
            const decodedJwt = parsed as DecodedJWT;
            const accessToken: AccessToken = {
              id: decodedJwt.jti,
              client: decodedJwt.iss,
              user: decodedJwt.sub,
              created: Instant.ofEpochSecond(decodedJwt.iat),
              expires: Instant.ofEpochSecond(decodedJwt.exp),
              scopes: decodedJwt.scopes
            };
            resolve(accessToken);
          }
        }
      );
    });
  }
}
