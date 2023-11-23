-- Table: public.user_service

-- DROP TABLE IF EXISTS public.user_service;

CREATE TABLE IF NOT EXISTS public.user_services
(
    id regclass NOT NULL,
    service_id integer NOT NULL DEFAULT 0,
    user_id integer NOT NULL DEFAULT 0,
    quota integer NOT NULL DEFAULT 0,
    created_at time with time zone,
    updated_at time with time zone,
    CONSTRAINT user_service_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user_services
    OWNER to "aheadIV";


-- SEQUENCE: public.user_service_seq

-- DROP SEQUENCE IF EXISTS public.user_service_seq;

CREATE SEQUENCE IF NOT EXISTS public.user_service_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1
    OWNED BY user_services.id;

ALTER SEQUENCE public.user_service_seq
    OWNER TO "aheadIV";