import {Divider, Dropdown, Menu, Space} from "antd";
import {history} from "umi";
import {Link} from "@@/exports";
import LedgerDropdown from "@/pages/layout/component/LedgerDropdown";
import {inject, observer} from "mobx-react";
import {clear} from "@/components/token";
import {DownOutlined, SettingOutlined, UserOutlined, PayCircleOutlined, LogoutOutlined} from "@ant-design/icons";

const rightContent = ({store}: any) => {

    const onQuit = () => {
        clear()
        history.push('/login')
    }

    const menu: any = [
        {label: <Link to='/register'>个人中心</Link>, key: 'self', icon: <UserOutlined/>},
        {label: <Link to='/register'>个人设置</Link>, key: 'setup', icon: <SettingOutlined/>},
        {type: 'divider'},
        {label: <a onClick={onQuit}>退出登录</a>, key: 'quit', icon: <LogoutOutlined/>}
    ]

    return (
        <Space split={<Divider type="vertical"/>} size={0}>
            <Dropdown visible={store.visible} overlay={<LedgerDropdown/>}
                      onVisibleChange={visible => store.setVisible(visible)}>
                <a style={{color: 'black'}}><PayCircleOutlined/> {store.currentLedger.name} <DownOutlined/></a>
            </Dropdown>
            <Dropdown overlay={<Menu style={{width: 150}} items={menu}/>}>
                <a style={{color: 'black'}}><UserOutlined/> Jason</a>
            </Dropdown>
        </Space>
    );
};

export default inject('store')(observer(rightContent))