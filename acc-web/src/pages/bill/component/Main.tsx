import {inject, observer} from "mobx-react";
import { Card, Col,  Divider, Row, Space, Table} from "antd";

import Search from "@/pages/bill/component/Search";

const main = ({store}: any) => {

    const dataSource = [
        {
            date: '2023-01-20',
            type: '支出',
            debitAccount: '吃饭',
            creditAccount: "微信",
            amount: '-100.05',
        },
        {
            date: '2023-01-21',
            type: '收入',
            debitAccount: '工资',
            creditAccount: "招行",
            amount: '2104',
        }, {
            date: '2023-01-25',
            type: '转账',
            debitAccount: '微信',
            creditAccount: "招行",
            amount: '-150',
        }, {
            date: '2023-02-25',
            type: '借入',
            debitAccount: '王冠军',
            creditAccount: "招行",
            amount: '200',
        }, {
            date: '2023-02-28',
            type: '借出',
            debitAccount: '刘振涛',
            creditAccount: "招行",
            amount: '-100',
        }, {
            date: '2023-03-10',
            type: '收债',
            debitAccount: '许振涛',
            creditAccount: "支付宝",
            amount: '200',
            remark: '因买车借钱'
        }, {
            date: '2023-03-20',
            type: '还债',
            debitAccount: '小欧',
            creditAccount: "招行",
            amount: '-400',
        }, {
            date: '2023-03-20',
            type: '支出',
            debitAccount: '余额调整',
            creditAccount: "招行",
            amount: '-400',
        }, {
            date: '2023-03-20',
            type: '收入',
            debitAccount: '余额调整',
            creditAccount: "招行",
            amount: '500',
        }
    ];

    const columns = [
        {
            title: '日期',
            dataIndex: 'date',
            key: 'date',
            width: '120px',
            align: 'center'
        }, {
            title: '业务类型',
            dataIndex: 'type',
            key: 'type',
            width: '80px',
            align: 'center'
        }, {
            title: '分类/账户',
            dataIndex: 'debitAccount',
            key: 'debitAccount',
        }, {
            title: '金额',
            dataIndex: 'amount',
            key: 'amount',
            align:'right',
        }, {
            title: '账户',
            dataIndex: 'creditAccount',
            key: 'creditAccount',
        }, {
            title: '项目',
            dataIndex: 'project',
            key: 'project',
        }, {
            title: '备注',
            dataIndex: 'remark',
            key: 'remark',
        }
    ];


    const extra = (
        <Space split={<Divider type="vertical"/>}>
            <span>总支出 <span style={{color: '#14ba89'}}>-11,300.00</span></span>
            <span>总收入 <span style={{color: '#f1523a'}}>+29,100.00</span></span>
            <span>结余 17，800.00 <span style={{color: '#aaa'}}>（单位：元）</span></span>
        </Space>
    )

    return (
        <Card title='账目清单' size={'small'} bordered={false} extra={extra}>
            <Row gutter={[0, 8]}>
                <Col span={24}>
                    <Search/>
                </Col>
                <Col span={24}>
                    <Table size={"small"} dataSource={dataSource} bordered columns={columns}/>
                </Col>
            </Row>

        </Card>

    );
};

export default inject('store')(observer(main))
