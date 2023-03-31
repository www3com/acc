import {notification, Tabs} from 'antd';

import Expenses from "@/pages/bill/component/Expenses";
import Income from "@/pages/bill/component/Income";
import Transfer from "@/pages/bill/component/Transaction";
import Transaction from "@/pages/bill/component/Transaction";
import {inject, observer} from "mobx-react";

const top = ({store}: any) => {
    const onFinish = async (values: any) => {
        await store.saveTransaction(values);
        notification.success({
            message: '保存成功',
        })
    };

    const transfer =
        <Transaction type={7}
                     account={{label: '转出', options: store.cpAccounts}}
                     cpAccount={{label: '转入', options: store.cpAccounts}}
                     onFinish={onFinish}/>;

    const items = [
        {label: '支出', key: '1', children: <Expenses/>},
        {label: '收入', key: '2', children: <Income/>},
        {label: '转账', key: '3', children: transfer},
        {label: '借入', key: '4', children: 'Content of Tab Pane 4'},
        {label: '借出', key: '5', children: 'Content of Tab Pane 5'},
        {label: '收债', key: '6', children: 'Content of Tab Pane 6'},
        {label: '还债', key: '7', children: 'Content of Tab Pane 7'}
    ];
    return <Tabs size='small' items={items}/>;
};

export default inject('store')(observer(top))