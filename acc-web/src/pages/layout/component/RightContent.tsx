import { Button, Divider, Dropdown, Menu, MenuProps, Space } from 'antd';
import { history } from 'umi';
import { Link } from '@@/exports';
import { inject, observer } from 'mobx-react';
import { clear, getCurrentLedger } from '@/components/session';
import {
  DownOutlined,
  SettingOutlined,
  UserOutlined,
  PayCircleOutlined,
  LogoutOutlined,
  AppstoreOutlined, BookOutlined,
} from '@ant-design/icons';
import styles from '../style.less';

const rightContent = ({ layoutStore }: any) => {

  const renderItems = (): any => {
    let items: any = [];
    for (const ledger of layoutStore.ledgers) {
      items.push({ key: ledger.id, label: ledger.name });
    }
    return items;
  };

  const onQuit = () => {
    clear();
    history.push('/login');
  };

  const items: any = [
    { label: <Link to='/register'>个人中心</Link>, key: 'self', icon: <UserOutlined /> },
    { label: <Link to='/register'>个人设置</Link>, key: 'setup', icon: <SettingOutlined /> },
    { type: 'divider' },
    { label: <a onClick={onQuit}>退出登录</a>, key: 'quit', icon: <LogoutOutlined /> },
  ];

  let ledger = getCurrentLedger();
  return (
    <Space split={<Divider type='vertical' />} size={0}>
      <Dropdown
        menu={{ items: renderItems(), selectable: true, defaultSelectedKeys: ledger.id.toString() }}
        dropdownRender={(menu) => (
          <div className={styles.dropdownContent} style={{ backgroundColor: 'white' }}>
            {menu}
            <Divider style={{ margin: 0 }} />
            <Space style={{ padding: 8 }}>
              <Button type='text' icon={<AppstoreOutlined />}>账本模版</Button>
              <Button type='text' icon={<BookOutlined />}>账本管理</Button>
            </Space>
          </div>
        )}
      >
        <a style={{ color: 'black' }}><PayCircleOutlined /> {ledger.name} <DownOutlined /></a>
      </Dropdown>
      <Dropdown menu={{ items }} overlayStyle={{ width: 150 }}>
        <a style={{ color: 'black' }}><UserOutlined /> Jason</a>
      </Dropdown>
    </Space>
  );
};

export default inject('layoutStore')(observer(rightContent));