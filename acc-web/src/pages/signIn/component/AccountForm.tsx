import {Button, Divider, Form, Input, message, Space} from "antd";
import {LockOutlined, UserOutlined} from "@ant-design/icons";
import {Link} from "umi";
import React from "react";
import {UserStore} from "@/stores/userStore";
import {inject, observer} from 'mobx-react';
import {setToken} from "@/components/token";

interface AccountProps {
  store?: UserStore
}

const accountForm = ({store}: AccountProps) => {
  const onFinish = async (values: any) => {
    let r = await store.login(values.username, values.password);
    if (r.code == 200) {
      setToken(r.data.token)
      location.href = './account'
    } else {
      message.info(r.message)
    }
  };

  return (<Form name="account" onFinish={onFinish}>
    <Form.Item name="username" rules={[{required: true, message: '请输入您的用户名'}]}>
      <Input prefix={<UserOutlined/>} placeholder="请输入账号名"/>
    </Form.Item>
    <Form.Item name="password" rules={[{required: true, message: '请输入您的密码'}]}>
      <Input prefix={<LockOutlined/>} type="password" placeholder="请输入密码"/>
    </Form.Item>
    <Form.Item>
      <Button type="primary" htmlType="submit" style={{width: '100%'}}>登录</Button>
    </Form.Item>
    <Form.Item wrapperCol={{offset: 13}}>
      <Space split={<Divider type="vertical"/>} size={0}>
        <a>忘记密码？</a>
        <Link to='/register'>免费注册</Link>
      </Space>
    </Form.Item>
  </Form>)
}

export default inject('store')(observer(accountForm))
