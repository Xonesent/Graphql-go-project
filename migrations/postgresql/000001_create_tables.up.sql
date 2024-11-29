CREATE SCHEMA IF NOT EXISTS shop;

CREATE TABLE IF NOT EXISTS shop.users (
    user_id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE IF NOT EXISTS shop.products (
    product_id SERIAL PRIMARY KEY,
    item TEXT,
    attributes_id TEXT
);

CREATE TABLE IF NOT EXISTS shop.orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES shop.users(user_id),
    product_id INT REFERENCES shop.products(product_id),
    price INT
);