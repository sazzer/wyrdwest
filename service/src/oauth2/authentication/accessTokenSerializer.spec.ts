import test from 'ava';
import { Instant } from 'js-joda';
import { AccessTokenSerializer } from './accessTokenSerializer';

/** The test subject */
const testSubject = new AccessTokenSerializer('mySecretKey', 'HS512');

test('Serialize an access token without scopes', async t => {
  const serialized = await testSubject.serialize({
    id: 'accessTokenId',
    client: 'clientId',
    user: 'userId',
    created: Instant.parse('2019-05-02T16:59:00Z'),
    expires: Instant.parse('2019-05-02T17:59:00Z')
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
    created: Instant.parse('2019-05-02T16:59:00Z'),
    expires: Instant.parse('2019-05-02T17:59:00Z'),
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
    created: Instant.parse('2019-05-02T16:59:00Z'),
    expires: Instant.parse('2019-05-02T17:59:00Z'),
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
    created: Instant.parse('2019-05-02T16:59:00Z'),
    expires: Instant.parse('2019-05-02T17:59:00Z'),
    scopes: ['admin', 'user']
  });
});
