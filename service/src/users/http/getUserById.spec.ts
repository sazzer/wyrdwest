import test from 'ava';
import supertest from 'supertest';
import td from 'testdouble';
import buildServer from '../../server';
import { Identity, Model } from '../../service';
import { UserRetriever } from '../retriever';
import { UserNotFoundError } from '../unknownUserError';
import { buildGetUserByIdHandler } from './getUserById';

/** The User ID to work with */
const USER_ID = '47FC5F8A-4065-40D0-A1E0-9F502F8C8666';

/**
 * Actually run the test
 * @param userRetriever the mock user retriever to use
 * @param url the URL to call
 */
function runTest(userRetriever: UserRetriever, url: string): supertest.Test {
  const server = buildServer([buildGetUserByIdHandler(userRetriever)]);

  return supertest(server).get(url);
}

test('Get Unknown User', async t => {
  const userRetriever = td.object(['getUserById']);
  td.when(userRetriever.getUserById(USER_ID)).thenReject(new UserNotFoundError(USER_ID));

  const response = await runTest(userRetriever, `/users/${USER_ID}`);
  t.deepEqual(response.status, 404);
  t.deepEqual(response.body, {
    type: 'tag:wyrdwest,2019:users/problems/unknown-user',
    title: 'The requested user could not be found',
    status: 404
  });
});

test('Get Minimal User', async t => {
  const user = new Model(new Identity(USER_ID, 'version', new Date(), new Date()), {
    name: 'Test User'
  });

  const userRetriever = td.object(['getUserById']);
  td.when(userRetriever.getUserById(USER_ID)).thenResolve(user);

  const response = await runTest(userRetriever, `/users/${USER_ID}`);
  t.deepEqual(response.status, 200);
  t.deepEqual(response.body, {
    self: `/users/${USER_ID}`,
    name: 'Test User'
  });
});

test('Get Full User', async t => {
  const user = new Model(new Identity(USER_ID, 'version', new Date(), new Date()), {
    name: 'Test User',
    email: 'testuser@example.com',
    password: 'hashedPassword'
  });

  const userRetriever = td.object(['getUserById']);
  td.when(userRetriever.getUserById(USER_ID)).thenResolve(user);

  const response = await runTest(userRetriever, `/users/${USER_ID}`);
  t.deepEqual(response.status, 200);
  t.deepEqual(response.body, {
    self: `/users/${USER_ID}`,
    name: 'Test User',
    email: 'testuser@example.com'
  });
});
