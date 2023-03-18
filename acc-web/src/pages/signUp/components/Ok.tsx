import {Button, Result} from "antd";
import {history} from "umi";

import styles from '../style.less'
import Footer from "@/components/Footer"
import Header from "@/pages/signUp/components/Header"
import FullDialog from "@/components/FullDialog";
import {inject, observer} from "mobx-react";


const ok = ({store}: any) => {
    const onClick = () => {
        store.showOk(false)
        history.push("/sign-in")
    }
    return (
        <FullDialog visible={store.okVisible}>
            <div className={styles.header}>
                <Header/>
            </div>
            <div className={styles.content}>
                <div className={styles.contentInner}>
                    <Result
                        status="success"
                        title="恭喜您，注册账户成功！"
                        extra={[
                            <Button type="primary" key={'btn'} onClick={onClick}>返回登录</Button>
                        ]}
                    />
                </div>
            </div>
            <div className={styles.footer}>
                <Footer/>
            </div>
        </FullDialog>
    );
}

export default inject('store')(observer(ok))
