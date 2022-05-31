import {Row, Col, Tabs} from "antd";

import React from "react";
import style from "@/pages/signIn/style.less";
import AccountForm from "@/pages/signIn/component/AccountForm";
import Header from "@/pages/signIn/component/Header";
import Footer from "@/components/Footer";
import {Provider} from "mobx-react";
import {UserStore} from "@/stores/userStore";

const {TabPane} = Tabs;

const store = new UserStore()
export default () => {

  return (
    <Provider store={store}>
      <div>
        <div className={style.header}>
          <Header/>
        </div>
        <div className={style.main}>
          <Row className={style.mainRow}>
            <Col span={16}>
              <img src="./login2.png" width='500px' height='350px'/>
            </Col>
            <Col span={8} style={{paddingTop: 30}}>
              <Tabs>
                <TabPane tab="账户登录" key="account">
                  <AccountForm/>
                </TabPane>
              </Tabs>
            </Col>
          </Row>
        </div>
        <div className={style.footer}>
          <Footer/>
        </div>
      </div>
    </Provider>
  );
}
