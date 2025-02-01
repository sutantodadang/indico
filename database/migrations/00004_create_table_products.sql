-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "i_products" (
  "product_id" uuid PRIMARY KEY NOT NULL,
  "sku" varchar NOT NULL,
  "name" varchar NOT NULL,
  "quantity" int NOT NULL,
  "location_id" uuid,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "i_products" ADD FOREIGN KEY ("location_id") REFERENCES "i_warehouses" ("warehouse_id");


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "i_products" DROP CONSTRAINT i_products_location_id_fkey;
DROP TABLE IF EXISTS "i_products";


-- +goose StatementEnd
