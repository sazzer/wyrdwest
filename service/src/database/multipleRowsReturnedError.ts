/**
 * Error indicating that we expected exactly 1 row but got more
 */
export class MultipleRowsReturnedError extends Error {
  constructor(public readonly count: number) {
    super(`Expected 1 row, but got ${count}`);
  }
}
