CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    username VARCHAR(20) UNIQUE,
    email VARCHAR(42) UNIQUE,
    name VARCHAR(100),
    password VARCHAR(100),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    enabled_at TIMESTAMP,
    disabled_at TIMESTAMP
);
