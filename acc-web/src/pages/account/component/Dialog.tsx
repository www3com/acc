import {Form, Input, Modal} from 'antd';
import {inject, observer} from 'mobx-react';
import {useEffect} from "react";

const dialog = ({store}: any) => {


        const [form] = Form.useForm();

        useEffect(() => {
            setTimeout(() => {
                form.resetFields();
            });
            // form.setFieldsValue(store.dialogProps.item)
        }, [store.dialogProps.item])

        const handleOk = () => {
            form.validateFields().then(async (values) => {
                await store.saveAccount(values);
                form.resetFields();
            }).catch((info) => {
                console.log('Validate Failed:', info);
            });
        };

        const handleCancel = () => {
            store.closeDialog()
        }
        const title = store.dialogProps.item.id == null ? '新建账户' : '编辑账户';
        return (<>
            <Modal  title={title} open={store.dialogProps.open} onOk={handleOk}
                   confirmLoading={store.dialogProps.loading}
                   onCancel={() => handleCancel()}>
                <Form  form={form}
                      style={{paddingRight: 60, paddingTop: 20, paddingBottom: 20}}
                      labelCol={{span: 6}}
                      wrapperCol={{span: 18}}
                      name='form_in_modal'
                      initialValues={store.dialogProps.item}>
                    <Form.Item name='id' hidden><Input/></Form.Item>
                    <Form.Item name='parentId' hidden><Input/></Form.Item>
                    <Form.Item name='name' label='名称'
                               rules={[{required: true, message: '请输入账户名称!'}]}>
                        <Input/>
                    </Form.Item>
                    <Form.Item name='remark' label='描述'>
                        <Input type='textarea'/>
                    </Form.Item>
                </Form>
            </Modal>
        </>);
    }
;

export default inject('store')(observer(dialog));
