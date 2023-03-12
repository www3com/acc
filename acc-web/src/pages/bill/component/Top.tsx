import {Tabs} from 'antd';

import Expenses from "@/pages/bill/component/Expenses";

export default () => {
    const items = [
        {label: '支出', key: '1', children: <Expenses/>},
        {label: '收入', key: '2', children: 'Content of Tab Pane 2'},
        {label: '转账', key: '3', children: 'Content of Tab Pane 3'},
        {label: '借入', key: '4', children: 'Content of Tab Pane 4'},
        {label: '借出', key: '5', children: 'Content of Tab Pane 5'},
        {label: '收债', key: '6', children: 'Content of Tab Pane 6'},
        {label: '还债', key: '7', children: 'Content of Tab Pane 7'}
    ];
    return <Tabs size='small' items={items}/>;
};

