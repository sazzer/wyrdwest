it('is healthy', () => {
  const url: string = Cypress.env('SERVICE_URI');

  cy.request(url + '/health').then(response => {
    expect(response.status).to.eq(200);
    expect(response.body).to.deep.equal({
      status: 'OK',
      details: {
        database: {
          status: 'OK'
        }
      }
    });
  });
});
