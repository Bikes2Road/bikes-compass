# bikes-compass
This repositorie is to search bikes on db

## Description

This microservice provides an API for querying bikes in a MongoDB database. It allows you to retrieve bikes information, filter by name, perform paginated searches, and exposes endpoints with clear and documented errors for easier consumption.

## Getting Started

### Prerequisites

- [Go 1.21+](https://golang.org/dl/)
- Access to a MongoDB instance

### Installation and Execution

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Bikes2Road/bikes-compass.git
   cd bikes-compass
   ```

2. **Configure environment variables:**

   Create a `.env` file with your MongoDB settings, for example:
   ```
   MONGODB_URI=mongodb://localhost:27017
   MONGODB_DATABASE=bikesdb
   MONGODB_COLLECTION=bikes
   ```

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run the application:**

   ```bash
   go run main.go
   ```

   The service will typically run on port `8080` or the port specified in your environment variables.

5. **Test the main endpoint:**

   You can make a sample request using curl:
   ```bash
   curl "http://localhost:8080/v1/bikes/search?page=1&cant=10&name=Yamaha"
   ```

### Swagger Documentation

Interactive Swagger documentation is available for testing the endpoints.

1. **Access Swagger UI:**

   Once the service is running, go to:
   ```
   http://localhost:8080/swagger/index.html
   ```
   or
   ```
   http://localhost:8080/swagger/
   ```
   depending on your router configuration.

2. **There you'll find documentation for all endpoints, input parameters, possible responses, and returned status codes.**

---

### Useful Commands

- **Start the service:**  
  ```cd cmd/api && go run main.go```
- **Update dependencies:**  
  ```go mod tidy```

---

## Main Endpoints

- `GET /v1/bikes/search`  
  Searches motorcycles in the database, optionally filtering by name, and using pagination (`page`, `cant`).

---

## Notes

- The project follows best practices for hexagonal architecture (ports and adapters).
- Error messages are standardized.
- Use the correct values for `page` (greater than or equal to 1) and `cant` (maximum 30).
- Name searches only accept letters and spaces.
