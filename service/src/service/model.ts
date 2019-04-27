import { Identity } from './identity';
/**
 * Representation of a Resource as it exists in the database
 */
export class Model<ID, DATA> {
  /**
   * Construct the resource model
   * @param identity The Identity of the resource
   * @param data The data in the resource
   */
  constructor(public readonly identity: Identity<ID>, public readonly data: DATA) {}
}
