import React, { useEffect } from 'react';
import { Card, Divider, Dropdown, message, Space, Statistic, Table } from 'antd';

import {
  DownOutlined,
  PayCircleOutlined, RedEnvelopeOutlined,
  TransactionOutlined,
} from '@ant-design/icons';
import { inject, observer } from 'mobx-react';
import Dialog from '@/pages/account/component/Dialog';

const root = ({ store }: any) => {

  useEffect(() => {
    store.list(4);
  }, []);

  const items = [
    { key: '1', label: '删除账户' },
    { key: '2', label: '调整余额' },
  ];

  const columns: any = [
    { title: '账户', dataIndex: 'name', key: 'name' },
    { title: '流入', dataIndex: 'debit', key: 'debit', align: 'right', width: '20%' },
    { title: '流出', dataIndex: 'credit', width: '20%', key: 'credit', align: 'right' },
    { title: '余额', dataIndex: 'balance', width: '20%', key: 'balance', align: 'right' },
    {
      title: '操作', dataIndex: 'operator', width: '250px', align: 'center',
      render: (_: any, record: any) =>
        <Space>
          <a onClick={() => store.showDialog({ parentId: record.id })}>新建</a>
          <Dropdown menu={{ items }}>
            <a>编辑 <DownOutlined /></a>
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
    <div style={{ backgroundColor: 'white' }}>
      <Card>
        <Space split={<Divider type='vertical' style={{ height: 50 }} />} size={50}>
          <Statistic title='总资产' value={100000} groupSeparator=',' prefix={<PayCircleOutlined />} />
          <Statistic title='总负债' value={2200} groupSeparator=',' prefix={<TransactionOutlined />} />
          <Statistic title='净资产' value={800000} groupSeparator=',' prefix={<RedEnvelopeOutlined />} />
        </Space>
      </Card>
      <Table
        rowKey={'id'}
        columns={columns}
        dataSource={store.accounts}
        pagination={false} />
      <Dialog />
    </div>
  );
};

export default inject('store')(observer(root));
