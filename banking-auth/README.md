# banking-auth

Make sure to define the following env variables to start the server

- SERVER_ADDRESS    `[IP Address of the machine]`
- SERVER_PORT       `[Port of the machine]`
- DB_USER           `[Database username]`
- DB_PASSWD         `[Database password]`
- DB_ADDR           `[IP address of the database]`
- DB_PORT           `[Port of the database]`
- DB_NAME           `[Name of the database]`

To run locally: 
SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD=root DB_ADDR=localhost DB_PORT=3307 DB_NAME=banking go run main.go
