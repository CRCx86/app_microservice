-- liquibase formatted sql

-- changeset zinov:2021-10-24-14-53-create-temps
CREATE TABLE IF NOT EXISTS temps
(
    id          serial      NOT NULL,
    number      varchar(15) NOT NULL,
    brand       varchar(20) NOT NULL,
    PRIMARY KEY (id)
);
-- rollback drop table temps;