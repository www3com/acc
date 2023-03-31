import React, {useState} from 'react';
import {Button, Form, Input, InputNumber, Popover, Space} from 'antd';
import {EditOutlined} from '@ant-design/icons';
import styles from '../style.less'
import {updateAccount} from "@/services/account";

const BalancePopover = ({record}: any) => {
    const [open, setOpen] = useState(false);
    const onOk = async (values: any) => {
        await updateAccount({type: 'balance', id: record.id, amount: record.balance - values.balance});
        record.balance = values.balance;
        setOpen(false);
    };

    const content = (<>
        <Form initialValues={record} onFinish={onOk}>
            <Form.Item name='balance' style={{width: 200}} rules={[{required: true, message: '请输入账户余额'}]}>
                <InputNumber controls={false} style={{width: 180}}/>
            </Form.Item>
            <Space style={{width: '100%'}} direction='vertical' align={'end'}>
                <Button type={'primary'} size={'small'} htmlType={'submit'}>确定</Button>
            </Space>
        </Form>
    </>);

    return (
        <div className={styles.item}>
            <span>{record.balance}</span>
            <Popover title='账户余额' open={open} placement="bottom" destroyTooltipOnHide
                     content={content} onOpenChange={() => setOpen(false)}>
                <a className={styles.edit} onClick={() => setOpen(true)}><EditOutlined/></a>
            </Popover>
        </div>);
};

export default BalancePopover;