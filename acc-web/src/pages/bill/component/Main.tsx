import {inject, observer} from "mobx-react";
import {Card, Col, Divider, Row, Space, Table} from "antd";

import Search from "@/pages/bill/component/Search";
import type {ColumnsType} from 'antd/es/table';
import TransactionDetail from "@/pages/bill/component/TransactionDetail";
import RemarkPopover from "@/pages/account/component/RemarkPopover";
import React from "react";

const main = ({store}: any) => {
    const columns: ColumnsType<any> = [
        {title: '日期', dataIndex: 'tradingTime', key: 'tradingTime', width: '120px', align: 'center'},
        {title: '业务类型', dataIndex: 'type', key: 'type', width: '80px', align: 'center'},
        {title: '分类/账户', dataIndex: 'account', key: 'account',},
        {
            title: '金额', dataIndex: 'amount', key: 'amount', align: 'right', render: (text: string,record:any) => {
                if (record.type == "支出") {
                    return <span style={{color: '#f1523a'}}>{text}</span>
                } else if (record.type == "收入") {
                    return <span style={{color: '#14ba89'}}>{text}</span>
                } else {
                    return <span>{text}</span>
                }
            }
        },
        {title: '账户', dataIndex: 'cpAccount', key: 'cpAccount',},
        {title: '项目', dataIndex: 'project', key: 'project',},
        {title: '备注', dataIndex: 'remark', key: 'remark',}];

    const extra = (
        <Space split={<Divider type="vertical"/>}>
            <span>总支出 <span style={{color: '#f1523a'}}>{store.totalTransaction.expense}</span></span>
            <span>总收入 <span style={{color: '#14ba89'}}>{store.totalTransaction.income}</span></span>
            <span>结余 {store.totalTransaction.balance} <span style={{color: '#aaa'}}>（单位：元）</span></span>
        </Space>
    )

    return (
        <Card title='账目清单' size={'small'} bordered={false} extra={extra}>
            <Row gutter={[0, 8]}>
                <Col span={24}>
                    <Search/>
                </Col>
                <Col span={24}>
                    <Table size={"small"} rowKey={'id'} bordered columns={columns}
                           dataSource={store.transactions}
                           expandable={{
                               expandedRowRender: record => <TransactionDetail data={record}/>
                           }}/>
                </Col>
            </Row>
        </Card>
    );
};

export default inject('store')(observer(main))
