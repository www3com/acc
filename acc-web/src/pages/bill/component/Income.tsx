import {Button, Col, DatePicker, Form, Input, notification, Row, Select, Tabs, TreeSelect} from 'antd';
import {inject, observer} from "mobx-react";
import dayjs from "dayjs";

const income = ({store}: any) => {
    const onFinish = async (values: any) => {
        await store.saveTransaction({...values, type: 2, tradingTime: values.tradingTime.valueOf()});
        notification.success({
            message: '保存收入明细成功',
        })
    };

    return (
        <Form name='income' autoComplete="off" colon={false} onFinish={onFinish} initialValues={{tradingTime: dayjs()}}>
            <Row gutter={36}>
                <Col span={6}>
                    <Form.Item label="分类" name="accountId" rules={[{required: true, message: '请选择分类'}]}>
                        <TreeSelect
                            style={{width: '100%'}}
                            dropdownStyle={{maxHeight: 400, overflow: 'auto'}}
                            placeholder="选择分类"
                            fieldNames={{label: 'name', value: 'id'}}
                            allowClear
                            treeData={store.incomes}
                        />
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="账户" name="cpAccountId" rules={[{required: true, message: '请选择账户'}]}>
                        <TreeSelect

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
                    <Form.Item label="金额" name="amount" rules={[{required: true, message: '请输入金额'}]}>
                        <Input/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="时间" name="tradingTime" rules={[{required: true, message: '请选择时间'}]}>
                        <DatePicker/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="项目" name="projectId">
                        <Select fieldNames={{label: 'name', value: 'id'}} options={store.projects}/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="成员" name="memberId">
                        <Select fieldNames={{label: 'name', value: 'id'}} options={store.members}/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="商家" name="supplierId">
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
                    <Button type="primary" size='middle' style={{width: 100}} htmlType="submit">保存</Button>
                </Col>
            </Row>
        </Form>
    );
};

export default inject('store')(observer(income))
