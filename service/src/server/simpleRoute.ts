import { Request, RequestHandler, Response } from 'express';
import { Problem } from '../problem';

/** Type representing the simplified handler */
export type SimpleHandler<T> = (req: Request) => Promise<T>;

/**
 * Simple wrapper around the response to send
 */
export class SimpleResponse<T> {
  /**
   * Construct the response
   * @param body The response body to send
   * @param statusCode The status code to send
   */
  constructor(public readonly body: T, public readonly statusCode: number = 200) {}
}
/**
 * Wrap a simple handler method to make it work with Express
 * @param handler The handler to wrap
 */
export function buildSimpleHandler<T>(handler: SimpleHandler<T>): RequestHandler {
  return async (req: Request, res: Response) => {
    try {
      const result = await handler(req);
      if (result instanceof SimpleResponse) {
        res.status(result.statusCode);
        res.json(result.body);
      } else {
        res.json(result);
      }
    } catch (e) {
      res.contentType('application/problem+json');
      if (e instanceof Problem) {
        res.status(e.status);
        res.json(e);
      } else if (e instanceof Error) {
        res.status(500);
        res.json({
          type: 'tag:wyrdwest,2019:problems/unexpected-error',
          title: e.message,
          status: 500
        });
      } else {
        res.status(500);
        res.json({
          type: 'tag:wyrdwest,2019:problems/unexpected-error',
          title: 'An unexpected error occurred',
          status: 500
        });
      }
    }
  };
}
