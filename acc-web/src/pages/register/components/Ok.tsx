import {Button, Result} from "antd";
import {history} from "umi";

import style from '../style.less'
import Footer from "@/components/Footer"
import Header from "@/pages/register/components/Header"
import FullDialog from "@/components/FullDialog";
import {inject, observer} from "mobx-react";


const ok = ({store}: any) => {
    const onClick = () => {
        store.showOk(false)
        history.push("/login")
    }
    return (
        <FullDialog visible={store.okVisible}>
            <div className={style.header}>
                <Header/>
            </div>
            <div className={style.main}>
                <div className={style.mainCenter}>
                    <Result
                        status="success"
                        title="恭喜您，注册账户成功！"
                        extra={[
                            <Button type="primary" key={'btn'} onClick={onClick}>返回登录</Button>
                        ]}
                    />
                </div>
            </div>
            <div className={style.footer}>
                <Footer/>
            </div>
        </FullDialog>
    );
}

export default inject('store')(observer(ok))
