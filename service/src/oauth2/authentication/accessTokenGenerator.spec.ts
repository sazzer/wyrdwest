import test from 'ava';
import { Clock, Duration, Instant, ZoneOffset } from 'js-joda';
import td from 'testdouble';

td.replace('uuid');
import { v4 as uuid } from 'uuid';

import { AccessTokenGenerator } from './accessTokenGenerator';

/** The "current time" */
const NOW = Instant.parse('2019-05-02T10:40:00Z');

/** The clock to use */
const CLOCK = Clock.fixed(NOW, ZoneOffset.ofHours(0));

/** The duration of an access token */
const DURATION = Duration.parse('PT1H');

/** The test subject */
const testSubject = new AccessTokenGenerator(CLOCK, DURATION);

test.before(() => {
  td.when(uuid()).thenReturn('theId');
});

test.after(() => {
  td.reset();
});

test('Generate with no scopes', t => {
  const accessToken = testSubject.generate('user', 'client');
  t.deepEqual(accessToken, {
    id: 'theId',
    client: 'client',
    user: 'user',
    created: NOW,
    expires: NOW.plus(DURATION),
    scopes: undefined
  });
});

test('Generate with scopes', t => {
  const accessToken = testSubject.generate('user', 'client', ['admin', 'other']);
  t.deepEqual(accessToken, {
    id: 'theId',
    client: 'client',
    user: 'user',
    created: NOW,
    expires: NOW.plus(DURATION),
    scopes: ['admin', 'other']
  });
});
