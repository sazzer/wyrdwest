import { AccessTokenModel } from './model';

/**
 * Representation of a function to handle a token grant
 */
export type TokenHandlerFunction = (params: { readonly [key: string]: string }) => Promise<AccessTokenModel>;
