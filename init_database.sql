CREATE DATABASE IF NOT EXISTS moveinsync;

USE moveinsync;

CREATE TABLE IF NOT EXISTS users (
    id varchar(100) PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);
