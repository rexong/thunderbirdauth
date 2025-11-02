# Features

## Register

To register a user, send credentials to `/register` endpoint in JSON format.
There is no GUI currently to register user using the `/register` endpoint.
Hence, to register a user, we would restort to the `curl` command.

```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{"username": "bob", "password": "1234"}'
```

## Login

To login, go to `/login` endpoint in the browser.
*this endpoint requires a `redirect url`*
A form will appear, where you can enter the user's credentials.

### Login Flow

1. User enter credentials, username and password
2. If credentials or `redirect url` invalid, error will be shown
3. Otherwise, create session token and store it as cookie
4. Then, redirect to the url provided as `redirect url`

## Auth

The `/auth` endpoint serves as a challenge-response authentication.
There is another endpont `/auth/basic` which is similar to `/auth`,
but writes a Authorization request header that statisfy basic authentication.

### Auth Flow

1. When invoked, it checkes if session token exist and valid.
2. If not, return status code 401.
3. Otherwise, if `/auth/basic` then add Authorization header,
return status code 200.
