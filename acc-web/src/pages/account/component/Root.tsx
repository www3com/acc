import React from 'react';
import {Layout} from 'antd';
import {useModel} from 'umi';

import styles from '../style.less';

import LeftTree from "@/pages/account/component/LeftTree";
import Main from "@/pages/account/component/Main";

const { Content, Sider } = Layout;

export default function ()  {
  return (
    <Layout className={styles.layout}>
      <Sider className={styles.sider} width={350}>
        <LeftTree/>
      </Sider>
      <Content>

        <Main/>
      </Content>
    </Layout>
  );
};
