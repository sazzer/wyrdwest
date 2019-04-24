import fastify from 'fastify';
import fastifyCors from 'fastify-cors';
import fastifyHelmet from 'fastify-helmet';
import fastifySensible from 'fastify-sensible';
import { IncomingMessage, Server, ServerResponse } from 'http';

export default function buildServer(): fastify.FastifyInstance<
  Server,
  IncomingMessage,
  ServerResponse
> {
  const server = fastify({
    logger: true
  });

  server.register(fastifyCors, {
    credentials: true,
    methods: ['GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'HEAD'],
    origin: true,
    preflightContinue: true
  });
  server.register(fastifyHelmet);
  server.register(fastifySensible);

  return server;
}
