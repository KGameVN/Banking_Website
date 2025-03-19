--USER
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    fullname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    logintoken VARCHAR(255),
    expiredtime VARCHAR(255),
    birthday TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Trigger to automatically update updated_at on each update
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE INDEX idx_users_email ON users (email);

--Transaction
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    transactiontime TIMESTAMP NOT NULL,
    from BIGINT NOT NULL,
    to BIGINT NOT NULL,
    amount BIGINT NOT NULL
);
