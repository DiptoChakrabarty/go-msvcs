# go-msvcs
Learn about microservices in go

### Services

- User Service
    * User DataBase Operations
    * User Creation,Deletion,Updation

Example Request
```sh
curl -X POST -H "Content-Type: application/json" -d '{"first_name": "pop", "last_name": "dop", "email": "dld@gmail.com"}' http://localhost:5000/users
```