import test from 'ava';
import { TestDatabase } from '../../database/testDatabase';
import { Identity, Model } from '../../service';
import { UserNotFoundError } from '../unknownUserError';
import { getUserById } from './getUserById';

test('Get Unknown User', async t => {
  const db = new TestDatabase();
  db.expect('SELECT * FROM users WHERE user_id = $1', ['1'], []);

  try {
    await getUserById(db, '1');
    t.fail('Expected an exception');
  } catch (e) {
    t.deepEqual(e, new UserNotFoundError('1'));
  }
});

test('Get Full User', async t => {
  const row = {
    user_id: '1',
    version: '2',
    created: new Date('2019-04-01T12:34:56Z'),
    updated: new Date('2019-04-27T20:07:00Z'),
    name: 'Graham',
    email: 'graham@grahamcox.co.uk',
    password: 'hashedPassword'
  };

  const db = new TestDatabase();
  db.expect('SELECT * FROM users WHERE user_id = $1', ['1'], [row]);

  const user = await getUserById(db, '1');
  t.deepEqual(
    user,
    new Model(new Identity('1', '2', new Date('2019-04-01T12:34:56Z'), new Date('2019-04-27T20:07:00Z')), {
      name: 'Graham',
      email: 'graham@grahamcox.co.uk',
      password: 'hashedPassword'
    })
  );
});

test('Get Minimal User', async t => {
  const row = {
    user_id: '1',
    version: '2',
    created: new Date('2019-04-01T12:34:56Z'),
    updated: new Date('2019-04-27T20:07:00Z'),
    name: 'Graham'
  };

  const db = new TestDatabase();
  db.expect('SELECT * FROM users WHERE user_id = $1', ['1'], [row]);

  const user = await getUserById(db, '1');
  t.deepEqual(
    user,
    new Model(new Identity('1', '2', new Date('2019-04-01T12:34:56Z'), new Date('2019-04-27T20:07:00Z')), {
      name: 'Graham',
      email: undefined,
      password: undefined
    })
  );
});
