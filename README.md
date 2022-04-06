# go-msvcs
Learn about microservices in go

### Services

- User Service
    * User DataBase Operations
    * User Creation,Deletion,Updation

Example Request
```sh
curl -X POST -H "Content-Type: application/json" -d '{"first": "pop", "last": "dop", "email": "dld@gmail.com"}' http://localhost:5000/users
```
