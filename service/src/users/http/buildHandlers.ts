import { RouteOptions } from 'fastify';
import { IncomingMessage, Server, ServerResponse } from 'http';
import { UserRetriever } from '../retriever';
import { buildGetUserByIdHandler } from './getUserById';

/**
 * Build all the handlers for working with users
 * @param userRetriever The means to retrieve users from the database
 */
export function buildUserHandlers(
  userRetriever: UserRetriever
): ReadonlyArray<RouteOptions<Server, IncomingMessage, ServerResponse>> {
  return [buildGetUserByIdHandler(userRetriever)];
}
