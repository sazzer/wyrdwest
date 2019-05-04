import bodyParser from 'body-parser';
import compression from 'compression';
import connectRid from 'connect-rid';
import cors from 'cors';
import errorhandler from 'errorhandler';
import express from 'express';
import helmet from 'helmet';
import morgan from 'morgan';
import responseTime from 'response-time';
import { Method, RouteDefinition } from './routes';

/**
 * Build the HTTP Server to work with
 */
export default function buildServer(handlers: readonly RouteDefinition[]): express.Express {
  const server = express();

  server.use(compression());
  server.use(bodyParser.json());
  server.use(bodyParser.urlencoded({ extended: false }));
  server.use(connectRid());
  server.use(cors());
  server.use(errorhandler());
  server.use(morgan('combined'));
  server.use(responseTime());
  server.use(helmet());

  const router = express.Router();

  const handlerDefinition = {
    [Method.GET]: router.get.bind(router),
    [Method.POST]: router.post.bind(router),
    [Method.PUT]: router.put.bind(router),
    [Method.PATCH]: router.patch.bind(router),
    [Method.DELETE]: router.delete.bind(router)
  };

  handlers.forEach(handler => {
    // tslint:disable
    console.log(`${handler.method} ${handler.url}`);

    const definitionFunction = handlerDefinition[handler.method];
    definitionFunction(handler.url, handler.handler);
  });

  server.use(router);
  return server;
}
