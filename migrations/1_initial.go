package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table items...")
		_, err := db.Exec(`create extension if not exists "uuid-ossp"`)
		if err != nil {
			return err
		}
		_, err = db.Exec(`CREATE TABLE items (
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
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table items...")
		_, err := db.Exec(`DROP TABLE items`)
		return err
	})
}
