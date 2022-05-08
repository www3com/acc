import React, {useEffect} from 'react';
import {Card, Col, ConfigProvider, Divider, message, Row, Space, Statistic, Switch, Table, Tabs, Tooltip} from 'antd';
import zhCN from 'antd/es/locale/zh_CN';
import {useModel} from 'umi';


import {
  CloseCircleOutlined,
  DollarCircleFilled,
  DollarCircleOutlined,
  DollarCircleTwoTone,
  DollarOutlined, EditOutlined,
  LikeOutlined, PayCircleOutlined, PlusCircleOutlined, PlusOutlined, RedEnvelopeOutlined, TransactionOutlined
} from "@ant-design/icons";


const columns: any = [
  {
    title: '账户',
    dataIndex: 'name',
    key: 'name',
    render: (text: string) => <Space><PayCircleOutlined/>{text}</Space>
  }, {
    title: '流入',
    dataIndex: 'debit',
    key: 'debit',
    align: 'right',
    width: '20%',
    render: (text: string) => '￥' + String(text).replace(/\B(?=(\d{3})+(?!\d))/g, ',')
  }, {
    title: '流出',
    dataIndex: 'credit',
    width: '20%',
    key: 'credit',
    align: 'right',
    render: (text: string) => '￥' + String(text).replace(/\B(?=(\d{3})+(?!\d))/g, ',')
  }, {
    title: '余额',
    dataIndex: 'balance',
    width: '20%',
    key: 'balance',
    align: 'right',
    render: (text: string) => '￥' + String(text).replace(/\B(?=(\d{3})+(?!\d))/g, ',')
  }, {
    title: '操作',
    dataIndex: 'operator',
    width: '250px',
    align: 'center',
    render: () =>
      <Space size={0} split={<Divider type="vertical"/>}>
        <Tooltip title='添加账户'>
          <a onClick={addAccount}><PlusCircleOutlined/></a>
        </Tooltip>
        <Tooltip title='删除账户'>
          <a onClick={deleteAccount}><CloseCircleOutlined/></a>
        </Tooltip>
        <Tooltip title='编辑'>
          <a onClick={addAccount}><EditOutlined/></a>
        </Tooltip>
      </Space>
  }];

const addAccount = () => {
  message.info('添加账户')
}

const deleteAccount = () => {
  message.info('删除账户')
}

export default () => {

  const model = useModel('accountModel');

  useEffect(() => {
    model.list(4)
  }, [])

  return (
    <ConfigProvider locale={zhCN}>
      <div style={{backgroundColor: "white"}}>
        <Card>
          <Space split={<Divider type="vertical" style={{height: 50}}/>} size={50}>
            <Statistic title="总资产" value={100000} groupSeparator=',' prefix={<PayCircleOutlined/>}/>
            <Statistic title="总负债" value={2200} groupSeparator=',' prefix={<TransactionOutlined/>}/>
            <Statistic title="净资产" value={800000} groupSeparator=',' prefix={<RedEnvelopeOutlined/>}/>
          </Space>
        </Card>
        <Table
          rowKey={'id'}
          columns={columns}
          dataSource={model.accounts}
          pagination={false}
        />
      </div>

    </ConfigProvider>
  );
}
