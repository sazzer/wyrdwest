import { RouteOptions } from 'fastify';
import { IncomingMessage, Server, ServerResponse } from 'http';
import { loadConfig } from './config';
import { DatabaseWrapper } from './database/databaseWrapper';
import { buildHealthcheckHandler } from './healthchecks/handlers';
import { buildOAuth2Handlers } from './oauth2/http/buildHandlers';
import buildServer from './server';
import { buildUsersDao } from './users/dao/dao';
import { buildUserHandlers } from './users/http/buildHandlers';

/**
 * Main entrypoint into the entire application
 */
async function main(): Promise<void> {
  const config = loadConfig();

  const database = new DatabaseWrapper(config.get('pg.uri'));
  await database.migrate();

  const usersDao = buildUsersDao(database);

  const handlers: ReadonlyArray<RouteOptions<Server, IncomingMessage, ServerResponse>> = [
    ...buildHealthcheckHandler({ database }),
    ...buildUserHandlers(usersDao),
    ...buildOAuth2Handlers()
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
