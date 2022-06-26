import {ConfigProvider} from "antd"
import ZH from 'antd/lib/locale/zh_CN'
import style from './style.less';
import Footer from "@/components/Footer";
import Header from "@/pages/register/components/Header";
import Success from "@/pages/register/components/Ok";
import Main from "@/pages/register/components/Main";
import {Provider} from "mobx-react";
import {Register} from "@/stores/register";

const store = new Register();

export default () => {
    return (
        <Provider store={store}>
            <ConfigProvider locale={ZH}>
                <div className={style.header}>
                    <Header/>
                </div>
                <div className={style.main}>
                    <Main/>
                    <div className={style.footer}>
                        <Footer/>
                    </div>
                </div>
                <Success/>
            </ConfigProvider>
        </Provider>
    );
}


