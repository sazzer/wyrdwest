import { Request, Response } from 'express';
import { Method, RouteDefinition } from '../../server/routes';
import { UserRetriever } from '../retriever';
import { UserNotFoundError } from '../unknownUserError';
import { translateUserToResponse } from './model';
import { SINGLE_ROUTE_URI } from './urls';

/**
 * Handler for retrieving a single User by ID
 * @param userRetriever The means to retrieve users from the database
 */
export function buildGetUserByIdHandler(userRetriever: UserRetriever): RouteDefinition {
  return {
    async handler(req: Request, res: Response): Promise<void> {
      const userId = req.params.userId;

      try {
        const user = await userRetriever.getUserById(userId);
        res.status(200);
        res.json(translateUserToResponse(user));
      } catch (e) {
        if (e instanceof UserNotFoundError) {
          res.status(404);
          res.json({
            type: 'tag:wyrdwest,2019:users/problems/unknown-user',
            title: 'The requested user could not be found',
            status: 404
          });
        } else {
          res.status(500);
          res.json({
            type: 'tag:wyrdwest,2019:problems/internal-server-error',
            title: 'An unexpected error occurred',
            status: 500
          });
        }
      }
    },
    method: Method.GET,
    url: SINGLE_ROUTE_URI
  };
}
