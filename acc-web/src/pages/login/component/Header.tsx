import style from "@/pages/login/style.less";
import {Col, Row} from "antd";
import {Link} from "umi";
import React from "react";

export default () => {
  return (<Row className={style.headerRow}>
    <Col flex="100px" style={{margin: 'auto'}}>
      <img src={'./logo3.png'} width={170} height={45}/>
    </Col>
    <Col flex="auto" style={{margin: 'auto'}}>
      <Link to='/signUp' style={{float: "right"}}>注册</Link>
    </Col>
  </Row>)
}
