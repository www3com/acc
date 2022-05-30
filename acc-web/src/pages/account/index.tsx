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
import {Provider} from "mobx-react";
import {AccountStore} from "@/stores/accountStore";
import Root from "@/pages/account/component/Root";

const store = new AccountStore();


export default () => {
  return (
    <Provider store={store}>
      <ConfigProvider locale={zhCN}>
        <Root/>
      </ConfigProvider>
    </Provider>
  );
}
