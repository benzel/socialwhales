---

SET statement_timeout           = 0;
SET client_encoding             = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies       = false;
SET client_min_messages         = warning;

DROP SEQUENCE   "s_accounts_id";
DROP TABLE      "t_accounts";
DROP TABLE      "t_profiles";

BEGIN;

CREATE SEQUENCE "s_accounts_id";
CREATE TABLE "t_accounts" (
    "id"                int8 PRIMARY KEY NOT NULL DEFAULT NEXTVAL('s_accounts_id'),
    "role"              int2 NOT NULL DEFAULT 0,
    "status"            int2 NOT NULL DEFAULT 0, -- 0 = not activated
    "email"             varchar(32) NOT NULL UNIQUE,
    "signed_up"         timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at"        timestamp with time zone NOT NULL DEFAULT NOW(),
    "hpassword"   		bytea NOT NULL,
    "token"				varchar(36) NOT NULL
);
CREATE TABLE "t_profiles" (
    "account_id"        int8 NOT NULL PRIMARY KEY REFERENCES t_accounts(id),
    "screen_name"       varchar(16),
    "first_name"        varchar(16),
    "surname"           varchar(16)
);

COMMIT;