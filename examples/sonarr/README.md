# Sonarr

## Setting up Sonarr with Basic Auth

When the Sonarr container is newly created, accessing the sonarr page will prompt
for Authentication Method, Username and Password.
Choose Basic Authentication, and ensure that your username and password matches the
`BASIC_USERNAME` and `BASIC_PASSWORD` in the `.env` file.

## Setting up Sonarr with Nginx

For Nginx to reach Sonarr, the URL Base needs to be updated.
Access the Sonarr Webpage, go to `settings` > `general` > `URL Base`.
Add `sonarr` to the input field and save the changes.
