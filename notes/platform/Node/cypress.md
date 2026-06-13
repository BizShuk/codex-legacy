# Cypress note

For Cypress to prevent flake

### Core

Cypress is built using Promises from Bluebird, doesn't return typical Promise instances. What it returns is called `Chainer`.

Queueing command style

### Common usage

cy.visit('url')

cy.get('input.post-title').type('My First Post') // .type => input text

cy.contains('Submit').click()

cy.url().should('include', '/posts/my-first-post')

cy.get('h1').should('contain', 'My First Post')

cy.get('.error').should('be.empty') // Assert that '.error' is empty
cy.contains('Login').should('be.visible') // Assert that el is visible
cy.wrap({ foo: 'bar' }).its('foo').should('eq', 'bar') // Assert the 'foo' property equals 'bar'

### select

cy.get('#main-content').find('.article').children('img[src^="/static"]').first()
Best Practice: Use data-* attributes to provide context to your selectors and isolate them from CSS or JS changes.

### timeout

default 4 sec

#### sync

// Cypress.$ is synchronous, so evaluates immediately
  // there is no element to find yet because
  // the cy.visit() was only queued to visit
  // and did not actually visit the application
  let el = Cypress.$('.new-el') // evaluates immediately as []

  When using a callback function with .should(), be sure that the entire function can be executed multiple times without side effects. Cypress applies its retry logic to these functions: if there's a failure, it will repeatedly rerun the assertions until the timeout is reached. That means your code should be retry-safe. The technical term for this means your code must be idempotent.

### file/folder organization

```file
cypress
    /fixture
        - example.json
    /intergration
        - *.spec.js
    /plugins    // execute before browser launch
        - index.js
    /support    
        - index.js // automatically included run ince before all specs

// generate during test
    /downloads
    /screenshots
    /videos
```

Cypress create

- screenshotsFolder
- videosFolder

cypress.json
cypress.env.json

DEBUG=cypress:server:specs npx cypress open

## or

DEBUG=cypress:server:specs npx cypress run

<https://docs.cypress.io/guides/core-concepts/variables-and-aliases#Avoiding-the-use-of-this>
Arror function this won't work with?
