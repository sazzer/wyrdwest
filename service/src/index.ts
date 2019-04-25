import { RouteOptions } from 'fastify';
import { IncomingMessage, Server, ServerResponse } from 'http';
import { buildHealthcheckHandler } from './healthchecks/handlers';
import { Status } from './healthchecks/healthcheck';
import buildServer from './server/index';

/**
 * Main entrypoint into the entire application
 */
function main(): void {
  const server = buildServer();

  const handlers: ReadonlyArray<RouteOptions<Server, IncomingMessage, ServerResponse>> = [
    ...buildHealthcheckHandler({
      failing: {
        checkHealth: () => {
          return Promise.resolve({
            detail: 'Failure',
            status: Status.FAIL
          });
        }
      }
    })
  ];

  handlers.forEach(handler => server.route(handler));

  server.listen(3000, (err, address) => {
    if (err) {
      throw err;
    }
    server.log.info(`server listening on ${address}`);
  });
}

main();
