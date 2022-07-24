
-- +migrate Up
CREATE TABLE users
(
    id bigserial not null primary key,
    email varchar(255) not null unique,
    password varchar(255) not null
);

-- +migrate Down
DROP TABLE users;