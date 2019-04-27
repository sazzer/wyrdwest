import { fc, testProp } from 'ava-fast-check';
import { checkPassword, hashPassword } from './password';

testProp('Hash password', [fc.string()], async password => {
  const hash = await hashPassword(password);
  return hash !== password;
});

testProp('Re-hash password', [fc.string()], async password => {
  const hash = await hashPassword(password);
  const hash2 = await hashPassword(password);
  return hash !== hash2;
});

testProp('Check password - correct', [fc.string()], async password => {
  const hash = await hashPassword(password);
  return checkPassword(hash, password);
});

testProp('Check password - incorrect', [fc.string()], async password => {
  const hash = await hashPassword('password');
  return !(await checkPassword(hash, password));
});
