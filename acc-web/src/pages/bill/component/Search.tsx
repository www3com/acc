import {
    Calendar,
    Col,
    Row,
    Space,
    Tree,
} from 'antd';

import React, {Key, useState} from "react";
import {
    ArrowRightOutlined,
    ClearOutlined,
} from "@ant-design/icons";
import SearchDisplay from "@/pages/bill/component/SearchDisplay";
import SearchItem from "@/pages/bill/component/SearchItem";
import {inject, observer} from "mobx-react";
import dayjs, {Dayjs} from "dayjs";

const search = ({store}: any) => {

    // const [startDate, setStartDate] = useState<Dayjs>(dayjs())
    // const [endDate, setEndDate] = useState<Dayjs>(dayjs())
    const [dateName, setDateName] = useState<string>('')

    // const [accounts, setAccounts] = useState<any[]>([])
    const [accountNames, setAccountNames] = useState<string[]>([])

    // const [cpAccounts, setCpAccounts] = useState<any[]>([])
    const [cpAccountNames, setCpAccountNames] = useState<string[]>([])

    // const [projects, setProjects] = useState<any[]>([])
    const [projectNames, setProjectNames] = useState<string[]>([])

    // const [members, setMembers] = useState<any[]>([])
    const [memberNames, setMemberNames] = useState<string[]>([])

    // const [suppliers, setSuppliers] = useState<any[]>([])
    const [supplierNames, setSupplierNames] = useState<string[]>([])


    const translate = (nodes: any[]) => {
        let keys: string[] = [];
        for (const node of nodes) {
            if (node.name != "") {
                keys.push(node.name)
            }
        }
        return keys;
    }

    const onSelectDate = () => {
        if (store.params.startDate == null) {
            store.params.startDate = dayjs()
        }
        if (store.params.endDate == null) {
            store.params.endDate = dayjs()
        }
        return setDateName(store.params.startDate.format('YYYY-MM-DD') + " 至 " + store.params.endDate.format('YYYY-MM-DD'))
    }

    const onSelect = (label: string, nodes: any[]) => {
        switch (label) {
            case 'account':
                store.params.accounts = nodes;
                break;
            case 'cpAccount':
                store.params.cpAccounts = nodes;
                // setCpAccounts(nodes);
                break;
            case 'project':
                store.params.projects = nodes;
                // setProjects(nodes);
                break;
            case 'member':
                store.params.members = nodes;
                // setMembers(nodes);
                break;
            case 'supplier':
                store.params.suppliers = nodes;
            // setSuppliers(nodes);
        }

    }

    const onClose = (label: string) => {
        switch (label) {
            case 'date':
                setDateName('')
                store.setParams({startDate: dayjs(), endDate: dayjs()})
                return;
            case 'account':
                setAccountNames([]);
                store.setParams({accounts: []})
                return;
            case 'cpAccount':
                setCpAccountNames([]);
                store.setParams({cpAccounts: []})
                return;
            case 'project':
                setProjectNames([]);
                store.setParams({projects: []})
                return;
            case 'member':
                setMemberNames([]);
                store.setParams({members: []})
                return;
            case 'supplier':
                setSupplierNames([]);
                store.setParams({suppliers: []})
        }
    }

    const onCloseAll = () => {
        onClose('date')
        onClose('account')
        onClose('cpAccount')
        onClose('project')
        onClose('member')
        onClose('supplier')
    }

    return <div>
        <Row wrap={false} align='middle'>
            <Col flex="auto">
                <Space>
                    <SearchItem title='选择时间' onOk={() => onSelectDate()} bodyStyle={{width: '600px'}}>
                        <Space>
                            <Calendar fullscreen={false} value={store.params.startDate}
                                      onSelect={(date: Dayjs) => store.setParams({startDate: date})}/>
                            <ArrowRightOutlined/>
                            <Calendar fullscreen={false} value={store.params.endDate}
                                      onSelect={(date: Dayjs) => store.setParams({endDate: date})}/>
                        </Space>
                    </SearchItem>

                    <SearchItem title='全部分类' onOk={() => setAccountNames(translate(store.params.accounts))}>
                        <Tree checkable fieldNames={{title: 'name', key: 'id'}} treeData={store.accounts}
                              checkedKeys={store.params.accounts.map((value: any) => value.id)}
                              onCheck={(selectedKey: any, e) => onSelect('account', e.checkedNodes)}/>
                    </SearchItem>
                    <SearchItem title='全部账户' onOk={() => setCpAccountNames(translate(store.params.cpAccounts))}>
                        <Tree checkable fieldNames={{title: 'name', key: 'id'}} treeData={store.cpAccounts}
                              checkedKeys={store.params.cpAccounts.map((value: any) => value.id)}
                              onCheck={(selectedKey: any, e) => onSelect('cpAccount', e.checkedNodes)}/>
                    </SearchItem>
                    <SearchItem title='全部项目' onOk={() => setProjectNames(translate(store.params.projects))}>
                        <Tree checkable fieldNames={{title: 'name', key: 'id'}} treeData={store.projects}
                              checkedKeys={store.params.projects.map((value: any) => value.id)}
                              onCheck={(selectedKey: any, e) => onSelect('project', e.checkedNodes)}/>
                    </SearchItem>
                    <SearchItem title='全部成员' onOk={() => setMemberNames(translate(store.params.members))}>
                        <Tree checkable fieldNames={{title: 'name', key: 'id'}} treeData={store.members}
                              checkedKeys={store.params.members.map((value: any) => value.id)}
                              onCheck={(selectedKey: any, e) => onSelect('member', e.checkedNodes)}/>
                    </SearchItem>
                    <SearchItem title='全部商家' onOk={() => setSupplierNames(translate(store.params.suppliers))}>
                        <Tree checkable fieldNames={{title: 'name', key: 'id'}} treeData={store.suppliers}
                              checkedKeys={store.params.suppliers.map((value: any) => value.id)}
                              onCheck={(selectedKey: any, e) => onSelect('supplier', e.checkedNodes)}/>
                    </SearchItem>
                </Space>
            </Col>
            <Col flex="34px">
                <a onClick={() => onCloseAll()}><ClearOutlined/></a>
            </Col>
        </Row>

        <div style={{marginLeft: 15}}>
            <SearchDisplay title='选择时间' data={dateName && dateName != '' ? [dateName] : []}
                           onClose={() => onClose('date')}></SearchDisplay>
            <SearchDisplay title='所有分类' data={accountNames} onClose={() => onClose('account')}></SearchDisplay>
            <SearchDisplay title='所有账户' data={cpAccountNames} onClose={() => onClose('cpAccount')}></SearchDisplay>
            <SearchDisplay title='全部项目' data={projectNames} onClose={() => onClose('project')}></SearchDisplay>
            <SearchDisplay title='全部成员' data={memberNames} onClose={() => onClose('member')}></SearchDisplay>
            <SearchDisplay title='全部商家' data={supplierNames} onClose={() => onClose('supplier')}></SearchDisplay>
        </div>

    </div>;
};

export default inject('store')(observer(search))