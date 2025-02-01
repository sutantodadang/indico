-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "i_users" (
  "user_id" uuid PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "role_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "i_users" ADD FOREIGN KEY ("role_id") REFERENCES "i_users_roles" ("user_role_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "i_users" DROP CONSTRAINT i_users_role_id_fkey;
DROP TABLE IF EXISTS "i_users";

-- +goose StatementEnd
