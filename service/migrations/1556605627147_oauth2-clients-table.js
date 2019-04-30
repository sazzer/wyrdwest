exports.shorthands = undefined;

exports.up = (pgm) => {
  pgm.createTable('oauth2_clients', {
    client_id: {
        type: 'uuid',
        primaryKey: true
    },
    version: {
        type: 'uuid',
        notNull: true
    },
    created: {
        type: 'timestamp',
        notNull: true
    },
    updated: {
        type: 'timestamp',
        notNull: true
    },
    name: {
      type: 'text',
      notNull: true
    },
    client_secret: {
      type: 'uuid',
      notNull: true
    },
    owner_id: {
      type: 'uuid',
      notNull: true,
      references: 'users',
      onDelete: 'RESTRICT',
      onUpdate: 'RESTRICT'
    },
    redirect_uris: {
      type: 'text[]',
      notNull: true
    },
    response_types: {
      type: 'text[]',
      notNull: true
    },
    grant_types: {
      type: 'text[]',
      notNull: true
    }
  });
};

exports.down = (pgm) => {
  pgm.dropTable('oauth2_clients');
};
