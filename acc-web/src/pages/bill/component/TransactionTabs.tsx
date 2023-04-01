import {notification, Tabs} from 'antd';

import Transaction from "@/pages/bill/component/TransactionTab";
import {inject, observer} from "mobx-react";
import {transactionType} from "@/consts";

const transactionTabs = ({store}: any) => {
    const onFinish = async (values: any) => {
        await store.saveTransaction(values);
        notification.success({
            message: '保存成功',
        })
    };

    const expense =
        <Transaction type={transactionType.expenses}
                     account={{label: '分类', options: store.expenses}}
                     cpAccount={{label: '账户', options: store.cpAccounts}}
                     project={{options: store.projects}}
                     member={{options: store.members}}
                     supplier={{options: store.suppliers}}
                     onFinish={onFinish}/>;

    const income =
        <Transaction type={transactionType.income}
                     account={{label: '分类', options: store.incomes}}
                     cpAccount={{label: '账户', options: store.cpAccounts}}
                     project={{options: store.projects}}
                     member={{options: store.members}}
                     supplier={{options: store.suppliers}}
                     onFinish={onFinish}/>;

    const transfer =
        <Transaction type={transactionType.transfer}
                     account={{label: '转出', options: store.cpAccounts}}
                     cpAccount={{label: '转入', options: store.cpAccounts}}
                     project={{options: store.projects}}
                     member={{options: store.members}}
                     supplier={{options: store.suppliers}}
                     onFinish={onFinish}/>;

    const lend =
        <Transaction type={transactionType.lend}
                     account={{label: '负债账户', options: store.debts}}
                     cpAccount={{label: '存入账户', options: store.cpAccounts}}
                     project={{options: store.projects}}
                     member={{disabled: true}}
                     supplier={{disabled: true}}
                     onFinish={onFinish}/>;
    const borrow =
        <Transaction type={transactionType.borrow}
                     account={{label: '转出账户', options: store.debts}}
                     cpAccount={{label: '债权账户', options: store.cpAccounts}}
                     project={{options: store.projects}}
                     member={{disabled: true}}
                     supplier={{disabled: true}}
                     onFinish={onFinish}/>;
    const debtIn =
        <Transaction type={transactionType.debtIn}
                     account={{label: '债权账户', options: store.debts}}
                     cpAccount={{label: '存入账户', options: store.cpAccounts}}
                     project={{options: store.projects}}
                     member={{disabled: true}}
                     supplier={{disabled: true}}
                     onFinish={onFinish}/>;
    const debtOut =
        <Transaction type={transactionType.debtOut}
                     account={{label: '转出账户', options: store.debts}}
                     cpAccount={{label: '负债账户', options: store.cpAccounts}}
                     project={{options: store.projects}}
                     member={{disabled: true}}
                     supplier={{disabled: true}}
                     onFinish={onFinish}/>;
    const items = [
        {label: '支出', key: '1', children: expense},
        {label: '收入', key: '2', children: income},
        {label: '转账', key: '3', children: transfer},
        {label: '借入', key: '4', children: borrow},
        {label: '借出', key: '5', children: lend},
        {label: '收债', key: '6', children: debtIn},
        {label: '还债', key: '7', children: debtOut}
    ];
    return <Tabs size='small' items={items}/>;
};

export default inject('store')(observer(transactionTabs))