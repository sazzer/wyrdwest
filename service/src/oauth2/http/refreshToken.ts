import { FastifyReply, FastifyRequest } from 'fastify';
import { IncomingMessage, ServerResponse } from 'http';
import { AccessTokenErrorCode, AccessTokenErrorModel, AccessTokenModel } from './model';

/**
 * Build the Token Handler to use for a Refresh Token grant
 */
export function buildRefreshTokenHandler(): (
  req: FastifyRequest<IncomingMessage>,
  res: FastifyReply<ServerResponse>
) => Promise<AccessTokenModel | AccessTokenErrorModel> {
  return async function handler(
    _1: FastifyRequest<IncomingMessage>,
    _2: FastifyReply<ServerResponse>
  ): Promise<AccessTokenModel | AccessTokenErrorModel> {
    return {
      error: AccessTokenErrorCode.INVALID_REQUEST
    };
  };
}
