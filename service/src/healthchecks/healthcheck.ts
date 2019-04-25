export enum Status {
  OK = 'OK',
  FAIL = 'FAIL'
}

/**
 * Representation of the health of some component
 */
export interface HealthcheckResult {
  readonly status: Status;
  readonly detail?: string;
}

/**
 * Interface describing some component that we can check the health of
 */
export interface Healthcheck {
  /** Check the health of the component */
  readonly checkHealth: () => Promise<HealthcheckResult>;
}
