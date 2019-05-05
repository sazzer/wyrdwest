import { TokenHandlerFunction } from './handlers';
import { AccessTokenErrorCode, AccessTokenErrorModel, AccessTokenModel } from './model';

/**
 * Build the Token Handler to use for an Refresh Token grant
 */
export function buildRefreshTokenHandler(): TokenHandlerFunction {
  return async function handleRefreshToken(_: { readonly [key: string]: string }): Promise<AccessTokenModel> {
    throw new AccessTokenErrorModel(AccessTokenErrorCode.UNSUPPORTED_GRANT_TYPE, 'Not yet implemented');
  };
}
