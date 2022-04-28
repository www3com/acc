import React from 'react';
import {ConfigProvider} from 'antd';
import zhCN from 'antd/es/locale/zh_CN';
import Root from './component/Root';

export default () => {
  return (
    <ConfigProvider componentSize={'middle'} locale={zhCN}>
      <div>rooo</div>
    </ConfigProvider>
  );
}
