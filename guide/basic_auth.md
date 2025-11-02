# Basic Authentication

HTTP provides simple challenge-response authentication framework.

## Challenge-Response Flow

1. Client access protected resources without authorization header
2. Server responds to client with 401 and a `WWW-Authenticate` response header
    - `WWW-Authenticate: <type> realm=<realm>`
    - type is `Basic`, realm describes the protected area.
3. Client authenticate with server with credentials
    - With `Authentication: <type> <credentials>` request header
    - Usually base64 encoded `username:password` string used as credentials
4. Server response 200 if valid otherwise 401
