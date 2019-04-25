import test from 'ava';
import { HealthcheckResult, Status } from './healthcheck';
import { checkHealth } from './healthchecker';

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
  const result = await checkHealth({});
  t.deepEqual(result, { status: 'OK', details: {} });
});

test('One passing Healthcheck', async t => {
  const result = await checkHealth({ passing });
  t.deepEqual(result, {
    details: {
      passing: {
        detail: 'Success',
        status: Status.OK
      }
    },
    status: Status.OK
  });
});

test('One failing healthcheck', async t => {
  const result = await checkHealth({ failing });
  t.deepEqual(result, {
    details: {
      failing: {
        detail: 'Failure',
        status: Status.FAIL
      }
    },
    status: Status.FAIL
  });
});

test('Mixed results', async t => {
  const result = await checkHealth({ passing, failing });
  t.deepEqual(result, {
    details: {
      failing: {
        detail: 'Failure',
        status: Status.FAIL
      },
      passing: {
        detail: 'Success',
        status: Status.OK
      }
    },
    status: Status.FAIL
  });
});

test('One slow healthcheck', async t => {
  const slow = {
    checkHealth: () => {
      return new Promise<HealthcheckResult>(resolve => {
        setTimeout(() => {
          resolve({
            detail: 'Slow',
            status: Status.OK
          });
        }, 1000);
      });
    }
  };

  const result = await checkHealth({ passing, failing, slow });
  t.deepEqual(result, {
    details: {
      failing: {
        detail: 'Failure',
        status: Status.FAIL
      },
      passing: {
        detail: 'Success',
        status: Status.OK
      },
      slow: {
        detail: 'Slow',
        status: Status.OK
      }
    },
    status: Status.FAIL
  });
});
