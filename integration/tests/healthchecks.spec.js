const frisby = require('frisby');

describe('Healthchecks', () => {
    it('Passes', () => {
        return frisby.get('/health')
            .expect('status', 200)
            .expect('jsonStrict', {
                "result": "OK",
                "components": {
                    "database": {
                        "result": "OK",
                        "detail": "Database OK"
                    }
                }
            });
    });
});
