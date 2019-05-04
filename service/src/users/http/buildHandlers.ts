import { RouteDefinition } from '../../server/routes';
import { UserRetriever } from '../retriever';
import { buildGetUserByIdHandler } from './getUserById';

/**
 * Build all the handlers for working with users
 * @param userRetriever The means to retrieve users from the database
 */
export function buildUserHandlers(userRetriever: UserRetriever): ReadonlyArray<RouteDefinition> {
  return [buildGetUserByIdHandler(userRetriever)];
}
