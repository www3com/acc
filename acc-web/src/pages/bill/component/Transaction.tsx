import {Button, Col, DatePicker, Form, Input, Row, Select, TreeSelect} from 'antd';
import dayjs from "dayjs";

interface Item {
    label: string,
    options: [],
    visible?: boolean
}

interface Transaction {
    type: number,
    account: Item,
    cpAccount: Item,
    project?: Item,
    member?: Item,
    supplier?: Item
    onFinish?: Function
}

export default ({type, account, cpAccount, project, member, supplier, onFinish}: Transaction) => {
    const finishHandler = async (values: any) => {
        if (onFinish) {
            onFinish({...values, type: type, tradingTime: values.tradingTime.valueOf()})
        }
    };

    return (
        <Form name='transfer' autoComplete="off" colon={false} onFinish={finishHandler}
              initialValues={{tradingTime: dayjs()}}>
            <Row gutter={36}>
                <Col span={6}>
                    <Form.Item label={account.label} name="accountId"
                               rules={[{required: true, message: '请选择' + account.label}]}>
                        <TreeSelect
                            style={{width: '100%'}}
                            dropdownStyle={{maxHeight: 400, overflow: 'auto'}}
                            fieldNames={{label: 'name', value: 'id'}}
                            allowClear
                            treeData={account.options}
                        />
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label={cpAccount.label} name="cpAccountId"
                               rules={[{required: true, message: '请选择' + cpAccount.label}]}>
                        <TreeSelect style={{width: '100%'}}
                                    dropdownStyle={{maxHeight: 400, overflow: 'auto'}}
                                    placeholder="选择账户"
                                    fieldNames={{label: 'name', value: 'id'}}
                                    allowClear
                                    treeData={cpAccount.options}/>
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
                        <Select fieldNames={{label: 'name', value: 'id'}} options={project?.options}/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="成员" name="memberId">
                        <Select fieldNames={{label: 'name', value: 'id'}} options={member?.options}/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="商家" name="supplierId">
                        <Select fieldNames={{label: 'name', value: 'id'}} options={supplier?.options}/>
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

