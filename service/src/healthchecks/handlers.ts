import { FastifyReply, FastifyRequest, RouteOptions } from 'fastify';
import { IncomingMessage, Server, ServerResponse } from 'http';

/**
 * Handler to actually perform the healthchecks and send back the result
 * @param reply the means to send the response
 */
function healthcheck(
  _: FastifyRequest<IncomingMessage>,
  reply: FastifyReply<ServerResponse>
): void {
  reply.send({
    status: 'OK'
  });
}

/**
 * Build the Fastify handler definition for the healthchecks
 * @return the healthcheck handlers to use
 */
export function buildHealthcheckHandler(): ReadonlyArray<
  RouteOptions<Server, IncomingMessage, ServerResponse>
> {
  return [
    {
      handler: healthcheck,
      method: 'GET',
      url: '/health'
    }
  ];
}
