import {Button, Checkbox, Form, Input, message, Space} from "antd"
import styles from '../style.less'
import React from "react";
import {OK} from "@/components/Request";
import {inject, observer} from "mobx-react";

const main = ({store}: any) => {
    const onFinish = async (values: any) => {
        let r = await store.register(values)
        if (r.code == OK) {
            store.show(true)
        } else {
            message.info(r.message)
        }
    };

    const onFinishFailed = (errorInfo: any) => {
        console.log('Failed:', errorInfo);
    };

    return (<div className={styles.contentInner}>
        <div
            style={{
                display: "flex",
                height: 100,
                marginLeft: 50,
                alignItems: 'center',
                justifyContent: 'center'
            }}>
            <Space>
                <span style={{fontSize: 24}}>注册</span>
                <span>欢迎成为可心记账的用户</span>
            </Space>

        </div>
        <Form name="register" requiredMark={false} labelCol={{span: 8}} wrapperCol={{span: 16}}
              onFinish={onFinish}
              onFinishFailed={onFinishFailed}>
            <Form.Item name="username" label="用户名" rules={[{required: true, message: '请输入用户名'}]}>
                <Input/>
            </Form.Item>
            <Form.Item name="nickname" label="昵称" rules={[{required: true, message: '请输入昵称'}]}>
                <Input/>
            </Form.Item>
            <Form.Item name="password" label="设定密码" rules={[{required: true, message: '请输入密码'},
                ({getFieldValue}) => ({
                    validator(_, value) {
                        if (!value || /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,20}$/.test(value)) {
                            return Promise.resolve();
                        }
                        return Promise.reject(new Error('至少包含大写字母、小写字母和数字'));
                    },
                })
            ]}>
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
                <Button type="primary" htmlType="submit" style={{width: '200px'}}>注册</Button>
            </Form.Item>
        </Form>
    </div>)
}

export default inject('store')(observer(main))