CREATE TABLE products 
(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name varchar(200) not null check(char_length(name) > 0),
    price bigint not null,
    stock bigint not null
);