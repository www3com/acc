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
            date: '2023-01-20',
            type: '支出',
            debitAccount: '吃饭',
            creditAccount: "微信",
            amount: '100.05',
        },
        {
            date: '2023-01-21',
            type: '收入',
            debitAccount: '工资',
            creditAccount: "招行",
            amount: '2104',
        },{
            date: '2023-01-25',
            type: '转账',
            debitAccount: '微信',
            creditAccount: "招行",
            amount: '150',
        },{
            date: '2023-02-25',
            type: '借入',
            debitAccount: '王冠军',
            creditAccount: "招行",
            amount: '200',
        },{
            date: '2023-02-28',
            type: '借出',
            debitAccount: '刘振涛',
            creditAccount: "招行",
            amount: '100',
        },{
            date: '2023-03-10',
            type: '收债',
            debitAccount: '许振涛',
            creditAccount: "支付宝",
            amount: '200',
            remark: '因买车借钱'
        },{
            date: '2023-03-20',
            type: '还债',
            debitAccount: '小欧',
            creditAccount: "招行",
            amount: '400',
        }
    ];

    const columns = [
        {
            title: '日期',
            dataIndex: 'date',
            key: 'date',
            width: '120px',
        },{
            title: '业务类型',
            dataIndex: 'type',
            key: 'type',
            width: '80px',
        }, {
            title: '分类/账户',
            dataIndex: 'debitAccount',
            key: 'debitAccount',
        }, {
            title: '金额',
            dataIndex: 'amount',
            key: 'amount',
        }, {
            title: '账户',
            dataIndex: 'creditAccount',
            key: 'creditAccount',
        }, {
            title: '项目',
            dataIndex: 'project',
            key: 'project',
        }, {
            title: '备注',
            dataIndex: 'remark',
            key: 'remark',
        }
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
