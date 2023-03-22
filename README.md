# GoMongo


Endpoints

GET http://localhost:4000/api/v1/user get all users

POST http://localhost:4000/api/v1/user create User

GET http://localhost:4000/api/v1/user/userId get user by id

PATCH http://localhost:4000/api/v1/user/userId update user

DELETE http://localhost:4000/api/v1/user/userId delete user


RUN 
$ docker-compose -f docker-compose.yaml up \n
<b>OR</b> \n 
$ go build main.go
$ ./main
