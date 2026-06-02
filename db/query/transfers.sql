-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount)
 VALUES ($1, $2, $3)
 RETURNING *;


-- name: GetTransferByFromAccountID :many
SELECT * FROM transfers
 WHERE from_account_id = $1
 ORDER BY id
 LIMIT $2 OFFSET $3;

-- name: GetTransferByToAccountID :many
SELECT * FROM transfers
 WHERE to_account_id = $1
 ORDER BY id
 LIMIT $2 OFFSET $3;

-- name: ListTransfersBetweenAccounts :many
SELECT * FROM transfers 
WHERE (from_account_id = $1 AND to_account_id = $2) 
   OR (from_account_id = $2 AND to_account_id = $1)
ORDER BY id
LIMIT $3 OFFSET $4;