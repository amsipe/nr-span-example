create schema if not exists example collate latin1_swedish_ci;

use example;

create table if not exists users
(
    id int unsigned auto_increment
        primary key,
    name varchar(255)
);

insert into users (name) values ('andy'), ('tom'), ('scott'), ('wallace');