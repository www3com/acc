import React, {useEffect} from 'react';
import {Card, Divider, Dropdown, message, Space, Statistic, Table, Modal, Tooltip} from 'antd';

import {
    CloseOutlined, DeleteOutlined,
    DownOutlined, ExclamationCircleOutlined, MinusOutlined,
    PayCircleOutlined, PlusOutlined, RedEnvelopeOutlined,
    TransactionOutlined,
} from '@ant-design/icons';
import {inject, observer} from 'mobx-react';
import Dialog from '@/pages/account/component/Dialog';
import NamePopover from "@/pages/account/component/NamePopover";
import BalancePopover from "@/pages/account/component/BalancePopover";
import RemarkPopover from "@/pages/account/component/RemarkPopover";

const {confirm} = Modal;
const root = ({store}: any) => {
    useEffect(() => {
        store.list();
    }, []);

    const onDelete = (name: string, code: string) => {
        confirm({
            title: "确认删除",
            icon: <ExclamationCircleOutlined/>,
            content: '确认删除账户[' + name + ']吗？',
            onOk: () => {
                store.deleteAccount(code)
            }
        })
    }

    const renderDelete = (record: any) => {
        if (record.level == 1) {
            return <></>
        }
        return <a onClick={() => onDelete(record.name, record.code)}>
            <Tooltip title='删除账户' placement='bottom'>
                <DeleteOutlined />
            </Tooltip>
        </a>
    }

    const columns: any = [
        {
            title: '账户',
            dataIndex: 'name',
            width: '20%',
            key: 'name',
            render: (_: any, record: any) => {
                return <NamePopover record={record}/>
            }
        }, {
            title: '描述',
            dataIndex: 'remark',
            key: 'remark',
            render: (_: any, record: any) => {
                return <RemarkPopover record={record}/>
            }
        }, {
            title: '余额',
            dataIndex: 'balance',
            width: '20%',
            key: 'balance',
            render: (_: any, record: any) => {
                return <BalancePopover record={record}/>
            }
        },
        {
            title: '操作',
            dataIndex: 'operator',
            width: '100px',
            align: 'center',
            render: (_: any, record: any) =>
                <Space style={{width: 20}}>
                    <a onClick={() => store.showDialog({parentId: record.id})}>
                        <Tooltip title='新建账户' placement='bottom'>
                            <PlusOutlined/>
                        </Tooltip>
                    </a>
                    {renderDelete(record)}
                </Space>,
        }];

    return (
        <div style={{backgroundColor: 'white'}}>
            <Card>
                <Space split={<Divider type='vertical' style={{height: 50}}/>} size={50}>
                    <Statistic title='总资产' value={store.accountDetails.total}
                               groupSeparator=','
                               prefix={<PayCircleOutlined/>}/>
                    <Statistic title='总负债' value={store.accountDetails.debt}
                               groupSeparator=','
                               prefix={<TransactionOutlined/>}/>
                    <Statistic title='净资产' value={store.accountDetails.netAmount}
                               groupSeparator=','
                               prefix={<RedEnvelopeOutlined/>}/>
                </Space>
            </Card>
            <Table
                size={"middle"}
                rowKey={'id'}
                columns={columns}
                dataSource={store.accountDetails.details}
                pagination={false}/>
            <Dialog/>
        </div>
    );
};

export default inject('store')(observer(root));
