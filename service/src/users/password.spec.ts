import test from 'ava';
import { checkPassword, hashPassword } from './password';

test('Hash password', async t => {
  const hash = await hashPassword('password');

  t.notDeepEqual(hash, 'password');
});

test('Re-hash password', async t => {
  const hash = await hashPassword('password');
  const hash2 = await hashPassword('password');

  t.notDeepEqual(hash, hash2);
});

test('Check password - Correct', async t => {
  const hash = await hashPassword('password');
  const result = await checkPassword(hash, 'password');

  t.true(result);
});

test('Check password - Incorrect', async t => {
  const hash = await hashPassword('password');
  const result = await checkPassword(hash, 'password2');

  t.false(result);
});
