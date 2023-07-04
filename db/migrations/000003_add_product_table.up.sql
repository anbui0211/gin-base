CREATE TABLE products (
    id VARCHAR(255) PRIMARY KEY NOT NULL ,
    product_id VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    search_string VARCHAR(100) NOT NULL,
    category_id INT NOT NULL,
    quantity INT NOT NULL,
    price MONEY NOT NULL,
    status VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);