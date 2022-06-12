import React, {useEffect, useState} from 'react';
import {Col, Divider, Dropdown, Menu, message, Row, Space, Typography} from "antd";
import {
  WalletOutlined,
  AppstoreOutlined,
  DownOutlined,
  AccountBookOutlined,
  SettingOutlined,
  LineChartOutlined
} from "@ant-design/icons";
import style from "@/layouts/style.less";
import {getSelectLedger, listLedgers, selectLedger} from "@/services/ledger";


const items = [{
  label: '概览',
  key: 'total',
  icon: <AppstoreOutlined/>,
}, {
  label: '记账',
  key: 'app',
  icon: <AccountBookOutlined/>,
}, {
  label: '账户',
  key: 'account',
  icon: <WalletOutlined/>,
}, {
  label: '报表',
  key: 'report',
  icon: <LineChartOutlined/>,
}, {
  label: '设置',
  key: 'setup',
  icon: <SettingOutlined/>,
}
];

export default () => {

  const [ledgers, setLedgers] = useState([]);
  const [ledger, setLedger] = useState({key: "", label: ""});

  const init = async () => {
    let ledgerArr = await listLedgers()
    let sLedger = getSelectLedger()

    ledgerArr = ledgerArr.filter(item => item.key != sLedger.key)

    setLedgers(ledgerArr)
    setLedger(sLedger)
  }

  const onClick = ({key}: any) => {
    selectLedger(key)
    init()
  };

  useEffect(() => {
    init()
  }, [])

  const menu = (<Menu onClick={onClick} items={ledgers}/>);

  return (
    <Row className={style.headerRow}>
      <Col flex="100px" style={{margin: 'auto'}}>
        <a href='/sign-in'>
          <img src={'./logo3.png'} width={140} height={40}/>
        </a>
      </Col>
      <Col flex="auto">
        <Menu className={style.menu}
              mode="horizontal"
              items={items}
        />
      </Col>
      <Col flex="300px" style={{margin: 'auto', textAlign: "right"}}>
        <Space split={<Divider type="vertical"/>} size={3}>
          <Dropdown overlay={menu}>
            <a onClick={e => e.preventDefault()}>
              <Space>
                {ledger.label}
                <DownOutlined/>
              </Space>
            </a>
          </Dropdown>
          <Typography.Link>Jason</Typography.Link>
        </Space>
      </Col>
    </Row>
  );
};

