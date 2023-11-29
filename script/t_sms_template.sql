-- public.t_sms_template definition

-- Drop table

-- DROP TABLE public.t_sms_template;

CREATE TABLE public.t_sms_template (
	id serial4 NOT NULL,
	template_name varchar NULL,
	"content" varchar NULL,
	sign_name varchar NULL,
	template_type varchar NULL DEFAULT 0, -- 0：验证码。\n1：短信通知。\n2：推广短信。\n3：国际/港澳台消息
	status int4 NULL DEFAULT 1, -- 状态 0：有效 1：无效
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	remark varchar NULL,
	CONSTRAINT t_sms_template_pk PRIMARY KEY (id)
);

-- Column comments

COMMENT ON COLUMN public.t_sms_template.template_type IS '0：验证码。\n1：短信通知。\n2：推广短信。\n3：国际/港澳台消息';
COMMENT ON COLUMN public.t_sms_template.status IS '状态 0：有效 1：无效';