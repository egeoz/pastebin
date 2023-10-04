-- name: GetEntry :one
select * from entries
where uuid = ? limit 1;

-- name: CreateEntry :one
insert into entries (
  uuid, title, content, content_type, is_encrypted, insert_date
) VALUES (
  ?, ?, ?, ?, ?, ?
)
returning *;

-- name: DeleteEntry :exec
delete from entries
where uuid = ?;