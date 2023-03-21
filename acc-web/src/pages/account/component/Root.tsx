import React, {useEffect} from 'react';
import {Card, Divider, Dropdown, message, Space, Statistic, Table, Modal} from 'antd';

import {
    DownOutlined, ExclamationCircleOutlined,
    PayCircleOutlined, RedEnvelopeOutlined,
    TransactionOutlined,
} from '@ant-design/icons';
import {inject, observer} from 'mobx-react';
import Dialog from '@/pages/account/component/Dialog';

const {confirm} = Modal;
const root = ({store}: any) => {

    useEffect(() => {
        store.list();
    }, []);

    const operatorClick = (type: string, record: any) => {
        // debugger
        if (type == '1') {
            store.showDialog(record)
        } else if (type == '2') {
            confirm({
                title: "确认删除",
                icon: <ExclamationCircleOutlined/>,
                content: '确认删除账户[' + record.name + ']吗？',
                onOk: () => {
                    store.deleteAccount(record.code)
                }
            })
        } else if (type == '3') {
            store.showDialog(record)
        }
    }

    const items = [
        {key: '1', label: '编辑账户'},
        {key: '2', label: '删除账户'},
        {key: '3', label: '调整余额'},
    ];

    const columns: any = [
        {title: '账户', dataIndex: 'name', key: 'name'},
        {title: '流入', dataIndex: 'debit', key: 'debit', align: 'right', width: '20%'},
        {title: '流出', dataIndex: 'credit', width: '20%', key: 'credit', align: 'right'},
        {title: '余额', dataIndex: 'balance', width: '20%', key: 'balance', align: 'right'},
        {
            title: '操作', dataIndex: 'operator', width: '250px', align: 'center',
            render: (_: any, record: any) =>
                <Space>
                    <a onClick={() => store.showDialog({parentId: record.id})}>新建</a>
                    <Dropdown menu={{items, onClick: e => operatorClick(e.key, record)}}>
                        <a onClick={() => store.showDialog(record)}>更多<DownOutlined/></a>
                    </Dropdown>
                </Space>,
        }];

    const addAccount = () => {
        message.info('添加账户');
    };

    const deleteAccount = () => {
        message.info('删除账户');
    };


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
