import {Row, Col, Tabs} from "antd";

import React from "react";
import style from "./styles.less";
import AccountForm from './components/AccountForm'
import Header from "./components/Header";
import Footer from "@/components/Footer";
import {Provider} from "mobx-react";
import {Login} from "@/stores/login";

const {TabPane} = Tabs;

const store = new Login()
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
                            <img src={"./login.png"} width='500px' height='350px'/>
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
