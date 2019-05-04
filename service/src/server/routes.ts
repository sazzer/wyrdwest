import { RequestHandler } from 'express';

/**
 * Enumeration of HTTP Methods that can be supported
 */
export enum Method {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
  DELETE = 'DELETE',
  PATCH = 'PATCH'
}

/**
 * Representation of a Route Definition to register
 */
export interface RouteDefinition {
  readonly url: string;
  readonly method: Method;
  readonly handler: RequestHandler;
}
