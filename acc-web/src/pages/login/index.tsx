import {Row, Col, Tabs, Form, Input, Button, Space, Divider} from "antd";

import React from "react";
import style from "@/pages/login/style.less";
import {LockOutlined, MailOutlined, UserOutlined} from "@ant-design/icons";

const {TabPane} = Tabs;

export default () => {
  return (
    <div style={{backgroundColor: "white"}}>
      <div className={style.headerWrapper}>
        <Row className={style.header}>
          <Col flex="100px" style={{margin: 'auto'}}>
            <img src={'./logo3.png'} width={170} height={45}/>
          </Col>
          <Col flex="auto">
            <a style={{float: "right", marginTop: 15}}>注册</a>
          </Col>
        </Row>
      </div>
      <div className={style.mainWrapper}>
        <div className={style.main}>
          <Row>
            <Col span={16}>
              <img src="./login2.png" width='500px' height='350px'/>
            </Col>
            <Col span={8}>
              <Tabs>
                <TabPane tab="账户登录" key="account">
                  <Form name="account">
                    <Form.Item name="username" rules={[{required: true, message: '请输入您的用户名'}]}>
                      <Input prefix={<UserOutlined/>} placeholder="请输入用户名"/>
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
                        <a>免费注册</a>
                      </Space>
                    </Form.Item>
                  </Form>
                </TabPane>
                <TabPane tab="邮箱登录" key="email">
                  <Form name="email">
                    <Form.Item name="email" rules={[{required: true, message: '请输入您的电子邮箱'}]}>
                      <Input prefix={<MailOutlined/>} placeholder="请输入电子邮箱"/>
                    </Form.Item>
                    <Form.Item>
                      <Row gutter={8}>
                        <Col flex='auto'>
                          <Input prefix={<LockOutlined/>} type="password" placeholder="请输入验证码"/>
                        </Col>
                        <Col flex='100px'>
                          <Button>获取验证码</Button>
                        </Col>
                      </Row>
                    </Form.Item>

                    <Form.Item>
                      <Button type="primary" htmlType="submit" style={{width: '100%'}}>登录</Button>
                    </Form.Item>
                  </Form>
                </TabPane>
              </Tabs>
            </Col>
          </Row>
        </div>
        <div className={style.footer}>
          Copyright © 2022-2023 upbos.com All Rights Reserved.
        </div>
      </div>
    </div>
  );
}
