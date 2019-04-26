import { RouteOptions } from 'fastify';
import { IncomingMessage, Server, ServerResponse } from 'http';
import { loadConfig } from './config';
import { Database } from './database';
import { buildHealthcheckHandler } from './healthchecks/handlers';
import buildServer from './server/index';

/**
 * Main entrypoint into the entire application
 */
function main(): void {
  const config = loadConfig();

  const database = new Database(config.get('pg.uri'));

  const handlers: ReadonlyArray<RouteOptions<Server, IncomingMessage, ServerResponse>> = [
    ...buildHealthcheckHandler({
      database
    })
  ];

  const server = buildServer();
  handlers.forEach(handler => server.route(handler));

  server.listen(config.get('http.port'), '::', (err, address) => {
    if (err) {
      throw err;
    }
    server.log.info(`server listening on ${address}`);
  });
}

main();
