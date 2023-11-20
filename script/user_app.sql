-- Table: public.user_app

-- DROP TABLE IF EXISTS public.user_app;

CREATE TABLE IF NOT EXISTS public.user_app
(
    id integer NOT NULL DEFAULT nextval('user_app_id_seq'::regclass),
    app_id integer NOT NULL,
    user_id integer NOT NULL,
    quota integer NOT NULL DEFAULT 0,
    created_at date,
    updated_at date,
    service_id integer,
    CONSTRAINT user_app_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user_app
    OWNER to manager;

-- SEQUENCE: public.user_app_id_seq

-- DROP SEQUENCE IF EXISTS public.user_app_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.user_app_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.user_app_id_seq
    OWNER TO manager;