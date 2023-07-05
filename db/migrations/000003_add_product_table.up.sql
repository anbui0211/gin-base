CREATE TABLE products (
    id VARCHAR(255) PRIMARY KEY NOT NULL ,
    name VARCHAR(255) NOT NULL,
    search_string VARCHAR(255) NOT NULL,
    category_id VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL NOT NULL,
    status VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);