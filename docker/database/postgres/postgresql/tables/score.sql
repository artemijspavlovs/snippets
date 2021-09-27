BEGIN
    TRANSACTION;

CREATE TABLE score
(
    id       bigserial unique primary key,
    username text   not null,
    score    bigint not null
);

COMMIT;