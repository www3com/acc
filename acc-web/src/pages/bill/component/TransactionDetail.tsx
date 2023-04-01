import {Button, Col, Descriptions, Row} from 'antd';
import {transactionType, transactionTypeName} from "@/consts";

export default ({data}: any) => {
    let accountLabel = "分类"
    let cpAccountLabel = "账户"
    let amountLabel = "金额"
    let tradingDateLabel = "时间"
    let projectLabel = "项目"
    let memberLabel = "成员"
    let supplierLabel = "商家"
    let remarkLabel = "备注"

    let visible = true;

    if (data.type == transactionTypeName.borrow) {
        accountLabel = "转出账户";
        cpAccountLabel = "债权账户";
        visible = false;
    } else if (data.type == transactionTypeName.lend) {
        accountLabel = "负债账户";
        cpAccountLabel = "存入账户";
        visible = false;
    } else if (data.type == transactionTypeName.debtIn) {
        accountLabel = "债权账户";
        cpAccountLabel = "存入账户";
        visible = false;
    } else if (data.type == transactionTypeName.debtOut) {
        accountLabel = "转出账户";
        cpAccountLabel = "负债账户";
        visible = false;
    }

    return (<div style={{paddingLeft: 40}}>
        <Row>
            <Col flex={"auto"}>
                <Descriptions size='small' column={4}>
                    <Descriptions.Item label={accountLabel}>{data.account}</Descriptions.Item>
                    <Descriptions.Item label={cpAccountLabel}>{data.cpAccount}</Descriptions.Item>
                    <Descriptions.Item label={amountLabel}>{data.amount}</Descriptions.Item>
                    <Descriptions.Item label={tradingDateLabel}>{data.tradingTime}</Descriptions.Item>
                    <Descriptions.Item label={projectLabel}>{data.project}</Descriptions.Item>
                    {visible ? <Descriptions.Item label={memberLabel}>{data.member}</Descriptions.Item> : <></>}
                    {visible ? <Descriptions.Item label={supplierLabel}>{data.supplier}</Descriptions.Item> : <></>}
                    <Descriptions.Item label={remarkLabel}>{data.remark}</Descriptions.Item>
                </Descriptions>
            </Col>
            <Col flex='50px'>
                <Button
                    type={"primary"} size={"small"}>删除</Button>
            </Col>
        </Row>

    </div>);
};

