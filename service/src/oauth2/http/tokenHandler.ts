import { FastifyReply, FastifyRequest, RouteOptions } from 'fastify';
import { IncomingMessage, Server, ServerResponse } from 'http';
import { AccessTokenErrorCode, AccessTokenErrorModel, AccessTokenModel } from './model';

export type TokenHandler = (
  req: FastifyRequest<IncomingMessage>,
  res: FastifyReply<ServerResponse>
) => Promise<AccessTokenModel | AccessTokenErrorModel>;

/**
 * Build the Token Handler to use
 */
export function buildTokenHandler(handlers: {
  readonly [grantType: string]: TokenHandler;
}): RouteOptions<Server, IncomingMessage, ServerResponse> {
  return {
    async handler(
      request: FastifyRequest<IncomingMessage>,
      response: FastifyReply<ServerResponse>
    ): Promise<AccessTokenModel | AccessTokenErrorModel> {
      const grantType = (request.body || {}).grant_type;

      if (grantType === undefined || grantType === '') {
        return {
          error: AccessTokenErrorCode.INVALID_REQUEST,
          error_description: 'No Grant Type was specified'
        };
      }

      const handler = handlers[grantType];

      if (handler) {
        return handler(request, response);
      } else {
        return {
          error: AccessTokenErrorCode.UNSUPPORTED_GRANT_TYPE,
          error_description: `The requested grant type is not supported: ${grantType}`
        };
      }
    },
    method: 'POST',
    url: '/oauth2/token'
  };
}
