import { Request } from 'express';
import { Method, RouteDefinition } from '../server/routes';
import { buildSimpleHandler, SimpleResponse } from '../server/simpleRoute';
import { Healthcheck, Status } from './healthcheck';
import { checkHealth, SystemHealth } from './healthchecker';

/**
 * Build the Fastify handler definition for the healthchecks
 * @return the healthcheck handlers to use
 */
export function buildHealthcheckHandler(healthchecks: {
  readonly [key: string]: Healthcheck;
}): ReadonlyArray<RouteDefinition> {
  return [
    {
      handler: buildSimpleHandler(async function healthcheck(_: Request): Promise<SimpleResponse<SystemHealth>> {
        const health = await checkHealth(healthchecks);

        const statusCode = health.status === Status.OK ? 200 : 503;

        return new SimpleResponse<SystemHealth>(health, statusCode);
      }),
      method: Method.GET,
      url: '/health'
    }
  ];
}
