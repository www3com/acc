import {Button, Col, DatePicker, Form, Input, Row, Select, Space, TreeSelect} from 'antd';
import dayjs from "dayjs";

interface Item {
    label?: string,
    options?: [],
    disabled?: boolean
}

interface TransactionTab {
    type: number,
    account: Item,
    cpAccount: Item,
    project?: Item,
    member?: Item,
    supplier?: Item
    onFinish?: Function
}

export default ({type, account, cpAccount, project, member, supplier, onFinish}: TransactionTab) => {
    const finishHandler = async (values: any) => {
        if (onFinish) {
            onFinish({...values, type: type, tradingTime: values.tradingTime.valueOf()})
        }
    };

    const labelCol = {span: 8};
    const formItemSpan = 6;
    return (
        <Form name='transfer' autoComplete="off" colon={false} requiredMark={false} onFinish={finishHandler}
              initialValues={{tradingTime: dayjs()}} style={{paddingRight: 20}}>
            <Row>
                <Col span={formItemSpan}>
                    <Form.Item label={account.label} labelCol={labelCol} name="accountId"
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
                <Col span={formItemSpan}>
                    <Form.Item label={cpAccount.label} labelCol={labelCol} name="cpAccountId"
                               rules={[{required: true, message: '请选择' + cpAccount.label}]}>
                        <TreeSelect style={{width: '100%'}}
                                    dropdownStyle={{maxHeight: 400, overflow: 'auto'}}
                                    placeholder="选择账户"
                                    fieldNames={{label: 'name', value: 'id'}}
                                    allowClear
                                    treeData={cpAccount.options}/>
                    </Form.Item>
                </Col>
                <Col span={formItemSpan}>
                    <Form.Item label="金额" labelCol={labelCol} name="amount"
                               rules={[{required: true, message: '请输入金额'}]}>
                        <Input/>
                    </Form.Item>
                </Col>
                <Col span={formItemSpan}>
                    <Form.Item label="时间" labelCol={labelCol} name="tradingTime"
                               rules={[{required: true, message: '请选择时间'}]}>
                        <DatePicker style={{width: '100%'}}/>
                    </Form.Item>
                </Col>
            </Row>
            <Row>
                <Col span={formItemSpan}>
                    <Form.Item label="项目" labelCol={labelCol} name="projectId">
                        <Select fieldNames={{label: 'name', value: 'id'}} options={project?.options}/>
                    </Form.Item>
                </Col>
                <Col span={formItemSpan}>
                    <Form.Item label="成员" labelCol={labelCol} name="memberId">
                        <Select disabled={member?.disabled} fieldNames={{label: 'name', value: 'id'}} options={member?.options}/>
                    </Form.Item>
                </Col>
                <Col span={formItemSpan}>
                    <Form.Item label="商家" labelCol={labelCol} name="supplierId">
                        <Select disabled={supplier?.disabled} fieldNames={{label: 'name', value: 'id'}} options={supplier?.options}/>
                    </Form.Item>
                </Col>
                <Col span={formItemSpan}>
                    <Form.Item label="备注" labelCol={labelCol} name="remark">
                        <Input/>
                    </Form.Item>
                </Col>
            </Row>
            <Space style={{width: '100%', justifyContent: 'end'}}>
                <Button type="primary" size='middle' style={{width: 150}} htmlType="submit">保存</Button>
            </Space>
        </Form>
    );
};

