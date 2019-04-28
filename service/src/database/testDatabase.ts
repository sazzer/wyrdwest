// tslint:disable:readonly-array
// tslint:disable:readonly-keyword

import { RowSet } from './database';
import { DatabaseBase } from './databaseBase';

interface ExpectedCalls {
  readonly sql: string;
  readonly binds: readonly any[];
  readonly result: RowSet;
}

/**
 * Implementation of the Database interface for testing against
 */
export class TestDatabase extends DatabaseBase {
  /** The expectations for the database */
  private readonly expectations: ExpectedCalls[] = [];

  /**
   * Expect that a query will be called
   * @param sql The SQL to expect
   * @param binds The binds to expect
   * @param result The result to return
   */
  public expect(sql: string, binds: readonly any[], result: RowSet): void {
    this.expectations.push({ sql, binds, result });
  }

  public async query(sql: string, ...args: readonly any[]): Promise<RowSet> {
    const result = this.expectations
      .filter(e => e.sql === sql)
      .filter(e => e.binds.length === args.length)
      .filter(e => JSON.stringify(e.binds) === JSON.stringify(args));

    if (result.length === 0) {
      throw new Error(`Unexpected SQL "${sql}" and binds [${args}]`);
    }

    return result[0].result;
  }
}
