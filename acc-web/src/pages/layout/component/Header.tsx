import React, {useEffect, useState} from 'react';
import {Col, Divider, Dropdown, Menu, message, Row, Space} from "antd";
import {
  WalletOutlined,
  AppstoreOutlined,
  DownOutlined,
  AccountBookOutlined,
  SettingOutlined,
  LineChartOutlined, PoweroffOutlined, UserOutlined, PayCircleOutlined, PlusCircleOutlined, PlusOutlined
} from "@ant-design/icons";
import style from "../style.less";
import {clearAll, getSelectLedger, Ledger, listLedgers, selectLedger} from "@/services/ledger";
import {history} from "umi";

const items = [
  {label: '概览', key: 'total', icon: <AppstoreOutlined/>},
  {label: '记账', key: 'app', icon: <AccountBookOutlined/>},
  {label: '账户', key: 'account', icon: <WalletOutlined/>},
  {label: '报表', key: 'report', icon: <LineChartOutlined/>},
  {label: '设置', key: 'setup', icon: <SettingOutlined/>}];

export default () => {

  const [ledgers, setLedgers] = useState([]);
  const [ledger, setLedger] = useState({key: "", label: "", icon: <PlusOutlined/>});

  const init = async () => {
    let ledgerArr = await listLedgers()
    let sLedger = getSelectLedger()

    ledgerArr = ledgerArr.filter(item => item.key != sLedger.key)

    ledgerArr.map((value: Ledger) => {
      value.icon = <PayCircleOutlined/>
      return value
    })

    ledgerArr.push({type: 'divider'})
    ledgerArr.push({key: "new", label: "新建账本", icon: <PlusOutlined/>})
    setLedgers(ledgerArr)
    setLedger({key: sLedger.key, label: sLedger.label, icon: <PlusOutlined/>})
  }

  const onClickLedger = ({key}: any) => {
    if (key == 'new') {
      message.info('暂时不支持新建账本')
      return
    }
    selectLedger(key)
    init()
  };

  const onClickUser = () => {
    message.info("用户设置")
  }

  const onClickQuit = () => {
    clearAll()
    history.push('/login')
  }


  useEffect(() => {
    init()
  }, [])

  const menu = (<Menu onClick={onClickLedger} items={ledgers}/>);

  return (
    <Row className={style.headerRow}>
      <Col flex="100px" style={{margin: 'auto'}}>
        <a href='/sign-in'>
          <img src={'./logo3.png'} width={140} height={40}/>
        </a>
      </Col>
      <Col flex="auto">
        <Menu className={style.menu} mode="horizontal" defaultSelectedKeys={['total']} items={items}/>
      </Col>
      <Col flex="300px" style={{margin: 'auto', textAlign: "right"}}>
        <Space split={<Divider type="vertical"/>} size={0}>
          <Dropdown overlay={menu}>
            <a style={{color: 'black'}}><PayCircleOutlined/> {ledger?.label} <DownOutlined/></a>
          </Dropdown>
          <a style={{color: 'black'}} onClick={onClickUser}><UserOutlined/> Jason</a>
          <a style={{color: 'black'}} onClick={onClickQuit}><PoweroffOutlined/> 退出</a>
        </Space>

      </Col>
    </Row>
  );
};

