# Internet Shop v1

### My First Big Project with Microservices Architecture

This project utilizes gRPC and REST API methods for authorization and authentication using JWT. The database used is PostgreSQL.

![Schema Project](SchemaProject.png)

---

### Description of Services

#### ClientService
All endpoints are integrated into this service, including the *REST API service authorization*.

**Endpoints:**

- **/registration**  
  Users can register in the database. After registration, their username, login, and hashed password are stored in the *UsersDB*.

- **/login**  
  This endpoint authenticates users. If the login is correct, the service returns a JWT token containing the user's ID and role (default: consumer).

- **/catalog**  
  JWT is required to access this page. The endpoint has two methods:
  - **GetItemsFromDB**: Requests data from the *gRPC CatalogService*. If the response is valid, all items from the database are displayed.
  - **AddItemToWishlist**: Sends the item ID to the *gRPC WishlistService*. If successful, a boolean response is returned.

- **/catalog/adder**  
  JWT authentication is required, and the role must be "admin"; otherwise, a StatusCode 403 is returned. The endpoint has two methods:
  - **AddItem**: Sends item information to the *gRPC AdderCatalogService*. If successful, a boolean response is returned.
  - **DeleteItem**: Sends the item ID to the *gRPC AdderCatalogService*. If successful, a boolean response is returned.

- **/admin/adder**  
  JWT authentication is required, and the role must be "admin"; otherwise, a StatusCode 403 is returned. The endpoint has two methods:
  - **AddAdmin**: Sends a request with the user ID to promote the user to "admin" using the *gRPC AdderAdminService*. A true response indicates success; otherwise, it returns false.
  - **DeleteAdmin**: Sends a request with the user ID to downgrade the user using the *gRPC AdderAdminService*. A true response indicates success; otherwise, it returns false.

- **/catalog/wishlist**  
  JWT is required to access this page. The endpoint has two methods:
  - **DeleteFromWishlist**: Sends the item ID to the *gRPC WishlistService* for deletion. The response will be true or false.
  - **GetWishlist**: Requests data from the *gRPC WishlistService*.

---

### Additional Services

#### AuthService (REST API)
This service handles requests from clients. If the endpoint is `/registration`, it checks that the username and login are not already in the *UsersDB*. If no errors are found, the password is hashed and stored along with the username and login.  
For the `/login` endpoint, the service checks if the login exists and compares the submitted password with the hashed password. If verified, a JWT is generated with the member's ID and role, and the server saves this JWT in its struct.

#### CatalogService (gRPC)
This service processes requests from clients:
- **GetItemsFromDB**: Returns a list of item structs.
- **AddItemToWishlist**: Adds the item to the wishlist by calling the `AddToWishlist` method from the *WishlistService*.

#### AdderCatalogService (gRPC)
This service manages item-related requests:
- **AddItem**: Receives item information and adds it to the *ItemsDB*.
- **DeleteItem**: Receives an item ID and deletes it from the *ItemsDB*. The response will be true or false.

#### AdderAdminService (gRPC)
This service manages admin-related requests:
- **AddAdmin**: Receives a request to promote a user to "admin". The service sends this request to the *UsersDB* and returns true or false.
- **DeleteAdmin**: Receives a request to downgrade a user to "consumer". The service sends this request to the *UsersDB* and returns true or false.

#### WishlistService (gRPC)
This service handles wishlist-related requests:
- **GetWishlist**: Returns structs about followed items from the *WishlistDB*.
- **DeleteFromWishlist**: Deletes an item from the wishlist and returns true or false.
- **AddItemToWishlist**: Receives an item ID, fetches data about the item from the *ItemsDB*, and adds it to the *WishlistDB*. The response will be true or false.

---

**P.S.** I hope that readers will understand at least the scheme!  
Also, I started learning Golang 1 month and 6 days ago.
