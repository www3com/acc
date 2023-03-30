import {Button, Col, DatePicker, Descriptions, Form, Input, notification, Row, Select, Tabs, TreeSelect} from 'antd';
import {inject, observer} from "mobx-react";

export default ({data}: any) => {
    let accountLabel = "分类"
    let cpAccountLabel = "账户"
    let amountLabel = "金额"
    let tradingDateLabel = "时间"
    let projectLabel = "项目"
    let memberLabel = "成员"
    let supplierLabel = "商家"
    let remarkLabel = "备注"
    return (<div style={{paddingLeft: 40}}>
        <Descriptions size='small'>
            <Descriptions.Item label={accountLabel}>{data.account}</Descriptions.Item>
            <Descriptions.Item label={cpAccountLabel}>{data.cpAccount}</Descriptions.Item>
            <Descriptions.Item label={amountLabel}>{data.amount}</Descriptions.Item>
            <Descriptions.Item label={tradingDateLabel}>{data.tradingTime}</Descriptions.Item>
            <Descriptions.Item label={projectLabel}>{data.project}</Descriptions.Item>
            <Descriptions.Item label={memberLabel}>{data.member}</Descriptions.Item>
            <Descriptions.Item label={supplierLabel}>{data.supplier}</Descriptions.Item>
            <Descriptions.Item label={remarkLabel}>{data.remark}</Descriptions.Item>
            <Descriptions.Item><Button type={"primary"} size={"small"}>删除</Button></Descriptions.Item>
        </Descriptions>
    </div>);
};

