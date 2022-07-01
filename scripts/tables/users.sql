DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id serial primary key,
    name varchar(255),
    phone varchar(15),
    mail varchar(255)
);
