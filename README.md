# Rest api - register and login user

## required
- Go version 1.18++
- Mysql
- Gin
- JWT

## Operation
- First go mod tidy
- Fill the .env same with your database
- make run -> running the project

REST-API-Users
- └── config          # load any config at .env.example
- └── entity          # seems like model
- └── transport       # standarization request from client and response data
- └── helper         # response general, success and error
- └── server          # routing and main directory
