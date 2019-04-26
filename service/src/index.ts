import { RouteOptions } from 'fastify';
import { IncomingMessage, Server, ServerResponse } from 'http';
import { loadConfig } from './config';
import { buildHealthcheckHandler } from './healthchecks/handlers';
import { Status } from './healthchecks/healthcheck';
import buildServer from './server/index';

/**
 * Main entrypoint into the entire application
 */
function main(): void {
  const config = loadConfig();
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

  server.listen(config.get('http.port'), (err, address) => {
    if (err) {
      throw err;
    }
    server.log.info(`server listening on ${address}`);
  });
}

main();
