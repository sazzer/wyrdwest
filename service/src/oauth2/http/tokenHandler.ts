import { Request, Response } from 'express';
import { Method, RouteDefinition } from '../../server/routes';
import { TokenHandlerFunction } from './handlers';
import { AccessTokenErrorCode, AccessTokenErrorModel } from './model';

/**
 * Build the Token Handler to use
 */
export function buildTokenHandler(handlers: { readonly [grantType: string]: TokenHandlerFunction }): RouteDefinition {
  return {
    handler: async (req: Request, res: Response) => {
      const grantType = (req.body || {}).grant_type;

      try {
        if (grantType === undefined || grantType === '') {
          throw new AccessTokenErrorModel(AccessTokenErrorCode.INVALID_REQUEST, 'No Grant Type was specified');
        }

        const handler = handlers[grantType];
        if (!handler) {
          throw new AccessTokenErrorModel(
            AccessTokenErrorCode.UNSUPPORTED_GRANT_TYPE,
            `The requested grant type is not supported: ${grantType}`
          );
        }

        const token = await handler(req.body);

        res.json(token);
      } catch (e) {
        if (e instanceof AccessTokenErrorModel) {
          res.status(400);
          res.json(e);
        } else {
          res.status(500);
          res.json({
            error: AccessTokenErrorCode.SERVER_ERROR,
            error_description: 'An unexpected error occurred'
          });
        }
      }
    },
    method: Method.POST,
    url: '/oauth2/token'
  };
}
