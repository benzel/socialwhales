---

DROP SEQUENCE "s_accounts_id";
DROP TABLE "t_accounts";
DROP TABLE "t_profiles";

BEGIN;

CREATE SEQUENCE "s_accounts_id";
CREATE TABLE "t_accounts" (
    "id"                int8 PRIMARY KEY NOT NULL DEFAULT NEXTVAL('s_accounts_id'),
    "role"              int2 NOT NULL DEFAULT 0,
    "email"             varchar(32) NOT NULL,
    "signed_up"         timestamp with time zone NOT NULL DEFAULT NOW(),
    "hashed_password"   bytea NOT NULL
);
CREATE TABLE "t_profiles" (
    "account_id"        int8 NOT NULL PRIMARY KEY REFERENCES t_accounts(id),
    "screen_name"       varchar(16),
    "first_name"        varchar(16),
    "surname"           varchar(16)
);

COMMIT;