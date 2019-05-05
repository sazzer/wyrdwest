import { RouteDefinition } from '../../server/routes';
import { GrantTypes } from '../clients/model';
import { buildAuthorizationCodeTokenHandler } from './authorizationCode';
import { buildClientCredentialsTokenHandler } from './clientCredentials';
import { buildRefreshTokenHandler } from './refreshToken';
import { buildTokenHandler } from './tokenHandler';

/**
 * Build all the handlers for working with OAuth2
 */
export function buildOAuth2Handlers(): ReadonlyArray<RouteDefinition> {
  return [
    buildTokenHandler({
      [GrantTypes.AUTHORIZATION_CODE.toString()]: buildAuthorizationCodeTokenHandler(),
      [GrantTypes.CLIENT_CREDENTIALS.toString()]: buildClientCredentialsTokenHandler(),
      [GrantTypes.REFRESH_TOKEN.toString()]: buildRefreshTokenHandler()
    })
  ];
}