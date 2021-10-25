-- liquibase formatted sql

-- changeset zinov:2021-10-25-14-06-create-users
CREATE TABLE IF NOT EXISTS users
(
    id          serial          NOT NULL,
    email       varchar(255)    NOT NULL,
    password    varchar(255)    NOT NULL,
    PRIMARY KEY (id)
    );
-- rollback drop table users;