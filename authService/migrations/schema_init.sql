CREATE TYPE USER_TYPE as ENUM ('USER', 'PREMIUM_USER', 'ADMIN');

CREATE TABLE "user"
(
    id       BIGSERIAL PRIMARY KEY,
    login      varchar(40) not null unique,
    password   varchar(50) not null,
    first_name varchar(40) default null,
    last_name  varchar(40) default null,
    phone      varchar(40) default null,
    type       USER_TYPE default 'USER'
);