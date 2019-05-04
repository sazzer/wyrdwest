import { Request, RequestHandler, Response } from 'express';
import { AccessTokenErrorCode } from './model';

/**
 * Build the Token Handler to use for an Client Credentials grant
 */
export function buildClientCredentialsTokenHandler(): RequestHandler {
  return async function handleClientCredentialsToken(_1: Request, res: Response): Promise<void> {
    res.status(400);
    res.json({
      error: AccessTokenErrorCode.INVALID_REQUEST
    });
  };
}
