-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "user_role" AS ENUM ('ADMIN', 'STAFF');
 
CREATE TABLE IF NOT EXISTS "i_users_roles" (
  "user_role_id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "unique_name" user_role NOT NULL,
  "name" varchar NOT NULL,
  "status" bool NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

INSERT INTO "i_users_roles" ("unique_name", "name", "status") VALUES ('ADMIN', 'Admin', true),('STAFF', 'Staff', true);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "i_users_roles";

DROP TYPE "user_role";
-- +goose StatementEnd
