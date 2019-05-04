import express from 'express';
import { Express } from 'express-serve-static-core';
import bodyParser from 'body-parser';
import compression from 'compression';
import connectRid from 'connect-rid';
import cors from 'cors';
import errorhandler from 'errorhandler';
import morgan from 'morgan';
import responseTime from 'response-time';
import expressSlash from 'express-slash';
import helmet from 'helmet';

/**
 * Build the HTTP Server to work with
 */
export default function buildServer(): Express {
  const server = express();

  server.use(compression());
  server.use(bodyParser.json());
  server.use(bodyParser.urlencoded({ extended: false }));
  server.use(connectRid());
  server.use(cors());
  server.use(errorhandler());
  server.use(morgan('combined'));
  server.use(responseTime());
  server.use(expressSlash());
  server.use(helmet());

  return server;
}
