import {Col, Menu, Row} from "antd";
import {
    WalletOutlined,
    AppstoreOutlined,
    AccountBookOutlined,
    SettingOutlined,
    LineChartOutlined
} from "@ant-design/icons";

import RightContent from "@/pages/layout/component/RightContent";

const items = [
    {label: '概览', key: 'total', icon: <AppstoreOutlined/>},
    {label: '记账', key: 'bill', icon: <AccountBookOutlined/>},
    {label: '账户', key: 'account', icon: <WalletOutlined/>},
    {label: '报表', key: 'report', icon: <LineChartOutlined/>},
    {label: '设置', key: 'settings', icon: <SettingOutlined/>}];

export default () => {
    return (
        <Row>
            <Col flex="100px" style={{margin: 'auto'}}>
                <a href='/sign-in'>
                    <img src={'./logo3.png'} width={140} height={40} alt=''/>
                </a>
            </Col>
            <Col flex="auto">
                <Menu mode="horizontal" defaultSelectedKeys={['account']} items={items}/>
            </Col>
            <Col flex="300px" style={{margin: 'auto', textAlign: "right"}}>
                <RightContent/>
            </Col>
        </Row>
    );
};

