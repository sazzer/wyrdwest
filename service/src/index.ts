import { buildHealthcheckHandler } from './healthchecks/handlers';
import buildServer from './server/index';

/**
 * Main entrypoint into the entire application
 */
function main(): void {
  const server = buildServer();

  const handlers: ReadonlyArray<any> = [...buildHealthcheckHandler()];

  handlers.forEach(handler => server.route(handler));

  server.listen(3000, (err, address) => {
    if (err) {
      throw err;
    }
    server.log.info(`server listening on ${address}`);
  });
}

main();
