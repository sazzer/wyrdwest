import pgMigrate from 'node-pg-migrate';
import { Pool } from 'pg';
import { HealthcheckResult, Status } from '../healthchecks/healthcheck';

/** Wrapper around a database */
export class Database {
  /** The connection pool */
  private readonly pool: Pool;

  /** The connection URI */
  private readonly uri: string;
  /**
   * Construct the database wrapper
   * @param uri the URI to use
   */
  constructor(uri: string) {
    this.uri = uri;
    this.pool = new Pool({
      connectionString: uri
    });
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
