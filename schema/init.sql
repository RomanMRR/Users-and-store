CREATE TABLE IF NOT EXISTS users  (
    id bigserial not null primary key,
    name varchar not null,
    surname varchar not null,
    patronymic varchar not null,
    age smallint,
    registrations_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS shops  (
    id bigserial not null primary key,
    name varchar not null,
    address varchar not null,
    working bool not null,
    owner varchar
);