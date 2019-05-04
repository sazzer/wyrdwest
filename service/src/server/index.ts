import bodyParser from 'body-parser';
import compression from 'compression';
import connectRid from 'connect-rid';
import cors from 'cors';
import errorhandler from 'errorhandler';
import express from 'express';
import helmet from 'helmet';
import morgan from 'morgan';
import responseTime from 'response-time';
import { RouteDefinition } from './routes';

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

  handlers.forEach(handler => {
    router.get(handler.url, handler.handler);
  });

  server.use(router);
  return server;
}
