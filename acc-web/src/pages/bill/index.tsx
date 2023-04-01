import {App, Col, ConfigProvider, Layout, Row} from 'antd';
import {Provider} from "mobx-react";
import style from "@/pages/bill/style.less";
import {BillStore} from "@/stores/bill";
import TransactionTabs from "@/pages/bill/component/TransactionTabs";
import TransactionList from "@/pages/bill/component/TransactionList";

const store = new BillStore();

export default () => {
    return (
            <Provider store={store}>
                <ConfigProvider>

                    <Row className={style.row} gutter={[0, 4]}>
                        <Col span={24}>
                            <TransactionTabs/>
                        </Col>
                        <Col span={24}>
                            <TransactionList/>
                        </Col>
                    </Row>

                </ConfigProvider>
            </Provider>
    );
}
