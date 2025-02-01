-- +goose Up
-- +goose StatementBegin
CREATE TYPE "type_order" AS ENUM (
  'RECEIVE',
  'SHIP'
);


CREATE TABLE IF NOT EXISTS "i_orders" (
  "order_id" uuid PRIMARY KEY NOT NULL,
  "user_id" uuid NOT NULL,
  "product_id" uuid NOT NULL,
  "quantity" int NOT NULL,
  "order_type" type_order NOT NULL,
  "order_status" int NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);


ALTER TABLE "i_orders" ADD FOREIGN KEY ("user_id") REFERENCES "i_users" ("user_id");

ALTER TABLE "i_orders" ADD FOREIGN KEY ("product_id") REFERENCES "i_products" ("product_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "i_orders" DROP CONSTRAINT i_orders_user_id_fkey;

ALTER TABLE "i_orders" DROP CONSTRAINT i_orders_product_id_fkey;

DROP TABLE IF EXISTS "i_orders";

-- +goose StatementEnd
