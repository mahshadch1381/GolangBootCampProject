CREATE TABLE voucher (
    id SERIAL PRIMARY KEY,
    number VARCHAR(64) NOT NULL UNIQUE,
    version INT NOT NULL DEFAULT 1
);