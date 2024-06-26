# Go architecture

# Api 

`Get all users`

GET  http://localhost:8080/users

`Create user`

POST http://localhost:8080/users

body

{
    "name": "ansar",
    "age": 26,
    "password": "ansar26"
}

`Update user`

PUT http://localhost:8080/users/:id

header

Authorization Bearer YW5zYXI6YW5zYXIyNg==

body 

{
    "name": "ansar",
    "age": 27,
    "password": "ansar26"
}

`Delete user`

DELETE http://localhost:8080/users/:id

header

Authorization Bearer YW5zYXI6YW5zYXIyNg==