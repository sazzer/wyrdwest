/** Type representing a single row returned from a query */
export interface Row {
  readonly [field: string]: any;
}

/** Type representing a set of rows returned from a query */
export type RowSet = readonly Row[];

/** Interface representing a database connection */
export interface Database {
  /** Query the database for a set of rows */
  readonly query: (sql: string, ...args: readonly any[]) => Promise<RowSet>;

  /** Query the database, expecting exactly one row */
  readonly queryOne: (sql: string, ...args: readonly any[]) => Promise<Row>;
}
