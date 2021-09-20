# Twitter API Integration - Demo Application
In order to run this application, you'll need:
### Credentials from twitter dev:
Create a dev account on twitter-dev, and allow creating tweets. Generate the API Key, Access tokens and secrets.
```
CONSUMER_KEY=
CONSUMER_SECRET=
ACCESS_TOKEN=
TOKEN_SCERET=
```

Running this applicaiton would need 2 terminals.

### Service 
on the first terminal,
`cd twitter-service`
`go build`
`./twitter-service`
This will run the golang service on port 8080

### Web
on the second terminal,
`cd twitter-web`
`npm i && npm run start`
This will run the React web applicaiton on port 3000.
