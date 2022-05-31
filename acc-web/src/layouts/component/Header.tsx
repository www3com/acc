import React from 'react';
import {Col, Divider, Dropdown, Menu, MenuProps, message, Row, Space, Typography} from "antd";
import {
  WalletOutlined,
  AppstoreOutlined,
  DownOutlined,
  LogoutOutlined,
  AccountBookOutlined,
  SettingOutlined,
  SoundOutlined,
  LineChartOutlined, PlusOutlined
} from "@ant-design/icons";
import style from "@/layouts/style.less";

const onClick = ({key}: any) => {
  message.info(`Click on item ${key}`);
};

const menu = (
  <Menu
    onClick={onClick}
    items={[

      <Divider/>,
      {
        icon: <PlusOutlined />,
        label: '新建账本',
        key: '3',
      },
    ]}
  />
);

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

export default function () {


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
                标准账本
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
