## Mini-Project Golang Celerates CAP

1 Domain (vehicle), 1 domain for Login and Register (users).
Creating a 2 domain REST API with GIN, GORM, and PostgreSQL

## Inside Mini Project
- CRUD Vehicles
- login Member
- Register Member
- JWT Token
- Hexagonal Architecture

## PostgreSQL
One of The SQL used in this project
### users
```sh
CREATE TABLE "users" (
  "username" varchar(20) PRIMARY KEY NOT NULL,
  "password" varchar(255) NOT NULL,
  "role" varchar(20) NOT NULL,
  "created_on" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
```
### vehicles
```sh
CREATE TABLE "vehicles" (
  "vehicle_id" serial PRIMARY KEY NOT NULL,
  "name" varchar(100) NOT NULL,
  "type" varchar(20) NOT NULL,
  "plat" varchar(20) NOT NULL,
  "color" varchar(100) NOT NULL
);
```
