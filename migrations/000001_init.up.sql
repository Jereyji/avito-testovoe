CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    balance INT NOT NULL DEFAULT 1000
);

CREATE TABLE merch (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    price INT NOT NULL CHECK (price > 0)
);

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sender_id UUID NOT NULL,
    recepient_id UUID NOT NULL,
    amount INT NOT NULL CHECK (amount > 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (recepient_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE purchases (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    merch_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (merch_id) REFERENCES merch(id) ON DELETE CASCADE
);