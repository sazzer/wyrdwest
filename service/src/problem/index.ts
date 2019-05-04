/**
 * Representation of a Problem as defined by RFC-7807
 */
export class Problem {
  constructor(
    public readonly type: string,
    public readonly title: string,
    public readonly status: number,
    public readonly detail?: string,
    public readonly instance?: string
  ) {}
}
