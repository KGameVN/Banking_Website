--USER
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    logintoken VARCHAR(255),
    expiredtime VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
--Amount
CREATE TABLE user_accounts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    account_number VARCHAR(20) UNIQUE NOT NULL,
    balance NUMERIC(18,2) DEFAULT 0.00,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Profile
CREATE TYPE gender_type as ENUM ('male', 'female');
CREATE TABLE user_profile (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    firstname VARCHAR(20),
    lastname VARCHAR(20),
    cmnd VARCHAR(20),
    address VARCHAR(20),
    gender  gender_type,
    birthday VARCHAR(20)
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
-- default value
INSERT INTO users (
    username,
    email,
    password,
    logintoken
) VALUES (
    'User1',
    'user@test.com',
    '1234',
    'token_123456'
), (
    'User2',
    'user2@test.com',
    '1111',
    'token123'
);

INSERT INTO user_accounts (
    user_id,
    account_number,
    balance
) VALUES (
    1,
    1,
    100
),(
    2,
    200,
    1000
);

--Transaction
-- SQL tạo bảng `transactions` trong PostgreSQL
CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,      -- ID tự tăng (mặc định cho `bigserial`)
    "TransactionTime" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Thời gian giao dịch
    "From" BIGINT NOT NULL,        -- ID tài khoản người gửi
    "To" BIGINT NOT NULL,          -- ID tài khoản người nhận
    "Amount" BIGINT NOT NULL,      -- Số tiền giao dịch
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,     -- Thời gian tạo giao dịch
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,     -- Thời gian cập nhật giao dịch
    CONSTRAINT fk_from_account FOREIGN KEY ("From") REFERENCES user_accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_to_account FOREIGN KEY ("To") REFERENCES user_accounts(id) ON DELETE CASCADE
);

--Token
CREATE TABLE login_tokens (
    id SERIAL PRIMARY KEY,                  -- Unique identifier for each login token
    token VARCHAR(255) NOT NULL,            -- JWT token (non-empty)
    expiredtime TIMESTAMP,                  -- Expiry time for the token (optional)
    user_id INT NOT NULL,                   -- Reference to the user who owns this token
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- The time when the token was created
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) -- Foreign key constraint to the users table
);

-- You may want to add an index for the `user_id` field for fast lookups
CREATE INDEX idx_user_id ON login_tokens(user_id);


