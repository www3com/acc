import {Row, Col, Tabs} from "antd";

import React from "react";
import style from "@/pages/login/style.less";
import AccountForm from "@/pages/login/component/AccountForm";
import EmailForm from "@/pages/login/component/EmailForm";
import Header from "@/pages/login/component/Header";
import Footer from "@/pages/components/Footer";

const {TabPane} = Tabs;

export default () => {

  return (
    <div>
      <div className={style.header}>
        <Header/>
      </div>
      <div className={style.main}>
        <Row className={style.mainRow}>
          <Col span={16}>
            <img src="./login2.png" width='500px' height='350px'/>
          </Col>
          <Col span={8}>
            <Tabs>
              <TabPane tab="账户登录" key="account">
                <AccountForm/>
              </TabPane>
              <TabPane tab="邮箱登录" key="email">
                <EmailForm/>
              </TabPane>
            </Tabs>
          </Col>
        </Row>
      </div>
      <div className={style.footer}>
        <Footer/>
      </div>
    </div>
  );
}
