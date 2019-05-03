/** The reason the access token was invalid */
export enum InvalidAccessTokenReason {
  MALFORMED_JWT = 'Malformed JWT',
  EXPIRY_IN_PAST = 'Expiry is in the past',
  CREATED_IN_FUTURE = 'Created date is in the future'
}

/**
 * Error indicating that an Access Token is invalid
 */
export class InvalidAccessTokenError extends Error {
  constructor(public readonly reason: InvalidAccessTokenReason) {
    super(`Invalid Access Token: ${reason}`);
  }
}
