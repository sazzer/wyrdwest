import { ClientID } from "../clients/model";
import { UserID } from "../../users/model";

/**
 * Type representing the ID of an Access Token
 */
export type AccessTokenID = string;

/**
 * Type representing an Access Token
 */
export interface AccessToken {
  id: AccessTokenID;
  created: Date;
  expires: Date;
  client: ClientID,
  user: UserID,
  scopes: string[]
}
