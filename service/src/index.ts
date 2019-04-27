import { RouteOptions } from 'fastify';
import { IncomingMessage, Server, ServerResponse } from 'http';
import { loadConfig } from './config';
import { DatabaseWrapper } from './database/databaseWrapper';
import { buildHealthcheckHandler } from './healthchecks/handlers';
import buildServer from './server';

/**
 * Main entrypoint into the entire application
 */
async function main(): Promise<void> {
  const config = loadConfig();

  const database = new DatabaseWrapper(config.get('pg.uri'));
  await database.migrate();

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
