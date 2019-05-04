import { NextFunction, Request, RequestHandler, Response } from 'express';
import { Method, RouteDefinition } from '../../server/routes';
import { AccessTokenErrorCode } from './model';

/**
 * Build the Token Handler to use
 */
export function buildTokenHandler(handlers: { readonly [grantType: string]: RequestHandler }): RouteDefinition {
  return {
    handler: async (req: Request, res: Response, next: NextFunction) => {
      const grantType = (req.body || {}).grant_type;

      if (grantType === undefined || grantType === '') {
        res.status(400);
        res.json({
          error: AccessTokenErrorCode.INVALID_REQUEST,
          error_description: 'No Grant Type was specified'
        });
        return;
      }

      const handler = handlers[grantType];

      if (handler) {
        handler(req, res, next);
      } else {
        res.status(400);
        res.json({
          error: AccessTokenErrorCode.UNSUPPORTED_GRANT_TYPE,
          error_description: `The requested grant type is not supported: ${grantType}`
        });
      }
    },
    method: Method.POST,
    url: '/oauth2/token'
  };
}
