Example commands
Start the containers: docker-compose up -d

Rebuild and start: docker-compose up -d --build

Login: docker exec -it postgres psql -U user

This command will be determined by the docker-compose.yml file if different env variables are used.

Eg. docker exec -it postgres psql -U user testdb as this takes a database name.

Show tables: \dt

Show databases: \l


# commands 

get list of databases
\list 

connect to "demo" database
\c demo

show tables
\dt

CREATE DATABASE testdb;

CREATE TABLE POSTS (ID INT PRIMARY KEY NOT NULL, TITLE text, BODY text);

INSERT INTO POSTS(ID, TITLE) VALUES(1, 'How to create a table in postgreSQL');

INSERT INTO POSTS VALUES(1, 'How to create a table in postgreSQL');

SELECT ID, TITLE FROM POSTS;

SELECT * FROM POSTS;

