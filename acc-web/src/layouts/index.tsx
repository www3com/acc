import {Row, Col, Layout, Menu, Space, Divider, Typography, Dropdown, message, ConfigProvider} from "antd";
import style from './style.less'
import {DownOutlined, LogoutOutlined, SoundOutlined} from "@ant-design/icons";
import Login from '@/pages/signIn'
import Register from "@/pages/signUp";
import Success from "@/pages/signUp/success";
import zhCN from 'antd/lib/locale/zh_CN';

const items: any = [
  {label: '概览'},
  {label: '账户'},
  {label: '记账'},
  {label: '设置'},
];

const onClick = ({key}: any) => {
  message.info(`Click on item ${key}`);
};

const menu = (
  <Menu
    onClick={onClick}
    items={[
      {
        label: '1st menu item',
        key: '1',
      },
      {
        label: '2nd menu item',
        key: '2',
      },
      {
        label: '3rd menu item',
        key: '3',
      },
    ]}
  />
);

export default (props: any) => {
  switch (props.location.pathname) {
    case '/': return <Login/>
    case '/signIn': return <Login/>
    case '/user': return <Register/>
    case '/user/success': return <Success/>
  }


  return (
    <ConfigProvider locale={zhCN}>
      <div>
        <div className={style.headerWrapper}>
          <Row className={style.header}>
            <Col flex="100px" style={{margin: 'auto'}}>
              <img src={'./logo.png'} width={110} height={30}/>
            </Col>
            <Col flex="auto">
              <Menu className={style.ddd}
                    mode="horizontal"
                    items={items}
              />
            </Col>
            <Col flex="300px" style={{margin: 'auto', textAlign: "right"}}>
              <Space split={<Divider type="vertical"/>}>

                <Dropdown overlay={menu}>
                  <a onClick={e => e.preventDefault()}>
                    <Space>
                      标准账户
                      <DownOutlined/>
                    </Space>
                  </a>
                </Dropdown>
                <Typography.Link>Jason</Typography.Link>
                <SoundOutlined/>
                <LogoutOutlined/>
              </Space>
            </Col>
          </Row>
        </div>
        <div className={style.mainWrapper}>
          <div className={style.main}>
            {props.children}
          </div>
          <div className={style.footer}>
            border
          </div>
        </div>
      </div>
    </ConfigProvider>
  );
}
