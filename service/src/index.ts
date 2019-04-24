import buildServer from './server/index';

function main(): void {
  const server = buildServer();

  server.listen(3000, (err, address) => {
    if (err) {
      throw err;
    }
    server.log.info(`server listening on ${address}`);
  });
}

main();
