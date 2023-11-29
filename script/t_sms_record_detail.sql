-- public.t_sms_record_detail definition

-- Drop table

-- DROP TABLE public.t_sms_record_detail;

CREATE TABLE public.t_sms_record_detail (
	id serial4 NOT NULL, -- 自增ID
	record_id int4 NULL, -- 关联发送短信记录id
	app_id varchar NULL,
	"content" varchar NULL,
	send_status int4 NULL DEFAULT '-1'::integer, -- -1：待发送 /  0：已发送  / 1 : 发送失败'
	report_status int4 NULL DEFAULT 0, -- 短信报告 0 ： 待回执  1：发送成功 2 : 发送失败
	mobile varchar NULL,
	msgid varchar NULL,
	channel_id int4 NULL,
	send_at int4 NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT t_sms_record_detail_pk PRIMARY KEY (id)
);

-- Column comments

COMMENT ON COLUMN public.t_sms_record_detail.id IS '自增ID';
COMMENT ON COLUMN public.t_sms_record_detail.record_id IS '关联发送短信记录id';
COMMENT ON COLUMN public.t_sms_record_detail.send_status IS '-1：待发送 /  0：已发送  / 1 : 发送失败''';
COMMENT ON COLUMN public.t_sms_record_detail.report_status IS '短信报告 0 ： 待回执  1：发送成功 2 : 发送失败';