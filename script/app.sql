-- Table: public.apps

-- DROP TABLE IF EXISTS public.apps;

CREATE TABLE IF NOT EXISTS public.apps
(
    id regclass NOT NULL DEFAULT nextval('app_id_seq'::regclass),
    appid integer,
    secret text COLLATE pg_catalog."default",
    status integer,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT apps_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.apps
    OWNER to manager;
-- SEQUENCE: public.app_sequence

-- DROP SEQUENCE IF EXISTS public.app_sequence;

CREATE SEQUENCE IF NOT EXISTS public.app_sequence
    INCREMENT 1
    START 100000
    MINVALUE 100000
    MAXVALUE 999999
    CACHE 1
    OWNED BY apps.appid;

ALTER SEQUENCE public.app_sequence
    OWNER TO manager;
-- SEQUENCE: public.app_id_seq

-- DROP SEQUENCE IF EXISTS public.app_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.app_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.app_id_seq
    OWNER TO manager;