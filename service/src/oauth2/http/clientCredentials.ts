import { TokenHandlerFunction } from './handlers';
import { AccessTokenErrorCode, AccessTokenErrorModel, AccessTokenModel } from './model';

/**
 * Build the Token Handler to use for an Client Credentials grant
 */
export function buildClientCredentialsTokenHandler(): TokenHandlerFunction {
  return async function handleClientCredentialsToken(_: { readonly [key: string]: string }): Promise<AccessTokenModel> {
    throw new AccessTokenErrorModel(AccessTokenErrorCode.UNSUPPORTED_GRANT_TYPE, 'Not yet implemented');
  };
}
