# TextCharge

## 组件:

   - API网关: 接收短信发送请求。
   - 身份验证服务: 验证用户身份。
   - 短信发送服务: 处理短信发送逻辑。
   - 计费服务: 计算发送短信的费用。
   - 用户账户服务: 管理用户的账户信息和余额。
   - 数据库: 存储用户数据和交易记录。
   - 消息队列: 在服务之间传递消息。

## 流程:

   - 接收请求：API网关接收短信发送请求。
   - 身份验证：验证用户身份和授权。
   - 短信发送：处理短信发送逻辑。
   - 计费处理：计算费用并更新用户账户。
   - 记录更新：更新数据库中的交易和用户账户信息。
   - 响应返回：向用户发送操作结果。

## 接口:

   - account/login
   - account/register
   
   - app/create
   - app/list
   - app/delete
   - app/info
   - app/edit

   - billing/log
   
   - sms/send

   - api/user
   - login
   - register
   - logout
   - reset-password
   - email/verification-notification
   - sanctum/csrf-cookie


## Flow
![TextCharge](./tech.jpg "TextCharge")