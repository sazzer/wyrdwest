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

  [
    ...buildHealthcheckHandler({ database }),
    ...buildUserHandlers(usersDao),
    ...buildOAuth2Handlers()
  ];

  const server = buildServer();

  const port = config.get('http.port');
  server.listen(port, () => {
  });
}

main();
