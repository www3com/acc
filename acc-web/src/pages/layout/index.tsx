import {ConfigProvider, Layout} from "antd";
import style from './style.less'
import zhCN from 'antd/lib/locale/zh_CN';
import Footer1 from "@/components/Footer";
import {Outlet} from 'umi';
import Top from "@/pages/layout/component/Top";
import {LayoutStore} from "@/stores/layout";
import {Provider} from "mobx-react";



const {Header, Content, Footer} = Layout
const store = new LayoutStore()

export default () => {

    return (
        <Provider store={store}>
            <ConfigProvider locale={zhCN}>
                <Layout className={style.layout}>
                    <Header className={style.header}>
                        <Top/>
                    </Header>
                    <Layout className={style.content}>
                        <Content>
                            <Outlet/>
                        </Content>
                        <Footer className={style.footer}>
                            <Footer1/>
                        </Footer>
                    </Layout>
                </Layout>
            </ConfigProvider>
        </Provider>
    );
}