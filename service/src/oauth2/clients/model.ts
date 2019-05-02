import { Model } from '../../service';
import { UserID } from '../../users/model';

/** Enumeration of possible response types */
enum ResponseTypes {
  CODE = 'code',
  TOKEN = 'token',
  ID_TOKEN = 'id_token'
}

/** Enumeration of possible grant types */
enum GrantTypes {
  AUTHORIZATION_CODE = 'authorization_code',
  IMPLICIT = 'implicit',
  CLIENT_CREDENTIALS = 'client_credentials',
  REFRESH_TOKEN = 'refresh_token'
}

/**
 * Type representing the data in an OAuth2 Client
 */
export interface ClientData {
  readonly name: string;
  readonly secret: string;
  readonly owner: UserID;
  readonly redirectUris: readonly ResponseTypes[];
  readonly responseTypes: readonly GrantTypes[];
  readonly grantTypes: readonly string[];
}

/**
 * Type representing the ID of an OAuth2 Client
 */
export type ClientID = string;

/**
 * Type representing the actual client resource model
 */
export type ClientModel = Model<ClientID, ClientData>;
