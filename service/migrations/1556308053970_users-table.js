exports.shorthands = undefined;

exports.up = (pgm) => {
    pgm.createTable('users', {
        user_id: {
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
        password: {
            type: 'text',
            notNull: false
        },
        email: {
            type: 'text',
            notNull: false
        }
    });
};

exports.down = (pgm) => {
    pgm.dropTable('users');
};
