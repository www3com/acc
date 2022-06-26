import {ConfigProvider} from "antd";
import style from './style.less'
import zhCN from 'antd/lib/locale/zh_CN';
import Footer from "@/components/Footer";
import { Link, Outlet } from 'umi';
import Header from "@/pages/layout/component/Header";

export default () => {
  return (
    <ConfigProvider locale={zhCN}>
      <div>
        <div className={style.header}>
            <Header/>
        </div>
        <div className={style.main}>
          <div className={style.mainRow}>
              <Outlet />
          </div>
          <div className={style.footer}>
            <Footer/>
          </div>
        </div>
      </div>
    </ConfigProvider>
  );
}
