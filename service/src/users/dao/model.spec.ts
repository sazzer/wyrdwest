import test from 'ava';
import { Identity, Model } from '../../service';
import { parseDatabaseRow } from './model';

test('Parse full user', t => {
  const row = {
    user_id: '1',
    version: '2',
    created: new Date('2019-04-01T12:34:56Z'),
    updated: new Date('2019-04-27T20:07:00Z'),
    name: 'Graham',
    email: 'graham@grahamcox.co.uk',
    password: 'hashedPassword'
  };

  const user = parseDatabaseRow(row);

  t.deepEqual(
    user,
    new Model(new Identity('1', '2', new Date('2019-04-01T12:34:56Z'), new Date('2019-04-27T20:07:00Z')), {
      name: 'Graham',
      email: 'graham@grahamcox.co.uk',
      password: 'hashedPassword'
    })
  );
});

test('Parse minimal user', t => {
  const row = {
    user_id: '1',
    version: '2',
    created: new Date('2019-04-01T12:34:56Z'),
    updated: new Date('2019-04-27T20:07:00Z'),
    name: 'Graham'
  };

  const user = parseDatabaseRow(row);

  t.deepEqual(
    user,
    new Model(new Identity('1', '2', new Date('2019-04-01T12:34:56Z'), new Date('2019-04-27T20:07:00Z')), {
      name: 'Graham',
      email: undefined,
      password: undefined
    })
  );
});
