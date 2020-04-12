create table if not exists users (
  id SERIAL NOT NULL,
  user_id text UNIQUE,
  password text,
  created_at timestamp with time zone not null default current_timestamp,
  updated_at timestamp with time zone not null default current_timestamp,
  PRIMARY KEY (id)
);
create table if not exists mikans (
  id SERIAL NOT NULL,
  name varchar(255),
  kind varchar(255),
  quality integer not null,
  created_at timestamp with time zone not null default current_timestamp,
  updated_at timestamp with time zone not null default current_timestamp,
  PRIMARY KEY (id)
);