const frisby = require('frisby');

beforeAll(() => {
    frisby.baseUrl(process.env.SERVICE_URI);
});