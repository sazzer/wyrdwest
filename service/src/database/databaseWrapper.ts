import pgMigrate from 'node-pg-migrate';
import { Pool } from 'pg';
import { HealthcheckResult, Status } from '../healthchecks';
import { RowSet } from './database';
import { DatabaseBase } from './databaseBase';

/** Wrapper around a database */
export class DatabaseWrapper extends DatabaseBase {
  /** The connection pool */
  private readonly pool: Pool;

  /** The connection URI */
  private readonly uri: string;
  /**
   * Construct the database wrapper
   * @param uri the URI to use
   */
  constructor(uri: string) {
    super();
    this.uri = uri;
    this.pool = new Pool({
      connectionString: uri
    });
  }

  /**
   * Execute a query against the database
   * @param sql The SQL to execute
   * @param args The arguments for the query
   */
  public async query(sql: string, ...args: readonly any[]): Promise<RowSet> {
    const result = await this.pool.query(sql, [...args]);
    return result.rows;
  }

  public async migrate(): Promise<void> {
    await pgMigrate({
      checkOrder: true,
      count: -1,
      databaseUrl: this.uri,
      dir: 'migrations',
      direction: 'up',
      ignorePattern: '.',
      migrationsTable: '_migrations',
      singleTransaction: true
    });
  }

  /**
   * Check the health of the database
   */
  public async checkHealth(): Promise<HealthcheckResult> {
    try {
      await this.pool.query('SELECT 1');
      return {
        status: Status.OK
      };
    } catch (e) {
      return {
        detail: e,
        status: Status.FAIL
      };
    }
  }
}
