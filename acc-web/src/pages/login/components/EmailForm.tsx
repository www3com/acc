import {Button, Col, Form, Input, Row} from "antd";
import {LockOutlined, MailOutlined} from "@ant-design/icons";
import DelayButton from "@/components/DelayButton";
import React from "react";

export default () => {
    const onClick = () => {
        return true
    }

    return (<Form name="email">
        <Form.Item name="email" rules={[{required: true, message: '请输入您的电子邮箱'}]}>
            <Input prefix={<MailOutlined/>} placeholder="请输入电子邮箱"/>
        </Form.Item>
        <Form.Item>
            <Row gutter={8}>
                <Col flex='auto'>
                    <Input prefix={<LockOutlined/>} type="password" placeholder="请输入验证码"/>
                </Col>
                <Col flex='100px'>
                    <DelayButton onClick={onClick}>获取验证码</DelayButton>
                </Col>
            </Row>
        </Form.Item>
        <Form.Item>
            <Button type="primary" htmlType="submit" style={{width: '100%'}}>登录</Button>
        </Form.Item>
    </Form>)
}
