import {Button, Dropdown, Space, Tabs, TreeSelect} from 'antd';

import Expenses from "@/pages/bill/component/Expenses";
import React from "react";
import {LightFilter, ProFormFieldSet, ProFormSelect, ProFormText, ProFormTreeSelect} from "@ant-design/pro-components";
import {FilterOutlined} from "@ant-design/icons";

export default () => {
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

    return <Space>
        <LightFilter size={"middle"}  collapseLabel={<FilterOutlined/>}>
            <ProFormTreeSelect
                name="category"
                label="分类"
                request={async () => treeData}
                fieldProps={{
                    fieldNames: {
                        label: 'title',
                    },
                    treeCheckable: true,
                    showCheckedStrategy: TreeSelect.SHOW_PARENT,
                    placeholder: 'Please select',
                }}
            />
            <ProFormTreeSelect
                name="account"
                label="账户"
                request={async () => treeData}
                fieldProps={{
                    fieldNames: {
                        label: 'title',
                    },
                    treeCheckable: true,
                    showCheckedStrategy: TreeSelect.SHOW_PARENT,
                    placeholder: 'Please select',
                }}
            />
            <ProFormSelect
                name="project"
                label='项目'
                valueEnum={{
                    man: '男',
                    woman: '女',
                }}
                placeholder="性别"
            />
            <ProFormSelect
                name="member"
                label='成员'
                valueEnum={{
                    man: '男',
                    woman: '女',
                }}
                placeholder="性别"
            />
            <ProFormSelect
                name="supplier"
                label='商家'
                valueEnum={{
                    man: '男',
                    woman: '女',
                }}
                placeholder="性别"
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
    </Space>;
};

