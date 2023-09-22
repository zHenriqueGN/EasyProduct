# EasyProduct
A simple API to manage products

## Running API:
### Configuring the environment
- create a .env file based on conf struct (configs/configs.go)

### Running tests
- configure the .env file to point to the test database configured on  test/test_database/docker-compose.yaml
- get up the test database container with docker compose:
  - "docker compose up -d"
- run all tests:
  - "go test ./..."
- you can test the API points by using the API calls located in test/http_calls or test using the swagger as well

### Accessing the API
- Go to http://localhost:8000/docs/ on your browser and you be able to access the API docs.
