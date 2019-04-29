import test from 'ava';
import { Identity, Model } from '../../service';
import { UserModel } from '../model';
import { translateUserToResponse } from './model';

test('Translate Full User', t => {
  const user: UserModel = new Model(new Identity('1', '2', new Date(), new Date()), {
    name: 'Graham',
    email: 'graham@grahamcox.co.uk',
    password: 'hashedPassword'
  });

  const response = translateUserToResponse(user);

  t.deepEqual(response, {
    self: '/users/1',
    name: 'Graham',
    email: 'graham@grahamcox.co.uk'
  });
});

test('Translate Minimal User', t => {
  const user: UserModel = new Model(new Identity('1', '2', new Date(), new Date()), {
    name: 'Graham'
  });

  const response = translateUserToResponse(user);

  t.deepEqual(response, {
    self: '/users/1',
    name: 'Graham',
    email: undefined
  });
});
