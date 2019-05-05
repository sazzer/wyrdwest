/** Representation of an Access Token as returned over the API */
export interface AccessTokenModel {
  readonly access_token: string;
  readonly token_type: string;
  readonly expires_in: number;
  readonly refresh_token?: string;
  readonly scope?: string;
}

/** Possible error codes for an access token error */
export enum AccessTokenErrorCode {
  INVALID_REQUEST = 'invalid_request',
  INVALID_CLIENT = 'invalid_client',
  INVALID_GRANT = 'invalid_grant',
  UNAUTHORIZED_CLIENT = 'unauthorized_client',
  UNSUPPORTED_GRANT_TYPE = 'unsupported_grant_type',
  INVALID_SCOPE = 'invalid_scope',
  ACCESS_DENIED = 'access_denied',
  UNSUPPORTED_RESPONSE_TYPE = 'unsupported_response_type',
  SERVER_ERROR = 'server_error',
  TEMPORARILY_UNAVAILABLE = 'temporarily_unavailable'
}

/** Representation of an access token error */
export class AccessTokenErrorModel {
  // tslint:disable:variable-name
  constructor(
    public readonly error: AccessTokenErrorCode,
    public readonly error_description?: string,
    public readonly error_uri?: string
  ) {}
}
