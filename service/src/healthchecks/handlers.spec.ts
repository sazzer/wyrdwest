import test from 'ava';
import fastify, { HTTPInjectResponse } from 'fastify';
import { buildHealthcheckHandler } from './handlers';
import { Healthcheck, Status } from './healthcheck';

function runTest(healthchecks: { readonly [key: string]: Healthcheck }): Promise<HTTPInjectResponse> {
  const server = fastify();

  buildHealthcheckHandler(healthchecks);//.forEach(handler => server.route(handler));

  return server.inject({
    method: 'GET',
    url: '/health'
  });
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
  t.deepEqual(response.statusCode, 200);
  t.deepEqual(JSON.parse(response.payload), {
    details: {},
    status: 'OK'
  });
});

test('One passing Healthcheck', async t => {
  const response = await runTest({ passing });
  t.deepEqual(response.statusCode, 200);
  t.deepEqual(JSON.parse(response.payload), {
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
  t.deepEqual(response.statusCode, 503);
  t.deepEqual(JSON.parse(response.payload), {
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
  t.deepEqual(response.statusCode, 503);
  t.deepEqual(JSON.parse(response.payload), {
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
