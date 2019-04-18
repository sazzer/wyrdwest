const frisby = require('frisby');

describe('Healthchecks', () => {
    it('Passes', () => {
        return frisby.get('/health')
            .expect('status', 200)
            .expect('jsonStrict', {
                "result": "OK",
                "components": {
                    "passing": {
                        "result": "OK",
                        "detail": "It Failed"
                    }
                }
            });
    });
});