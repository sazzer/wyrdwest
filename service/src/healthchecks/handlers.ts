import { FastifyReply, FastifyRequest, RouteOptions } from 'fastify';
import { IncomingMessage, Server, ServerResponse } from 'http';
import { Healthcheck, Status } from './healthcheck';
import { checkHealth, SystemHealth } from './healthchecker';

/**
 * Build the Fastify handler definition for the healthchecks
 * @return the healthcheck handlers to use
 */
export function buildHealthcheckHandler(healthchecks: {
  readonly [key: string]: Healthcheck;
}): ReadonlyArray<RouteOptions<Server, IncomingMessage, ServerResponse>> {
  return [
    {
      handler: async function healthcheck(
        _: FastifyRequest<IncomingMessage>,
        reply: FastifyReply<ServerResponse>
      ): Promise<SystemHealth> {
        const health = await checkHealth(healthchecks);
        reply.status(health.status === Status.OK ? 200 : 503);

        return health;
      },
      method: 'GET',
      url: '/health'
    }
  ];
}
