import React, {useState} from 'react';
import {Button, Form, Input, Modal, Popover, Space} from 'antd';
import {EditOutlined} from '@ant-design/icons';
import styles from '../style.less'
import {updateAccount} from "@/services/account";


const EditPopover = ({record}: any) => {

    const [open, setOpen] = useState(false);
    const onOk = async (values: any) => {
        await updateAccount({type: 'name', id: record.id, name: values.name});
        record.name = values.name;
        setOpen(false);
    };

    const content = (<>
        <Form initialValues={record} onFinish={onOk}>
            <Form.Item name='name' style={{width: 200}} rules={[{required: true, message: '请输入账户名称'}]}>
                <Input autoComplete="off"/>
            </Form.Item>
            <Space style={{width: '100%'}} direction='vertical' align={'end'}>
                <Button type={'primary'} size={'small'} htmlType={'submit'}>确定</Button>
            </Space>
        </Form>
    </>);

    return (
        <div className={styles.item}>
            <span>{record.name}</span>
            <Popover title='账户名称' open={open} placement="bottom" destroyTooltipOnHide
                     content={content} onOpenChange={() => setOpen(false)}>
                <a className={styles.edit} onClick={() => setOpen(true)}><EditOutlined/></a>
            </Popover>
        </div>);
};

export default EditPopover;