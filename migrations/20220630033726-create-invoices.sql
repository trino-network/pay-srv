-- +migrate Up
CREATE TABLE invoices (
                       id uuid NOT NULL DEFAULT uuid_generate_v4 (),
                       user_id uuid NOT NULL,
                       title varchar(256)  NOT NULL,
                       network varchar(32)  NOT NULL,
                       recipient varchar(256)  NOT NULL,
                       amount bigint,
                       jump_url varchar(256)  NOT NULL,
                       notify_url varchar(256)  NOT NULL,
                       status varchar(32)  NOT NULL,
                       metadata varchar(512)  NOT NULL,
                       txn_hash varchar (256) DEFAULT NULL,
                       created_at timestamptz NOT NULL,
                       updated_at timestamptz NOT NULL,
                       CONSTRAINT invoices_pkey PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS invoices;