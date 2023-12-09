CREATE TABLE category
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(100)
);


CREATE TABLE product
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(100),
    category_id BIGINT references category(id),
    image       VARCHAR(100),
    price       NUMERIC(16, 2)
);

CREATE TABLE user_favorite
(
    user_id    BIGINT,
    product_id BIGINT
);

CREATE TABLE user_cart
(
    user_id    BIGINT,
    product_id BIGINT
);