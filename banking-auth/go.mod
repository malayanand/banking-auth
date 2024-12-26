module github.com/malayanand/banking-auth

go 1.21.1

require (
	github.com/gorilla/mux v1.8.1
	github.com/jmoiron/sqlx v1.4.0
	github.com/malayanand/banking v0.0.0
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
)

replace github.com/malayanand/banking => ../banking/
