exports.shorthands = undefined;

const SYSTEM_USER_ID = '00000000-0000-0000-0000-000000000000';
const SYSTEM_CLIENT_ID = '00000000-0000-0000-0000-000000000000';
const SYSTEM_CLIENT_SECRET = 'B99F9947-E026-4D17-BC34-52DA82955F0C';

exports.up = (pgm) => {
  pgm.sql(`INSERT INTO users(
      user_id,
      version,
      created,
      updated,
      name
    )
    VALUES(
      '${SYSTEM_USER_ID}',
      '00000000-0000-0000-0000-000000000000',
      current_timestamp,
      current_timestamp,
      'System OAuth2 User'
    )`);

  pgm.sql(`INSERT INTO oauth2_clients(
      client_id,
      version,
      created,
      updated,
      name,
      client_secret,
      owner_id,
      redirect_uris,
      response_types,
      grant_types
    )
    VALUES(
      '${SYSTEM_CLIENT_ID}',
      '00000000-0000-0000-0000-000000000000',
      current_timestamp,
      current_timestamp,
      'System OAuth2 Client',
      '${SYSTEM_CLIENT_SECRET}',
      '${SYSTEM_USER_ID}',
      '{}',
      '{}',
      '{}'
    )`);
};

exports.down = (pgm) => {
  pgm.sql(`DELETE FROM oauth2_clients WHERE client_id = '${SYSTEM_CLIENT_ID}'`);
  pgm.sql(`DELETE FROM users WHERE client_id = '${SYSTEM_USER_ID}'`);
};
