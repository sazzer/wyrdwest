import { Healthcheck, HealthcheckResult, Status } from './healthcheck';

/**
 * Representation of the overall healthcheck result for the system
 */
export interface SystemHealth {
  readonly status: Status;
  readonly details: { readonly [key: string]: HealthcheckResult };
}

/**
 * Check the health of the provided components and return the collected detils
 * @param healthchecks the healthchecks to perform
 * @return the overall system health
 */
export async function checkHealth(healthchecks: { readonly [key: string]: Healthcheck }): Promise<SystemHealth> {
  const keys = Object.keys(healthchecks);
  const checks = Object.keys(healthchecks)
    .map(hc => healthchecks[hc])
    .map(hc => hc.checkHealth());
  const results = await Promise.all(checks);

  const details = keys
    .map((key, index) => ({
      key,
      result: results[index]
    }))
    .reduce((obj, next) => {
      return {
        ...obj,
        [next.key]: next.result
      };
    }, {});

  const overallStatus = results.map(result => result.status).every(status => status === Status.OK)
    ? Status.OK
    : Status.FAIL;

  return {
    details,
    status: overallStatus
  };
}
