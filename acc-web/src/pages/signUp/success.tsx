import {Button, Result} from "antd";

import React from "react";
import style from "@/pages/signUp/style.less";
import Header from "@/pages/signUp/component/Header";
import Footer from "@/components/Footer";


export default () => {
  const onClick = () => {
    location.href = '/signIn'
  }
  return (
    <div>
      <div className={style.header}>
        <Header/>
      </div>
      <div className={style.main}>
        <div className={style.mainCenter}>
          <Result
            status="success"
            title="恭喜您，注册账户成功！"
            extra={[
              <Button type="primary" onClick={onClick}>返回登录</Button>
            ]}
          />
        </div>
      </div>
      <div className={style.footer}>
        <Footer/>
      </div>
    </div>
  );
}
