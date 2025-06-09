CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password TEXT NOT NULL,
    role INTEGER CHECK ( role <> 0 ),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
