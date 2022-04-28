import React from 'react';
import {Button, Card, Checkbox, Col, Divider, Empty, Form, Input, Layout, Row, Space, Tabs} from 'antd';
import {useModel} from 'umi';
import MonacoEditor from 'react-monaco-editor';
import styles from '../style.less'

const {Header, Content, Footer} = Layout
const {TabPane} = Tabs

export default () => {
  const {setUserData} = useModel("account.account");


  const onFinish = (values: any) => {
    setUserData(values)
    console.log('Success:', values);
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  const options: any = {
    minimap: { // 关闭代码缩略图
      enabled: false // 是否启用预览图
    },
    foldingStrategy: 'indentation', // 代码可分小段折叠
    // cursorSurroundingLines: 0,
    automaticLayout: true,
    selectOnLineNumbers: true,
    scrollbar: {
      verticalScrollbarSize: 8,
      horizontalScrollbarSize: 8,
    },
    renderLineHighlight: false,
    overviewRulerBorder: false,
    scrollBeyondLastLine: false,
    tabSize: 2
  };

  return (
    <Tabs size={'small'} type="editable-card" hideAdd>
      <TabPane tab="Tab 1" key="1">
        <Row style={{marginBottom:8}} gutter={8}>
          <Col span={12}>
            <Input style={{width: '100%'}}></Input>
          </Col>
          <Col>
            <Button type="primary">发送</Button>
          </Col>
        </Row>
        <Row>
          <Col span={12}>
            <Card title='参数' size='small' bodyStyle={{height: "calc(100vh - 300px)"}}>
              <MonacoEditor
                language="json"
                options={options}
                value='{}'
              />
            </Card>
          </Col>
          <Col span={12}>
            <Card title='响应' size='small' bodyStyle={{
              height: "calc(100vh - 300px)",
              width: "100%",
              display: "flex",
              justifyContent: "center",
              alignItems: "center"
            }}>
              <Empty description={'无数据'}/>
            </Card>

          </Col>

        </Row>
      </TabPane>
    </Tabs>
  );
};
