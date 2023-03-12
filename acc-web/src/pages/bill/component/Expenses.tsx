import {Button, Col, DatePicker, Form, Input, Row, Select, Tabs} from 'antd';

import {inject, observer} from "mobx-react";

const { Option } = Select;


const expenses = ({store}: any) => {
    return (
        <Form name='expenses' autoComplete="off" colon={false}>
            <Row gutter={36}>
                <Col span={6}>
                    <Form.Item label="分类" name="username">
                        <Select defaultValue="lucy"  >
                            <Option value="jack">Jack</Option>
                            <Option value="lucy">Lucy</Option>
                            <Option value="disabled" disabled>
                                Disabled
                            </Option>
                            <Option value="Yiminghe">yiminghe</Option>
                        </Select>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="账户" name="account">
                        <Select defaultValue="lucy"  >
                            <Option value="jack">Jack</Option>
                            <Option value="lucy">Lucy</Option>
                            <Option value="disabled" disabled>
                                Disabled
                            </Option>
                            <Option value="Yiminghe">yiminghe</Option>
                        </Select>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="金额" name="amount">
                        <Input/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="时间" name="date">
                        <DatePicker/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="成员" name="member">
                        <Select defaultValue="lucy"  >
                            <Option value="jack">Jack</Option>
                            <Option value="lucy">Lucy</Option>
                            <Option value="disabled" disabled>
                                Disabled
                            </Option>
                            <Option value="Yiminghe">yiminghe</Option>
                        </Select>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="项目" name="project">
                        <Select defaultValue="lucy"  >
                            <Option value="jack">Jack</Option>
                            <Option value="lucy">Lucy</Option>
                            <Option value="disabled" disabled>
                                Disabled
                            </Option>
                            <Option value="Yiminghe">yiminghe</Option>
                        </Select>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="商家" name="supplier">
                        <Select defaultValue="lucy"  >
                            <Option value="jack">Jack</Option>
                            <Option value="lucy">Lucy</Option>
                            <Option value="disabled" disabled>
                                Disabled
                            </Option>
                            <Option value="Yiminghe">yiminghe</Option>
                        </Select>
                    </Form.Item>
                </Col>
            </Row>
            <Row gutter={36}>
                <Col span={18}>
                    <Form.Item label="备注" name="supplier">
                        <Input/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Button type="primary" size='middle' style={{width: 100}}>保存</Button>
                </Col>
            </Row>
        </Form>
    );
};

export default inject('store')(observer(expenses))
