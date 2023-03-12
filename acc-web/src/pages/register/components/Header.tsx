import {Col, Row} from "antd"
import {Link} from "umi"
import style from '../style.less'
export default () => {
    return (<Row className={style.headerInner}>
        <Col flex="100px" style={{margin: 'auto'}}>
            <img src={'/logo3.png'} width={170} height={45}/>
        </Col>
        <Col flex="auto" style={{margin: 'auto'}}>
            <Link to='/login' style={{float: "right"}}>登录</Link>
        </Col>
    </Row>)
}
