import {Button, Col, DatePicker, Form, Input, Row, Select, Tabs, TreeSelect} from 'antd';
import {inject, observer} from "mobx-react";

const expenses = ({store}: any) => {
    return (
        <Form name='transfer' autoComplete="off" colon={false}>
            <Row gutter={36}>
                <Col span={6}>
                    <Form.Item label="转出" name="account">
                        <TreeSelect
                            showSearch
                            style={{width: '100%'}}
                            dropdownStyle={{maxHeight: 400, overflow: 'auto'}}
                            placeholder="选择分类"
                            fieldNames={{label: 'name', value: 'id'}}
                            allowClear
                            treeData={store.cpAccounts}
                        />
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="转入" name="cpAccount">
                        <TreeSelect
                            showSearch
                            style={{width: '100%'}}
                            dropdownStyle={{maxHeight: 400, overflow: 'auto'}}
                            placeholder="选择账户"
                            fieldNames={{label: 'name', value: 'id'}}
                            allowClear
                            treeData={store.cpAccounts}
                        />
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="金额" name="amount">
                        <Input/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="时间" name="tradingTime">
                        <DatePicker/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="项目" name="project">
                        <Select fieldNames={{label: 'name', value: 'id'}} options={store.projects}/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="成员" name="member">
                        <Select fieldNames={{label: 'name', value: 'id'}} options={store.members}/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="商家" name="supplier">
                        <Select fieldNames={{label: 'name', value: 'id'}} options={store.suppliers}/>
                    </Form.Item>
                </Col>
            </Row>
            <Row gutter={36}>
                <Col span={18}>
                    <Form.Item label="备注" name="remark">
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
