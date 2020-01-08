#CRUD APIs using MYSQL-GOLANG

BASIC APIs created to login/register/validate token/get user list/delete/update set of users

## Starting the GO server

Below command will start the GO server.

```bash
go run *.go
```
## MYSQL 

create below tables in mysql:-

```
CREATE TABLE login_tokens(
 username varchar(50) PRIMARY KEY NOT NULL,
 token varchar(120) DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=latin1;
```
```
CREATE TABLE users (
  username varchar(50) PRIMARY KEY NOT NULL,
  first_name varchar(200) NOT NULL,
  last_name varchar(200) NOT NULL,
  password varchar(120) DEFAULT NULL,
 token varchar(120) DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=latin1;
```