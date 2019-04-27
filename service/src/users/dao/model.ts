import { Row } from '../../database/database';
import { Identity, Model } from '../../service';
import { UserModel } from '../model';

/**
 * Parse a database row into a User model
 * @param row The row to parse
 * @return the parsed row
 */
export function parseDatabaseRow(row: Row): UserModel {
  return new Model(new Identity(row.user_id, row.version, row.created, row.updated), {
    name: row.name,
    password: row.password,
    email: row.email
  });
}
