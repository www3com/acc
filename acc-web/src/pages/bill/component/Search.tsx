import {Col, Row, Space} from 'antd';
import React, {FC, useRef, useState} from "react";
import {ClearOutlined} from "@ant-design/icons";
import SearchResult from "@/pages/bill/component/SearchResult";
import {inject, observer} from "mobx-react";
import {Dayjs} from "dayjs";
import SearchTree, {SearchTreeHandler} from "@/pages/bill/component/SearchTree";
import SearchDate, {SearchDateHandler} from "@/pages/bill/component/SearchDate";

const search: FC = ({store}: any) => {
    const dateRef = useRef<SearchDateHandler>(null);
    const accountRef = useRef<SearchTreeHandler>(null);
    const cpAccountRef = useRef<SearchTreeHandler>(null);
    const projectRef = useRef<SearchTreeHandler>(null);
    const memberRef = useRef<SearchTreeHandler>(null);
    const supplierRef = useRef<SearchTreeHandler>(null);

    const onClose = (label: string) => {
        switch (label) {
            case 'date':
                dateRef.current?.reset();
                store.setParams({startTime: null, endTime: null});
                break;
            case 'account':
                accountRef.current?.reset();
                store.setParams({accounts: []})
                break;
            case 'cpAccount':
                cpAccountRef.current?.reset();
                store.setParams({cpAccounts: []})
                break;
            case 'project':
                projectRef.current?.reset();
                store.setParams({projects: []})
                break;
            case 'member':
                memberRef.current?.reset();
                store.setParams({members: []})
                break;
            case 'supplier':
                supplierRef.current?.reset();
                store.setParams({suppliers: []})
        }
    }

    const onCloseAll = () => {
        dateRef.current?.reset();
        accountRef.current?.reset();
        cpAccountRef.current?.reset();
        projectRef.current?.reset();
        memberRef.current?.reset();
        supplierRef.current?.reset();
        store.setParams({
            startTime: null,
            endTime: null,
            accounts: [],
            cpAccounts: [],
            projects: [],
            members: [],
            suppliers: []
        });

    }
    let dateName = store.params.startTime == null ? "" :
        store.params.startTime.format('YYYY-MM-DD') + " 至 " + store.params.endTime.format('YYYY-MM-DD')
    return <div>
        <Row wrap={false} align='middle'>
            <Col flex="auto">
                <Space>
                    <SearchDate title='选择时间' ref={dateRef}
                                bodyStyle={{width: '600px'}}
                                onOk={(startTime: Dayjs, endTime: Dayjs) => {
                                    store.setParams({startTime: startTime, endTime: endTime})
                                }}/>
                    <SearchTree title='全部分类' ref={accountRef} treeData={store.accounts}
                                onOk={(value: any) => store.setParams({accounts: value})}/>
                    <SearchTree title='全部账户' ref={cpAccountRef} treeData={store.cpAccounts}
                                onOk={(value: any) => store.setParams({cpAccounts: value})}/>
                    <SearchTree title='全部项目' ref={projectRef} treeData={store.projects}
                                onOk={(value: any) => store.setParams({projects: value})}/>
                    <SearchTree title='全部成员' ref={memberRef} treeData={store.members}
                                onOk={(value: any) => store.setParams({members: value})}/>
                    <SearchTree title='全部商家' ref={supplierRef} treeData={store.suppliers}
                                onOk={(value: any) => store.setParams({suppliers: value})}/>
                </Space>
            </Col>
            <Col flex="34px">
                <a onClick={() => onCloseAll()}><ClearOutlined/></a>
            </Col>
        </Row>

        <div style={{marginLeft: 15}}>
            <SearchResult title='选择时间' data={dateName && dateName != '' ? [dateName] : []}
                          onClose={() => onClose('date')}></SearchResult>
            <SearchResult title='所有分类' data={store.params.accounts.map((value: any) => value.name)}
                          onClose={() => onClose('account')}></SearchResult>
            <SearchResult title='所有账户' data={store.params.cpAccounts.map((value: any) => value.name)}
                          onClose={() => onClose('cpAccount')}></SearchResult>
            <SearchResult title='全部项目' data={store.params.projects.map((value: any) => value.name)}
                          onClose={() => onClose('project')}></SearchResult>
            <SearchResult title='全部成员' data={store.params.members.map((value: any) => value.name)}
                          onClose={() => onClose('member')}></SearchResult>
            <SearchResult title='全部商家' data={store.params.suppliers.map((value: any) => value.name)}
                          onClose={() => onClose('supplier')}></SearchResult>
        </div>

    </div>;
};

export default inject('store')(observer(search))