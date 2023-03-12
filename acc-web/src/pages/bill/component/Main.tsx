import {inject, observer} from "mobx-react";
import {Button, Card, Col, DatePicker, Divider, Row, Space, Table, TreeSelect} from "antd";
import {
    LightFilter,
    ProFormSelect,
    ProFormTreeSelect,
    ProFormDateRangePicker,
    ProFormText,
    ProFormFieldSet,
    ProFormRadio
} from "@ant-design/pro-components";
import {FilterOutlined} from "@ant-design/icons";
// import moment from "moment";

const {RangePicker} = DatePicker;

const main = ({store}: any) => {

    const dataSource = [
        {
            key: '1',
            name: '胡彦斌',
            age: 32,
            address: '西湖区湖底公园1号',
        },
        {
            key: '2',
            name: '胡彦祖',
            age: 42,
            address: '西湖区湖底公园1号',
        },
    ];

    const columns = [
        {
            title: '姓名',
            dataIndex: 'name',
            key: 'name',
        },
        {
            title: '年龄',
            dataIndex: 'age',
            key: 'age',
        },
        {
            title: '住址',
            dataIndex: 'address',
            key: 'address',
        },
    ];

    const treeData = [
        {
            title: 'Node1',
            value: '0-0',
            key: '0-0',
            children: [
                {
                    title: 'Child Node1',
                    value: '0-0-0',
                    key: '0-0-0',
                },
            ],
        },
        {
            title: 'Node2',
            value: '0-1',
            key: '0-1',
            children: [
                {
                    title: 'Child Node3',
                    value: '0-1-0',
                    key: '0-1-0',
                },
                {
                    title: 'Child Node4',
                    value: '0-1-1',
                    key: '0-1-1',
                },
                {
                    title: 'Child Node5',
                    value: '0-1-2',
                    key: '0-1-2',
                },
                {
                    title: 'Child Node6',
                    value: '0-1-3',
                    key: '0-1-3',
                },
                {
                    title: 'Child Node7',
                    value: '0-1-4',
                    key: '0-1-4',
                },
                {
                    title: 'Child Node8',
                    value: '0-1-5',
                    key: '0-1-5',
                },
                {
                    title: 'Child Node9',
                    value: '0-1-6',
                    key: '0-1-6',
                },
                {
                    title: 'Child Node10',
                    value: '0-1-7',
                    key: '0-1-7',
                },
                {
                    title: 'Child Node11',
                    value: '0-1-8',
                    key: '0-1-8',
                },
                {
                    title: 'Child Node12',
                    value: '0-1-9',
                    key: '0-1-9',
                },
                {
                    title: 'Child Node13',
                    value: '0-1-10',
                    key: '0-1-10',
                },
            ],
        },
    ];

    const extra = (
        <Space split={<Divider type="vertical"/>}>
            <span>总支出 <span style={{color: '#14ba89'}}>-11,300.00</span></span>
            <span>总收入 <span style={{color: '#f1523a'}}>+29,100.00</span></span>
            <span>结余 17，800.00 <span style={{color: '#aaa'}}>（单位：元）</span></span>
        </Space>
    )

    return (
        <Card title='账目清单' size={'small'} bordered={false} extra={extra}>
            <Row gutter={[0, 10]}>
                <Col span={24}>
                    <Space>
                        <LightFilter size={"middle"} bordered collapseLabel={<FilterOutlined/>}>
                            <ProFormSelect
                                name="sex1"
                                label='选择分类'
                                valueEnum={{
                                    man: '男',
                                    woman: '女',
                                }}
                                placeholder="性别"
                            />
                            <ProFormSelect
                                name="sex2"
                                label='所有账户'
                                valueEnum={{
                                    man: '男',
                                    woman: '女',
                                }}
                                placeholder="性别"
                            />
                            <ProFormSelect
                                name="sex3"
                                label='所有成员'
                                valueEnum={{
                                    man: '男',
                                    woman: '女',
                                }}
                                placeholder="性别"
                            />
                            <ProFormSelect
                                name="sex4"
                                label='所有项目'
                                valueEnum={{
                                    man: '男',
                                    woman: '女',
                                }}
                                placeholder="性别"
                            />
                            <ProFormTreeSelect
                                initialValue={['0-0', '0-1']}
                                label="所有商家"
                                request={async () => treeData}
                                fieldProps={{
                                    fieldNames: {
                                        label: 'title',
                                    },
                                    treeCheckable: true,
                                    showCheckedStrategy: TreeSelect.SHOW_PARENT,
                                    placeholder: 'Please select',
                                }}
                                name="treeSelect"
                            />
                            <ProFormFieldSet name='aaaa' label='日期'>
                                {/*<RangePicker*/}
                                {/*    format='YYYY-MM-DD'*/}
                                {/*    ranges={{*/}
                                {/*        '本周': [moment().startOf('week'), moment()],*/}
                                {/*        '本月': [moment().startOf('month'), moment()],*/}
                                {/*        '本季': [moment().startOf('quarter'), moment()],*/}
                                {/*        '本年': [moment().startOf('year'), moment()],*/}
                                {/*    }}*/}
                                {/*/>*/}
                            </ProFormFieldSet>
                            <ProFormText name='remark' label='备注'/>
                        </LightFilter>
                        <Button>重置</Button>
                    </Space>
                </Col>
                <Col span={24}>
                    <Table size={"small"} dataSource={dataSource} columns={columns}/>
                </Col>
            </Row>

        </Card>

    );
};

export default inject('store')(observer(main))
