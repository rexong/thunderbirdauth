# Thunderbirdauth

## What?

This is a light weight IDP built using go and sqlite3.
Thunderbirdauth is meant to be used with a reverse proxy like Nginx.
Application that requires authentication would be routed to
Thunderbirdauth by the reverse proxy to be veified.

## Current Feature

- Register User
- Verifiy User
  - Create Session Token as Cookie

## Authentication Flow/Protocol Supported

- No Flow/Protocol
  - Allows Thunderbirdauth to handle authentication on its own.
  - Trust the Session Token set by Thunderbirdauth
- Basic Authentication Flow
  - Allows Application that supports Basic Auth
