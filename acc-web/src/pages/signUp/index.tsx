import {Row, Col, Tabs, Form, Input, Button, Space, Divider, Checkbox, ConfigProvider, message} from "antd";

import React from "react";
import style from "@/pages/signUp/style.less";
import {Link} from 'umi'
import DelayButton from "@/components/DelayButton";
import Header from "@/pages/signUp/component/Header";
import Footer from "@/pages/components/Footer";
import zhCN from "antd/lib/locale/zh_CN";
import {useModel} from "@@/plugin-model/useModel";

export default () => {
  const model = useModel('signUpModel');

  const onFinish = async (values: any) => {
    values.captcha = Number(values.captcha)
    let res = await model.register(values)

    if (res.code == 1000) {
      message.info('必须同意协议！')
      return
    } else if (res.code == 1001) {
      message.info('您输入的用户名已经存在！')
      return
    } else if (res.code == 1002) {
      message.info('您输入的电子邮箱已经存在！')
      return
    }
    location.href = './signUp/success'
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };
  return (
    <ConfigProvider locale={zhCN}>
      <div>
        <div className={style.header}>
          <Header/>
        </div>
        <div className={style.main}>
          <div className={style.mainCenter}>
            <div style={{display: "flex", height: 100, marginLeft: 50, alignItems: 'center', justifyContent: 'center'}}>
              <Space>
                <span style={{fontSize: 24}}>注册</span>
                <span>欢迎成为可心记账的用户</span>
              </Space>

            </div>
            <Form name="register" requiredMark={false} labelCol={{span: 8}} wrapperCol={{span: 16}} onFinish={onFinish}
                  onFinishFailed={onFinishFailed}>
              <Form.Item name="email" label="电子邮箱" rules={[{required: true, message: '请输入电子邮箱'}, {type: 'email'}]}>
                <Input/>
              </Form.Item>
              <Form.Item name="captcha" label='验证码' rules={[{required: true, message: '请输入验证码'}]}>
                <Row gutter={8}>
                  <Col flex='auto'><Input/> </Col>
                  <Col flex='100px'><DelayButton>获取验证码</DelayButton> </Col>
                </Row>
              </Form.Item>
              <Form.Item name="username" label="用户名" rules={[{required: true, message: '请输入用户名'}]}>
                <Input/>
              </Form.Item>
              <Form.Item name="nickname" label="昵称" rules={[{required: true, message: '请输入昵称'}]}>
                <Input/>
              </Form.Item>
              <Form.Item name="password" label="设定密码" rules={[{required: true, message: '请输入密码'}]}>
                <Input.Password/>
              </Form.Item>
              <Form.Item
                name="repeatPassword"
                label="重复密码"
                dependencies={['password']}
                rules={[{required: true, message: '请重复输入一次进行确认'},
                  ({getFieldValue}) => ({
                    validator(_, value) {
                      if (!value || getFieldValue('password') === value) {
                        return Promise.resolve();
                      }
                      return Promise.reject(new Error('您两次输入的密码不匹配'));
                    },
                  }),]}>
                <Input.Password/>
              </Form.Item>
              <Form.Item name="agreement" valuePropName="checked" wrapperCol={{span: 16, offset: 8}}>
                <Checkbox>
                  <a href="">我已阅读并同意 《用户服务条款》</a>
                </Checkbox>
              </Form.Item>
              <Form.Item wrapperCol={{span: 16, offset: 8}}>
                <Button type="primary" htmlType="submit" style={{width: '120px'}}>注册</Button>
              </Form.Item>
            </Form>
          </div>
          <div className={style.footer}>
            <Footer/>
          </div>
        </div>
      </div>
    </ConfigProvider>
  );
}