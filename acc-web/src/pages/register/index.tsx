import { ConfigProvider, Layout } from 'antd';
import styles from './style.less';
import Footer from '@/components/Footer';
import Header from '@/pages/register/components/Header';
import Success from '@/pages/register/components/Ok';
import Main from '@/pages/register/components/Main';
import { Provider } from 'mobx-react';
import { Register } from '@/stores/register';
import { theme } from '@/components/Theme';

const store = new Register();

export default () => {
  return (
    <Provider store={store}>
      <ConfigProvider theme={theme}>
        <Layout className={styles.layout}>
          <Layout.Header className={styles.header}>
            <Header />
          </Layout.Header>
          <Layout.Content className={styles.content}>
            <Main />
          </Layout.Content>
          <Layout.Footer className={styles.footer}>
            <Footer />
          </Layout.Footer>
        </Layout>
        <Success />
      </ConfigProvider>
    </Provider>
  );
}


