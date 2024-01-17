CREATE TABLE users
( 
    id SERIAL NOT NULL UNIQUE, 
    name VARCHAR(255) not null,
    surname VARCHAR(255) not null,
    patronymic VARCHAR(255),
    age INT,
    gender VARCHAR,
    country VARCHAR(255)
);