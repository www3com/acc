import {Row, Col, Layout, Menu} from "antd";

const { Header, Content, Footer } = Layout;


export default (props: any) => {
  if (props.location.pathname === '/login') {
    return <div>login</div>
  }

  return (
    <Layout style={{width: 1024, margin:"auto", backgroundColor:'white'}}>
      <Header style={{ position: 'fixed', zIndex: 1, width: 1024, backgroundColor:'white', borderBottom:'1px solid #f0f0f0' }}>
        <Row>
          <Col flex="100px">100px</Col>
          <Col flex="auto">
              <Menu
                mode="horizontal"
                defaultSelectedKeys={['2']}
                items={new Array(3).fill(null).map((_, index) => ({
                  key: String(index + 1),
                  label: `nav ${index + 1}`,
                }))}
              />
          </Col>
        </Row>
      </Header>
      <Content  style={{ padding: '0 50px', marginTop: 80,  minHeight: 380 }}>

          {props.children}

      </Content>
      <Footer style={{ textAlign: 'center' }}>Ant Design ©2018 Created by Ant UED</Footer>
    </Layout>

  );
}
