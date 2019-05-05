import { TokenHandlerFunction } from './handlers';
import { AccessTokenErrorCode, AccessTokenErrorModel, AccessTokenModel } from './model';

/**
 * Build the Token Handler to use for an Authorization Code grant
 */
export function buildAuthorizationCodeTokenHandler(): TokenHandlerFunction {
  return async function handleAuthorizationCodeToken(_: { readonly [key: string]: string }): Promise<AccessTokenModel> {
    throw new AccessTokenErrorModel(AccessTokenErrorCode.UNSUPPORTED_GRANT_TYPE, 'Not yet implemented');
  };
}
