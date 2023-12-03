CREATE TYPE "account_type" AS ENUM (
    'client',
    'admin'
    );

CREATE TYPE "account_status" AS ENUM (
    'pending',
    'activated'
    );

CREATE TYPE "project_status" AS ENUM (
    'registering',
    'progressing',
    'finished'
    );

CREATE TABLE IF NOT EXISTS "user_account"
(
    "user_id"    serial PRIMARY KEY,
    "username"   text UNIQUE              NOT NULL,
    "password"   text                     NOT NULL,
    "type"       account_type             NOT NULL DEFAULT 'client',
    "status"     account_status           NOT NULL DEFAULT 'pending',
    "created_at" timestamp with time zone NOT NULL DEFAULT (now()),
    "updated_at" timestamp with time zone NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "user_profile"
(
    "id"         serial PRIMARY KEY,
    "first_name" text                     NOT NULL,
    "last_name"  text                     NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT (now()),
    "updated_at" timestamp with time zone NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "project"
(
    "id"           serial PRIMARY KEY,
    "user_profile" integer,
    "name"         text                     NOT NULL,
    "description"  text                     NOT NULL,
    "price"        integer                  NOT NULL,
    "paid"         integer                  NOT NULL,
    "status"       project_status           NOT NULL DEFAULT 'registering',
    "start_time"   timestamp with time zone NOT NULL,
    "end_time"     timestamp with time zone NOT NULL,
    "created_at"   timestamp with time zone NOT NULL DEFAULT (now()),
    "updated_at"   timestamp with time zone NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "project"."description" IS 'Which technologies and algorithms used';

COMMENT ON COLUMN "project"."price" IS 'Price of the project';

COMMENT ON COLUMN "project"."paid" IS 'How much money client paid';

ALTER TABLE "user_profile"
    ADD FOREIGN KEY ("id") REFERENCES "user_account" ("user_id") ON DELETE CASCADE;

ALTER TABLE "project"
    ADD FOREIGN KEY ("user_profile") REFERENCES "user_profile" ("id") ON DELETE CASCADE;
