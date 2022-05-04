import {Row, Col} from "antd";

import React from "react";
import style from "@/pages/login/style.less";

export default () => {
  return (
    <div style={{backgroundColor: "white"}}>
      <div className={style.headerWrapper}>
        <Row className={style.header}>
          <Col flex="100px" style={{margin: 'auto'}}>
            <img src={'./logo.png'} width={140} height={50}/>
          </Col>
          <Col flex="auto">

          </Col>
        </Row>
      </div>
      <div className={style.mainWrapper}>
        <div className={style.main}>
            <Row>
              <Col span={12}>
                image
              </Col>
              <Col span={12}>
                input
              </Col>
            </Row>
        </div>
        <div className={style.footer}>
          border
        </div>
      </div>
    </div>
  );
}
