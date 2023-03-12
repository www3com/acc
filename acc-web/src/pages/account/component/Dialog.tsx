import { Form, Input, Modal, Radio } from 'antd';
import { inject, observer } from 'mobx-react';

const dialog = ({ store }: any) => {

    const [form] = Form.useForm();

    const handleOk = () => {
      form.validateFields().then((values) => {
        store.saveAccount(values);
        // form.resetFields();
      }).catch((info) => {
        console.log('Validate Failed:', info);
      });
    };

    const title = store.dialogProps.item.id == null ? '新建账户' : '编辑账户';
    return (
      <>
        <Modal title={title} open={store.dialogProps.open} onOk={handleOk} onCancel={() => store.closeDialog()}>
          <Form form={form}
                style={{ paddingRight: 60, paddingTop: 20, paddingBottom: 20 }}
                labelCol={{ span: 6 }}
                wrapperCol={{ span: 18 }}
                name='form_in_modal'
                initialValues={{ modifier: 'public' }}>
            <Form.Item name='name' label='名称'
                       rules={[{ required: true, message: '请输入账户名称!' }]}>
              <Input />
            </Form.Item>
            <Form.Item name='remark' label='描述'>
              <Input type='textarea' />
            </Form.Item>
          </Form>
        </Modal>
      </>
    );
  }
;

export default inject('store')(observer(dialog));
