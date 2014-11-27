# 数据库

## user

id					ID				记录唯一标识
username			用户名			用户唯一标识，可修改
password			密码			密码由客户端md5加密串与盐连接，然后再次md5
salt				盐				随机串
email				Email			用户标识，找回密码依据，可修改
level				用户级别		区分普通用户、学生、教师、管理员
registor_time		注册时间		注册时间，mysql生成
nickname			昵称			昵称，可修改

## problem

id					题目ID
type				类型			题目分为`断言式`[assert]和`输入输出流`[IO]式
description			题干描述		对题目的描述
pre_code			预代码			预代码仅对assert类型题目有效
io_data				IO测试数据		IO测试数据仅对IO类型题目有效
tags				类型标签		题目分类的依据
level				等级			目前仅分公开题目和私有题目
