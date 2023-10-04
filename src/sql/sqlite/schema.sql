create table entries(
  id integer primary key asc,
  uuid text not null unique,
  title text not null,
  content text not null,
  content_type text not null,
  is_encrypted text not null,
  insert_date text not null
  );