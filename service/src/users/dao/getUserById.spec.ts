import test from 'ava';
import { TestDatabase } from '../../database/testDatabase';
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
