create table if not exists admins
(
    id           int generated always as identity primary key,
    name         text not null,
    login        text unique not null,
    password     text not null
);

create table if not exists leaders
(
    id           int generated always as identity primary key,
    name         text not null,
    phone_number text not null,
    login        text unique not null,
    password     text not null
);

create table if not exists members
(
    id           int generated always as identity primary key,
    name         text not null,
    phone_number text not null,
    role         text not null,
    login        text unique not null,
    password     text not null
);