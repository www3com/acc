import {Col, ConfigProvider, Layout, Row} from 'antd';
import zhCN from 'antd/locale/zh_CN';
import {Provider} from "mobx-react";
import style from "@/pages/bill/style.less";
import {BillStore} from "@/stores/bill";
import Top from "@/pages/bill/component/Top";
import Main from "@/pages/bill/component/Main";
import { theme } from '@/components/Theme';

const store = new BillStore();

const {Header, Content} = Layout

export default () => {
    return (
        <Provider store={store}>
            <ConfigProvider>
                <Row className={style.row}>
                    <Col span={24} style={{padding: '0px 12px'}}>
                        <Top/>
                    </Col>
                    <Col span={24}>
                        <Main/>
                    </Col>
                </Row>
            </ConfigProvider>
        </Provider>
    );
}
