# To-Do-App
It is a todo app built using Golang. CRUD operations are implemented following clean-architecture. Here, echo framework and mysql is used for the execution of the app. 

# Layers
This project has 4 layers :
* Models Layer
* Repository Layer
* Usecase Layer
* Delivery Layer

# How to run the projects
#### Here is the steps to run it with docker-compose

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
### If the database is empty, database can be updated, using to ways

## Way 1
## Way 2
