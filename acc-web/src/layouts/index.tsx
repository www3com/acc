import {Row, Col, Layout, Menu, Space, Divider, Typography, Dropdown, message, ConfigProvider} from "antd";
import style from './style.less'
import {DownOutlined, LogoutOutlined, SoundOutlined} from "@ant-design/icons";
import Login from '@/pages/signIn'
import Register from "@/pages/signUp";
import Success from "@/pages/signUp/success";
import zhCN from 'antd/lib/locale/zh_CN';
import Footer from "@/components/Footer";
import Header from "@/layouts/component/Header";

export default (props: any) => {
  switch (props.location.pathname) {
    case '/': return <Login/>
    case '/sign-in': return <Login/>
    case '/user': return <Register/>
    case '/user/success': return <Success/>
  }

  return (
    <ConfigProvider locale={zhCN}>
      <div>
        <div className={style.header}>
          <Header/>
        </div>
        <div className={style.main}>
          <div className={style.mainRow}>
            {props.children}
          </div>
          <div className={style.footer}>
            <Footer/>
          </div>
        </div>
      </div>
    </ConfigProvider>
  );
}
