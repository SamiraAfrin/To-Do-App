# Project Description
It is a todo app built using Golang. CRUD operations are implemented following clean-architecture. Here, echo framework and mysql is used for the execution of the app. 

# Layers
This project has 4 layers :
* Models Layer
* Repository Layer
* Usecase Layer
* Delivery Layer

# How to run the project
#### Here is the steps to run it with ```docker-compose```

```
#move to directory
$ cd workspace

# Clone into your workspace
$ git clone https://github.com/SamiraAfrin/To-Do-App.git

#move to project
$ cd To-Do-App

# Run the application
$ docker compose up -d mysql adminer
$ docker compose up web - - build
```
# Database
### If the database is empty, database can be updated using two ways

## Way 1
```
From the adminer GUI
# In the browser, type 
- localhost:8080
# Fill the fields using the following credentials
- System --> MySQL
- Server --> mysql
- Username --> root
- Password --> 123
- Database --> recordings
# Update the database
- SQL commands are provided in the database.sql file
```
## Way 2
```
From the container terminal
# To be in the container terminal
$ docker exec -it mysql container id bash
$ To connect mysql server
$ mysql -u root -h localhost --protocol tcp -P 3306 -p
$ Passwords ? --> 123
# Update the database
- SQL commands are provided in the database.sql file
```
# API
## Way 1
```
#move to project
$ cd To-Do-App

#Execute the call in another terminal
$ curl localhost:8000/tasks
```
## Way 2
```
# Postman might be a good option 
# choose the http method
- Example: Choose GET method 
# In the url section,type, localhost:8000/endpoint
- Example: localhost:8000/tasks
# All the api patterns are mentioned in the following path
- To-Do-App/Task/delivery/http/task_handler.go
```
# Unit Test
### To create mocks file, use mockery tools
[Mockery Site](https://vektra.github.io/mockery/)

```
# To run all the test files from the roor directory
$ go test -v ./...
# to run specific test files,specify the exact location
$ go test -v ./Task/usecase/ 
```