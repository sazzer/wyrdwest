import { Request, Response } from 'express';
import { Method, RouteDefinition } from '../server/routes';
import { Healthcheck, Status } from './healthcheck';
import { checkHealth } from './healthchecker';

/**
 * Build the Fastify handler definition for the healthchecks
 * @return the healthcheck handlers to use
 */
export function buildHealthcheckHandler(healthchecks: {
  readonly [key: string]: Healthcheck;
}): ReadonlyArray<RouteDefinition> {
  return [
    {
      handler: async function healthcheck(_: Request, res: Response): Promise<void> {
        const health = await checkHealth(healthchecks);
        res.status(health.status === Status.OK ? 200 : 503);
        res.json(health);
      },
      method: Method.GET,
      url: '/health'
    }
  ];
}
