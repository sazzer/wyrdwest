import { Pool } from 'pg';
import { HealthcheckResult, Status } from '../healthchecks/healthcheck';

/** Wrapper around a database */
export class Database {
  /** The connection pool */
  public readonly pool: Pool;

  /**
   * Construct the database wrapper
   * @param uri the URI to use
   */
  constructor(uri: string) {
    this.pool = new Pool({
      connectionString: uri
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
