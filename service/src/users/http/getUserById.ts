import { FastifyReply, FastifyRequest, RouteOptions } from 'fastify';
import { IncomingMessage, Server, ServerResponse } from 'http';
import { Problem } from '../../problem';
import { UserRetriever } from '../retriever';
import { UserNotFoundError } from '../unknownUserError';
import { translateUserToResponse, UserResponseModel } from './model';
import { SINGLE_ROUTE_URI } from './urls';

/**
 * Handler for retrieving a single User by ID
 * @param userRetriever The means to retrieve users from the database
 */
export function buildGetUserByIdHandler(
  userRetriever: UserRetriever
): RouteOptions<Server, IncomingMessage, ServerResponse> {
  return {
    async handler(
      request: FastifyRequest<IncomingMessage>,
      reply: FastifyReply<ServerResponse>
    ): Promise<UserResponseModel | Problem> {
      const userId = request.params.userId;

      try {
        const user = await userRetriever.getUserById(userId);
        reply.status(200);
        return translateUserToResponse(user);
      } catch (e) {
        if (e instanceof UserNotFoundError) {
          reply.status(404);
          return {
            type: 'tag:wyrdwest,2019:users/problems/unknown-user',
            title: 'The requested user could not be found',
            status: 404
          };
        } else {
          reply.status(500);
          return {
            type: 'tag:wyrdwest,2019:problems/internal-server-error',
            title: 'An unexpected error occurred',
            status: 500
          };
        }
      }
    },
    method: 'GET',
    url: SINGLE_ROUTE_URI
  };
}
