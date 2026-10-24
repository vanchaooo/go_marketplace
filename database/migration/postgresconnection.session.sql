CREATE TABLE users 
(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(200) NOT NULL CHECK (char_length(name) > 0),
    age SMALLINT NOT NULL CHECK (age > 0 and age < 150), 
    email VARCHAR NOT NULL CHECK (char_length(email) > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);