import React from 'react';
import {Card, Input, Layout, Tooltip, Tree} from 'antd';
import styles from '../style.less'
import {SearchOutlined} from '@ant-design/icons';

import {useModel} from 'umi';

const {Header, Content} = Layout;
const {Search} = Input;
export default () => {
  const treeData = [
    {
      title: 'parent 1',
      key: '0-0',
      children: [
        {
          title: 'leaf',
          key: '0-0-0',
        },
        {
          title: 'leaf',
          key: '0-0-1',
        }, {
          title: 'leaf',
          key: '0-0-2',
        }, {
          title: 'leaf',
          key: '0-0-3',
        }, {
          title: 'leaf',
          key: '0-0-4',
        }, {
          title: 'leaf',
          key: '0-0-5',
        },
      ],
    },
  ];

  return (
    <Layout className={styles.layout}>
      <Header style={{padding: 0, height: 80, backgroundColor: "white"}}>
        <Card size="small" title="Protos" extra={<a href="#">More</a>} bodyStyle={{padding: 3}}>
          <Input size={"small"} placeholder="filter" suffix={<SearchOutlined/>}/>
        </Card>
      </Header>
      <Content className={styles.content}>
        <Tree
          showIcon
          defaultExpandAll
          defaultSelectedKeys={['0-0-0']}
          treeData={treeData}
        />
      </Content>
    </Layout>

  );
}
