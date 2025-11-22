# Jellyfin LDAP-AUTH Setup

1. Install Jellyfin's LDAP-Auth plugin

- Login to jellyfin as admin user
- Navigate `Dashboard` > `Plugins` > `All` > `Search` LDAP Authentication
- Install LDAP Authentication
- Restart Jellyfin Application

1. Configure LDAP-Auth plugin

- Navigate `Dashboard` > `Plugins` > `Installed` > `LDAP-Authentication` > `Settings`
- Fill in the form:
  - LDAP Server: Docker Container Name (if in same dockntwrk)/
  local machine ip addr
  - LDAP Port
  - Secure LDP: Unchecked
  - LDAP Bind User:
    - `your-bind-user-dn`
    - Example: `cn=admin,dc=example,dc=com`
  - LDAP Bind User Password:
    - `your-bind-user-password`
    - Example: `adminpassword`
  - LDAP Base DN for searches:
    - `your-base-dn`
    - Example: `dc=example,dc=com`
  - Save and Test LDAP Server Settings
  - LDAP Search Filter: `(&(objectClass=person)(uid=myuser))`
  - Save and Test LDAP Filter Settings
  - LDAP Search Filter: Leave Empty
  - LDAP Search Attributes: `uid`
  - LDAP Uid Attributes: `uid`
  - LDAP Username Attributes: `uid`
  - LDAP Password Attributes: `userPassword`
