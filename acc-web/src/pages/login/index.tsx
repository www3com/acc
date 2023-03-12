import { Row, Col, Tabs, Layout, ConfigProvider } from 'antd';

import React from 'react';
import styles from './styles.less';
import AccountForm from './components/AccountForm';
import Header from './components/Header';
import Footer from '@/components/Footer';
import { Provider } from 'mobx-react';
import { Login } from '@/stores/login';
import { theme } from '@/components/Theme';

const store = new Login();
export default () => {

  const items = [{ label: '账户登录', key: 'account', children: <AccountForm /> }];
  return (
    <Provider store={store}>
      <ConfigProvider theme={theme}>
        <Layout className={styles.layout}>
          <Layout.Header className={styles.header}>
            <Header />
          </Layout.Header>
          <Layout.Content className={styles.content}>
            <Row className={styles.contentInner}>
              <Col span={16}>
                <img src={'./login.png'} width='500px' height='350px' />
              </Col>
              <Col span={8} style={{ paddingTop: 50 }}>
                <Tabs items={items} />
              </Col>
            </Row>
          </Layout.Content>
          <Layout.Footer className={styles.footer}>
            <Footer />
          </Layout.Footer>
        </Layout>
      </ConfigProvider>
    </Provider>
  );
}
