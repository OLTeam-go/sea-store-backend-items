CREATE TABLE items (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  merchant_id uuid NOT NULL,
  "name" varchar NOT NULL,
  category varchar NOT NULL,
  description text NULL,
  price numeric NOT NULL,
  quantity int8 NOT NULL,
  created_at timestamp NULL DEFAULT now(),
  updated_at timestamp NULL DEFAULT now(),
  deleted_at timestamp NULL,
  CONSTRAINT items_pk PRIMARY KEY (id)
);