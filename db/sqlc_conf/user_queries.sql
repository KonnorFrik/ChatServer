-- name: CreateUser :one
INSERT INTO users (
    name, email, password, role
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUserName :exec
UPDATE users
    SET name = $2
    WHERE id = $1;

-- name: UpdateUserEmail :exec
UPDATE users
    SET email = $2
    WHERE id = $1;

-- name: UpdateUserNameEmail :exec
UPDATE users
    SET name = $2,
        email = $3
    WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
    WHERE id = $1;

