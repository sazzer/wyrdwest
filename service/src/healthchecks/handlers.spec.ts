import test from 'ava';
import supertest from 'supertest';
import buildServer from '../server';
import { buildHealthcheckHandler } from './handlers';
import { Healthcheck, Status } from './healthcheck';

function runTest(healthchecks: { readonly [key: string]: Healthcheck }): supertest.Test {
  const server = buildServer(buildHealthcheckHandler(healthchecks));

  return supertest(server).get('/health');
}

const passing = {
  checkHealth: () => {
    return Promise.resolve({
      detail: 'Success',
      status: Status.OK
    });
  }
};
const failing = {
  checkHealth: () => {
    return Promise.resolve({
      detail: 'Failure',
      status: Status.FAIL
    });
  }
};

test('No Healthchecks', async t => {
  const response = await runTest({});
  t.deepEqual(response.status, 200);
  t.deepEqual(response.body, {
    details: {},
    status: 'OK'
  });
});

test('One passing Healthcheck', async t => {
  const response = await runTest({ passing });
  t.deepEqual(response.status, 200);
  t.deepEqual(response.body, {
    details: {
      passing: {
        detail: 'Success',
        status: 'OK'
      }
    },
    status: 'OK'
  });
});

test('One failing Healthcheck', async t => {
  const response = await runTest({ failing });
  t.deepEqual(response.status, 503);
  t.deepEqual(response.body, {
    details: {
      failing: {
        detail: 'Failure',
        status: 'FAIL'
      }
    },
    status: 'FAIL'
  });
});

test('Mixed Healthchecks', async t => {
  const response = await runTest({ failing, passing });
  t.deepEqual(response.status, 503);
  t.deepEqual(response.body, {
    details: {
      failing: {
        detail: 'Failure',
        status: 'FAIL'
      },
      passing: {
        detail: 'Success',
        status: 'OK'
      }
    },
    status: 'FAIL'
  });
});
