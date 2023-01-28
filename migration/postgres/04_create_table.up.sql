CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    login VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    typeU VARCHAR,
    phone_number VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);