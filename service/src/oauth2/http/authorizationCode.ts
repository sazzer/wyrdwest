import { Request, RequestHandler, Response } from 'express';
import { AccessTokenErrorCode } from './model';

/**
 * Build the Token Handler to use for an Authorization Code grant
 */
export function buildAuthorizationCodeTokenHandler(): RequestHandler {
  return async function handleAuthorizationCodeToken(_1: Request, res: Response): Promise<void> {
    res.status(400);
    res.json({
      error: AccessTokenErrorCode.INVALID_REQUEST
    });
  };
}
