import React from 'react';
import {Col, Menu, MenuProps, Row, Space} from "antd";
import {AppstoreOutlined, MailOutlined} from "@ant-design/icons";

const items = [{
  label: '概览',
  key: 'total',
  icon: <MailOutlined/>,
}, {
  label: '记账',
  key: 'app',
  icon: <AppstoreOutlined/>,
}, {
  label: '账户',
  key: 'account',
  icon: <AppstoreOutlined/>,
}, {
  label: '报表',
  key: 'report',
  icon: <AppstoreOutlined/>,
}, {
  label: '设置',
  key: 'setup',
  icon: <AppstoreOutlined/>,
}
];

export default function () {
  return (
    <Row>
      <Col flex="100px">100px</Col>
      <Col flex="auto"><Menu mode="horizontal" theme="light" items={items}/></Col>
    </Row>
  );
};
