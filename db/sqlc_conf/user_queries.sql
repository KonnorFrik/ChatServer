-- name: CreateUser :one
INSERT INTO users (
    name, email, password, role, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, NOW(), NOW()
)
RETURNING *;

-- name: UserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUserName :exec
UPDATE users
    SET name = $2,
        updated_at = NOW()
    WHERE id = $1;

-- name: UpdateUserEmail :exec
UPDATE users
    SET email = $2,
        updated_at = NOW()
    WHERE id = $1;

-- name: UpdateUserNameEmail :exec
UPDATE users
    SET name = $2,
        email = $3,
        updated_at = NOW()
    WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
    WHERE id = $1;

