-- name: GetEntry :one
select * from entries
where uuid = $1 limit 1;

-- name: CreateEntry :one
insert into entries (
  uuid, title, content, content_type, is_encrypted, insert_date
) VALUES (
  $1, $2, $3, $4, $5, $6
)
returning *;

-- name: DeleteEntry :exec
delete from entries
where uuid = $1;