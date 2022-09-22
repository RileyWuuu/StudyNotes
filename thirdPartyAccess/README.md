# Third-party access (Oauth2)

- Provides consented access and restricts actions of what the client app can perform on resources on behalf of the user, without sharing user's credentials.
- An authorization protocol (as a means of granting access to a set of resources)
- Uses Access Tokens (doesn’t define a specific format)

### Roles:
- Resource Owner : user or system that owns the protected resources and can grant access to them.
- Client : system that requires access to the protected resources. 
- Authorization Server : receives requests from Client for Access Tokens and issues them upon successful authentication and consent by the Resource Owner.
- Resource Server : server that protects the user’s resources and receives access requests from the Client.

