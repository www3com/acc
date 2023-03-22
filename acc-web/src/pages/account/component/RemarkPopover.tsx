import React, {useState} from 'react';
import {Button, Form, Input, Modal, Popover, Space} from 'antd';
import {inject, observer} from 'mobx-react';
import {EditOutlined} from '@ant-design/icons';
import styles from '../style.less'
import {updateRemark} from "@/services/account";


const RemarkPopover = ({record}: any) => {

    const [open, setOpen] = useState(false);
    const onOk = async (values: any) => {
        await updateRemark(record.id, values.remark);
        record.remark = values.remark;
        setOpen(false);
    };

    const content = (<>
        <Form initialValues={record} onFinish={onOk}>
            <Form.Item name='remark' style={{width: 300}} rules={[{required: true, message: '请输入描述'}]}>
                <Input autoComplete="off"/>
            </Form.Item>
            <Space style={{width: '100%'}} direction='vertical' align={'end'}>
                <Button type={'primary'} size={'small'} htmlType={'submit'}>确定</Button>
            </Space>
        </Form>
    </>);

    return (
        <div className={styles.item}>
            <span>{record.remark == "" ? <>&nbsp;</> : record.remark}</span>
            <Popover title='描述' open={open} placement="bottom" destroyTooltipOnHide
                     content={content} onOpenChange={() => setOpen(false)}>
                <a className={styles.edit} onClick={() => setOpen(true)}><EditOutlined/></a>
            </Popover>
        </div>);
};

export default RemarkPopover;