/**
 * Error indicating that we expected rows but didn't get any
 */
export class NoRowsReturnedError extends Error {
  constructor() {
    super('No rows returned');
  }
}
