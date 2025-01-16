```markdown
# Go Lang JWT Example Project

This project demonstrates how to implement a secure authentication system using JWT (JSON Web Tokens) in Go (Golang). It uses the **Fiber** web framework for fast HTTP handling, **MongoDB** for user data storage, and **bcrypt** for secure password hashing. With Docker Compose, this project can be easily set up and run in a containerized environment.

## Features

- **JWT Authentication**: Secure user authentication using JWT tokens.
- **Password Hashing**: User passwords are hashed using bcrypt for security.
- **MongoDB Integration**: MongoDB is used to store user data and authentication-related information.
- **Protected Routes**: Certain API endpoints are protected and require a valid JWT token for access.
- **Dockerized**: The project comes with a Docker Compose configuration to easily set up and manage the required services (MongoDB, Go backend, and Nginx proxy).

## Technologies Used

- **Go (Golang)**: Backend programming language.
- **Fiber**: A fast and lightweight web framework for Go.
- **JWT**: Token-based authentication mechanism for secure access.
- **MongoDB**: Database for storing user data and authentication information.
- **bcrypt**: Secure password hashing library.
- **Docker Compose**: Tool for defining and running multi-container Docker applications.

## Project Setup

This project uses **Docker Compose** for seamless setup. It includes the following services:

- **MongoDB**: A container running MongoDB to store authentication-related data.
- **Auth Service**: The Go backend service, which handles JWT authentication and user management.
- **Auth Proxy**: An Nginx reverse proxy for handling requests and routing them to the auth service.

### Running the Project with Docker

To get started, you can run the project with Docker Compose. The steps are simple and quick:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/hamzapro305/GoJwt.git
   cd GoJwt
   ```

2. **Build and start the services**:
   ```bash
   docker-compose up --build
   ```

3. **Access the API**:
   - The API will be available at `http://localhost:3000`.
   - Available routes include:
     - **POST** `/api/v1/auth/register`: Register a new user.
     - **POST** `/api/v1/auth/login`: Log in and receive a JWT token.
     - **GET** `/api/v1/protected/test`: Access a protected route (requires a valid JWT token).

### Docker Compose Overview

The **docker-compose.yml** file sets up the following services:

- **MongoDB**: The database for storing user authentication details.
- **Auth Service**: The core backend service written in Go, which interacts with MongoDB and handles authentication-related operations.
- **Auth Proxy**: A reverse proxy using Nginx that ensures secure routing and access to the auth service.

## Key Endpoints

- **/api/v1/auth/register** (POST): Register a new user.
- **/api/v1/auth/login** (POST): Log in to receive a JWT token.
- **/api/v1/protected/test** (GET): A protected route that requires a valid JWT token for access.

## Environment Variables

To ensure the correct configuration, the following environment variables are required:

- `MONGO_INITDB_ROOT_USERNAME`: Root username for MongoDB.
- `MONGO_INITDB_ROOT_PASSWORD`: Root password for MongoDB.
- `MONGO_INITDB_DATABASE`: Name of the MongoDB database.
- `MONGODB_URI`: URI to connect to MongoDB from the Go service.

These variables can be set in the **docker-compose.yml** file or passed directly in the environment.

## Dependencies

The project relies on several Go dependencies, including:

- **Fiber**: For routing and handling HTTP requests.
- **JWT**: For generating and validating JWT tokens.
- **MongoDB**: For storing user data.
- **bcrypt**: For securely hashing passwords.

The full list of dependencies is specified in the `go.mod` file, ensuring that all necessary libraries are included for the project to run smoothly.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

For further questions or contributions, feel free to open an issue or pull request on the repository.
```