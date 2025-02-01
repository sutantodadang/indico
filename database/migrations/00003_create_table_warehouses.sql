-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "i_warehouses" (
  "warehouse_id" uuid PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "capacity" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "i_warehouses";

-- +goose StatementEnd
