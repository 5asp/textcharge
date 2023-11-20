-- Table: public.services

-- DROP TABLE IF EXISTS public.services;

CREATE TABLE IF NOT EXISTS public.services
(
    id integer NOT NULL DEFAULT nextval('service_id_seq'::regclass),
    service_name text COLLATE pg_catalog."default",
    status integer NOT NULL DEFAULT 0,
    created_at date,
    updated_at date,
    CONSTRAINT services_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.services
    OWNER to manager;


    -- SEQUENCE: public.service_id_seq

-- DROP SEQUENCE IF EXISTS public.service_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.service_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.service_id_seq
    OWNER TO manager;