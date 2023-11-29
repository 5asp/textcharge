-- public.t_sms_appinfo definition

-- Drop table

-- DROP TABLE public.t_sms_appinfo;

CREATE TABLE public.t_sms_appinfo (
	id serial4 NOT NULL, -- 主键id
	app_key varchar NULL, -- 应用ID
	app_name varchar NULL, -- 应用名称
	app_secret varchar NULL, -- 应用密钥
	status int4 NULL, -- 应用状态
	remark varchar NULL, -- 应用说明
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT t_sms_appinfo_pk PRIMARY KEY (id),
	CONSTRAINT t_sms_appinfo_un UNIQUE (app_key)
);

-- Column comments

COMMENT ON COLUMN public.t_sms_appinfo.id IS '主键id';
COMMENT ON COLUMN public.t_sms_appinfo.app_key IS '应用ID';
COMMENT ON COLUMN public.t_sms_appinfo.app_name IS '应用名称';
COMMENT ON COLUMN public.t_sms_appinfo.app_secret IS '应用密钥';
COMMENT ON COLUMN public.t_sms_appinfo.status IS '应用状态';
COMMENT ON COLUMN public.t_sms_appinfo.remark IS '应用说明';