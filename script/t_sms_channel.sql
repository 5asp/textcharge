-- public.t_sms_channel definition

-- Drop table

-- DROP TABLE public.t_sms_channel;

CREATE TABLE public.t_sms_channel (
	id serial4 NOT NULL, -- 主键id
	channel_name varchar NULL, -- 渠道名称
	channel_type varchar NULL, -- 渠道类型
	channel_appkey varchar NULL, -- 渠道用户名
	channel_appsecret varchar NULL, -- 渠道密码
	channel_domain varchar NULL, -- 渠道请求地址
	ext_properties varchar NULL,
	status int2 NULL, -- 状态0：启用 1：禁用
	send_order int4 NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	quota int4 NOT NULL DEFAULT 0, -- 通道余额
	CONSTRAINT t_sms_channel_pk PRIMARY KEY (id)
);

-- Column comments

COMMENT ON COLUMN public.t_sms_channel.id IS '主键id';
COMMENT ON COLUMN public.t_sms_channel.channel_name IS '渠道名称';
COMMENT ON COLUMN public.t_sms_channel.channel_type IS '渠道类型';
COMMENT ON COLUMN public.t_sms_channel.channel_appkey IS '渠道用户名';
COMMENT ON COLUMN public.t_sms_channel.channel_appsecret IS '渠道密码';
COMMENT ON COLUMN public.t_sms_channel.channel_domain IS '渠道请求地址';
COMMENT ON COLUMN public.t_sms_channel.status IS '状态0：启用 1：禁用';
COMMENT ON COLUMN public.t_sms_channel.quota IS '通道余额';