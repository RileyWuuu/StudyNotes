# Third-party access (Oauth2)

- Provides consented access and restricts actions of what the client app can perform on resources on behalf of the user, without sharing user's credentials.
- An authorization protocol (as a means of granting access to a set of resources)
- Uses Access Tokens (doesn’t define a specific format)

### Roles:
- Resource Owner : user or system that owns the protected resources and can grant access to them.
- Client : system that requires access to the protected resources. 
- Authorization Server : receives requests from Client for Access Tokens and issues them upon successful authentication and consent by the Resource Owner.
- Resource Server : server that protects the user’s resources and receives access requests from the Client.

![$60 00 $65 00 $80 00](https://user-images.githubusercontent.com/71340325/191670853-27ebca76-4d8a-468f-81b1-83a35781b54c.jpg)
### As shown on the picture
1. Client requests authorization from Authorization server
2. Authorization server authenticates the client and verifies
3. Resource owner interacts with Authorization server to grant access
4. Authorization server redirects back to client with code/token


### Grand Types:
- Authorization Code : The Authorization server returns a single-use Authorization Code to the Client, which is then exchanged for an Access Token.
- Implicit Grant : A simplified flow where the Access Token is returned directly to the Client.
