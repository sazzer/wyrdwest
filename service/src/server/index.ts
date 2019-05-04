import express from 'express';
import { Express } from 'express-serve-static-core';

/**
 * Build the HTTP Server to work with
 */
export default function buildServer(): Express {
  const server = express();

  return server;
}
