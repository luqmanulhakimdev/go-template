CREATE OR REPLACE FUNCTION update_modified_column() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';
CREATE TABLE public.settings (
	id serial4 NOT NULL,
	"name" varchar(50) NULL,
	value jsonb NOT NULL DEFAULT '{}'::jsonb,
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	CONSTRAINT settings_pkey PRIMARY KEY (id)
);
CREATE TRIGGER settings BEFORE UPDATE ON settings FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
GRANT ALL ON SEQUENCE public.settings_id_seq TO amarthaplus;
GRANT SELECT,INSERT,UPDATE,DELETE ON settings TO amarthaplus;
