import { Clock, Instant } from 'js-joda';
import jwt from 'jsonwebtoken';
import { AccessToken } from './accessToken';
import { InvalidAccessTokenError, InvalidAccessTokenReason } from './invalidAccessTokenError';

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
  constructor(private readonly clock: Clock, private readonly key: string, private readonly algorithm: string) {}

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
            reject(new InvalidAccessTokenError(InvalidAccessTokenReason.MALFORMED_JWT));
          } else if (typeof parsed === 'string') {
            reject(new InvalidAccessTokenError(InvalidAccessTokenReason.MALFORMED_JWT));
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

            const now = this.clock.instant();
            if (now.isAfter(accessToken.expires)) {
              // The token has expired
              reject(new InvalidAccessTokenError(InvalidAccessTokenReason.EXPIRY_IN_PAST));
            } else if (now.isBefore(accessToken.created)) {
              // The token hasn't been created yet
              reject(new InvalidAccessTokenError(InvalidAccessTokenReason.CREATED_IN_FUTURE));
            }

            resolve(accessToken);
          }
        }
      );
    });
  }
}
