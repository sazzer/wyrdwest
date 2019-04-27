import { Database, Row, RowSet } from './database';
import { MultipleRowsReturnedError } from './multipleRowsReturnedError';
import { NoRowsReturnedError } from './noRowsReturnedError';

/**
 * Base class that the Database implementations can inherit from
 */
export abstract class DatabaseBase implements Database {
  public abstract query(sql: string, ...args: readonly any[]): Promise<RowSet>;

  /** Query the database, expecting exactly one row */
  public async queryOne(sql: string, ...args: readonly any[]): Promise<Row> {
    const rows = await this.query(sql, args);

    if (rows.length === 0) {
      // No rows returned
      throw new NoRowsReturnedError();
    } else if (rows.length >= 2) {
      // Too many rows returned
      throw new MultipleRowsReturnedError(rows.length);
    } else {
      return rows[0];
    }
  }
}
