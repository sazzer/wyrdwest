/**
 * Representation of the Identity of some resource
 */
export class Identity<T> {
  /**
   * Construct the resource identity
   * @param id The ID of the resource
   * @param version The version of the resource
   * @param created When the resource was created
   * @param updated When the resource was last updated
   */
  constructor(
    public readonly id: T,
    public readonly version: string,
    public readonly created: Date,
    public readonly updated: Date
  ) {}
}
