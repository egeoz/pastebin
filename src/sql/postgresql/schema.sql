create table entries(
  id bigserial primary key,
  uuid text not null unique,
  title text not null,
  content text not null,
  content_type text not null,
  is_encrypted text not null,
  insert_date datetime not null
  );