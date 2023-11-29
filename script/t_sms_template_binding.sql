-- public.t_sms_template_binding definition

-- Drop table

-- DROP TABLE public.t_sms_template_binding;

CREATE TABLE public.t_sms_template_binding (
	id serial4 NOT NULL,
	template_id int4 NULL,
	template_code varchar NULL,
	template_content varchar NULL,
	channel_id int8 NULL,
	status int4 NULL DEFAULT 0, -- 0 : 待提交 1：待审核  2：审核成功 3：审核失败
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT t_sms_template_binding_pk PRIMARY KEY (id)
);

-- Column comments

COMMENT ON COLUMN public.t_sms_template_binding.status IS '0 : 待提交 1：待审核  2：审核成功 3：审核失败';