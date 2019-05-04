import { Request, RequestHandler, Response } from 'express';
import { AccessTokenErrorCode } from './model';

/**
 * Build the Token Handler to use for an Refresh Token grant
 */
export function buildRefreshTokenHandler(): RequestHandler {
  return async function handleRefreshToken(_1: Request, res: Response): Promise<void> {
    res.status(400);
    res.json({
      error: AccessTokenErrorCode.INVALID_REQUEST
    });
  };
}
