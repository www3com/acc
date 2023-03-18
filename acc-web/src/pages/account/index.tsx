import {ConfigProvider} from 'antd';
import zhCN from 'antd/locale/zh_CN';

import {Provider} from "mobx-react";
import {Account} from "@/stores/account";
import Root from "@/pages/account/component/Root";

const store = new Account();

export default () => {
  return (
    <Provider store={store}>
      {/*<ConfigProvider locale={zhCN}>*/}
        <Root/>
      {/*</ConfigProvider>*/}
    </Provider>
  );
}
