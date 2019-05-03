import test from 'ava';
import { Clock, Instant, ZoneOffset } from 'js-joda';
import { AccessTokenSerializer } from './accessTokenSerializer';
import { InvalidAccessTokenError, InvalidAccessTokenReason } from './invalidAccessTokenError';

/** The "current time" */
const NOW = Instant.parse('2019-05-02T16:59:00Z');

/** The clock to use */
const CLOCK = Clock.fixed(NOW, ZoneOffset.ofHours(0));

/** The expiry of an access token */
const EXPIRY = NOW.plusSeconds(3600);

/** The test subject */
const testSubject = new AccessTokenSerializer(CLOCK, 'mySecretKey', 'HS512');

test('Serialize an access token without scopes', async t => {
  const serialized = await testSubject.serialize({
    id: 'accessTokenId',
    client: 'clientId',
    user: 'userId',
    created: NOW,
    expires: EXPIRY
  });

  t.deepEqual(
    serialized,
    'eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTY4MTk5NDAsIm5iZiI6MTU1NjgxNjM0MCwiaWF0IjoxNTU2ODE2MzQwLCJhdWQiOiJjbGllbnRJZCIsImlzcyI6ImNsaWVudElkIiwic3ViIjoidXNlcklkIiwianRpIjoiYWNjZXNzVG9rZW5JZCJ9.nNViXO_b5omsLUyPqCw77w3jAj1rpadGVlvzGdQh78BluwYN3vuJSBUQhPDfw7DaIiV5WS62QFUQ2Ne0bgl0Dg'
  );
});

test('Serialize an access token with scopes', async t => {
  const serialized = await testSubject.serialize({
    id: 'accessTokenId',
    client: 'clientId',
    user: 'userId',
    created: NOW,
    expires: EXPIRY,
    scopes: ['admin', 'user']
  });

  t.deepEqual(
    serialized,
    'eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTY4MTk5NDAsIm5iZiI6MTU1NjgxNjM0MCwiaWF0IjoxNTU2ODE2MzQwLCJhdWQiOiJjbGllbnRJZCIsImlzcyI6ImNsaWVudElkIiwic3ViIjoidXNlcklkIiwic2NvcGVzIjpbImFkbWluIiwidXNlciJdLCJqdGkiOiJhY2Nlc3NUb2tlbklkIn0.mdYcKr-dCOK_ooP-dTnz7qKn18sDPf6gUG9BF0TaAATf-oQ2rDuBfUQqDHjl0JQfjSPFyj7rSNcvDZRWUnRTVw'
  );
});

test('Deserialize an access token without scopes', async t => {
  const token = await testSubject.deserialize(
    'eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTY4MTk5NDAsIm5iZiI6MTU1NjgxNjM0MCwiaWF0IjoxNTU2ODE2MzQwLCJhdWQiOiJjbGllbnRJZCIsImlzcyI6ImNsaWVudElkIiwic3ViIjoidXNlcklkIiwianRpIjoiYWNjZXNzVG9rZW5JZCJ9.nNViXO_b5omsLUyPqCw77w3jAj1rpadGVlvzGdQh78BluwYN3vuJSBUQhPDfw7DaIiV5WS62QFUQ2Ne0bgl0Dg'
  );

  t.deepEqual(token, {
    id: 'accessTokenId',
    client: 'clientId',
    user: 'userId',
    created: NOW,
    expires: EXPIRY,
    scopes: undefined
  });
});

test('Deserialize an access token with scopes', async t => {
  const token = await testSubject.deserialize(
    'eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTY4MTk5NDAsIm5iZiI6MTU1NjgxNjM0MCwiaWF0IjoxNTU2ODE2MzQwLCJhdWQiOiJjbGllbnRJZCIsImlzcyI6ImNsaWVudElkIiwic3ViIjoidXNlcklkIiwic2NvcGVzIjpbImFkbWluIiwidXNlciJdLCJqdGkiOiJhY2Nlc3NUb2tlbklkIn0.mdYcKr-dCOK_ooP-dTnz7qKn18sDPf6gUG9BF0TaAATf-oQ2rDuBfUQqDHjl0JQfjSPFyj7rSNcvDZRWUnRTVw'
  );

  t.deepEqual(token, {
    id: 'accessTokenId',
    client: 'clientId',
    user: 'userId',
    created: NOW,
    expires: EXPIRY,
    scopes: ['admin', 'user']
  });
});

test('Serialize and Deserialize an access token', async t => {
  const input = {
    id: 'accessTokenId',
    client: 'clientId',
    user: 'userId',
    created: NOW,
    expires: EXPIRY,
    scopes: ['admin', 'user']
  };

  const serialized = await testSubject.serialize(input);

  const token = await testSubject.deserialize(serialized);

  t.deepEqual(token, input);
});

test('Deserialize an expired access token', async t => {
  const input = {
    id: 'accessTokenId',
    client: 'clientId',
    user: 'userId',
    created: NOW.minusSeconds(7200),
    expires: EXPIRY.minusSeconds(7200),
    scopes: ['admin', 'user']
  };

  const serialized = await testSubject.serialize(input);

  const e = await t.throwsAsync(() => testSubject.deserialize(serialized));

  t.deepEqual(e, new InvalidAccessTokenError(InvalidAccessTokenReason.EXPIRY_IN_PAST));
});

test('Deserialize an access token created in the future', async t => {
  const input = {
    id: 'accessTokenId',
    client: 'clientId',
    user: 'userId',
    created: NOW.plusSeconds(7200),
    expires: EXPIRY.plusSeconds(7200),
    scopes: ['admin', 'user']
  };

  const serialized = await testSubject.serialize(input);

  const e = await t.throwsAsync(() => testSubject.deserialize(serialized));

  t.deepEqual(e, new InvalidAccessTokenError(InvalidAccessTokenReason.CREATED_IN_FUTURE));
});

test('Deserialize a malformed access token', async t => {
  const e = await t.throwsAsync(() => testSubject.deserialize('imInvalid'));

  t.deepEqual(e, new InvalidAccessTokenError(InvalidAccessTokenReason.MALFORMED_JWT));
});
