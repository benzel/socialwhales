---
BEGIN;

CREATE SEQUENCE "s_users_id";
CREATE TABLE "t_users" (
    "id"                int8 NOT NULL DEFAULT NEXTVAL('s_users_id') PRIMARY KEY,
    "role"              int2 NOT NULL DEFAULT 0,
    "email"             varchar(32),
    "screen_name"       varchar(16),
    "first_name"        varchar(16),
    "surname"           varchar(16),
    "signed_up"         timestamp with time zone,
    "hashed_password"   bytea
);

COMMIT;