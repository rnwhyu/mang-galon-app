-- +migrate Up
-- +migrate StatementBegin
-- public.roles definition

-- Drop table

-- DROP TABLE public.roles;

CREATE TABLE public.roles (
	id serial4 NOT NULL,
	role_name varchar NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT roles_pk PRIMARY KEY (id)
);
-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	full_name varchar(256) NOT NULL,
	username varchar NOT NULL,
	email varchar NOT NULL,
	"password" varchar NOT NULL,
	role_id int4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT users_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX users_username_idx ON public.users USING btree (username);


-- public.users foreign keys

ALTER TABLE public.users ADD CONSTRAINT users_fk FOREIGN KEY (role_id) REFERENCES public.roles(id);

-- public.item_galon definition

-- Drop table

-- DROP TABLE public.item_galon;

CREATE TABLE public.item_galon (
	id serial4 NOT NULL,
	brand_name varchar(256) NOT NULL,
	stock int4 NOT NULL,
	updatestock_at timestamp NOT NULL DEFAULT now(),
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT item_galon_pk PRIMARY KEY (id)
);
-- public.orders definition

-- Drop table

-- DROP TABLE public.orders;

CREATE TABLE public.orders (
	id serial4 NOT NULL,
	user_id int4 NOT NULL,
	galon_id int4 NOT NULL,
	total_order int4 NOT NULL,
	status varchar NOT NULL,
	updated_at timestamp NOT NULL DEFAULT now(),
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT orders_pk PRIMARY KEY (id)
);


-- public.orders foreign keys

ALTER TABLE public.orders ADD CONSTRAINT orders_fk FOREIGN KEY (user_id) REFERENCES public.users(id);
ALTER TABLE public.orders ADD CONSTRAINT orders_fk_1 FOREIGN KEY (galon_id) REFERENCES public.item_galon(id);

-- +migrate StatementEnd