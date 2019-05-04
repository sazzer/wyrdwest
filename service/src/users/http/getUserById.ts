import { Request } from 'express';
import { Problem } from '../../problem';
import { Method, RouteDefinition } from '../../server/routes';
import { buildSimpleHandler } from '../../server/simpleRoute';
import { UserRetriever } from '../retriever';
import { UserNotFoundError } from '../unknownUserError';
import { translateUserToResponse, UserResponseModel } from './model';
import { SINGLE_ROUTE_URI } from './urls';

/**
 * Handler for retrieving a single User by ID
 * @param userRetriever The means to retrieve users from the database
 */
export function buildGetUserByIdHandler(userRetriever: UserRetriever): RouteDefinition {
  return {
    handler: buildSimpleHandler<UserResponseModel>(async (req: Request) => {
      const userId = req.params.userId;

      try {
        const user = await userRetriever.getUserById(userId);
        return translateUserToResponse(user);
      } catch (e) {
        if (e instanceof UserNotFoundError) {
          throw new Problem(
            'tag:wyrdwest,2019:users/problems/unknown-user',
            'The requested user could not be found',
            404
          );
        } else {
          throw e;
        }
      }
    }),
    method: Method.GET,
    url: SINGLE_ROUTE_URI
  };
}
