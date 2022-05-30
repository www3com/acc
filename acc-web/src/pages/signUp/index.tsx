import {ConfigProvider} from "antd";

import React from "react";
import zhCN from "antd/lib/locale/zh_CN";
import {Provider} from "mobx-react";
import {UserStore} from "@/stores/userStore";
import Root from "@/pages/signUp/component/Root";

const store = new UserStore()

export default () => {
  return (
    <Provider store={store}>
      <ConfigProvider locale={zhCN}>
        <Root/>
      </ConfigProvider>
    </Provider>
  );
}


