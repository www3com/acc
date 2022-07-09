import {Button, Card, Col, Divider, List, Row, Space} from "antd";
import {AppstoreOutlined, BookOutlined, CheckOutlined} from "@ant-design/icons";
import {inject, observer} from "mobx-react";
import style from "../style.less";
import React, {useState} from 'react'

const ledgerDropdown = ({store}: any) => {
    const [current, setCurrent] = useState(store.currentLedger.id)

    const onClick = (item: any) => {
        setCurrent(item.id)
        store.setCurrentLedger(item)
    }

    const renderItem = (item: any) => {
        return (
            <List.Item className={style.item} onClick={() => onClick(item)}>
                <Row style={{width: '100%'}}>
                    <Col flex="auto">{item.name}</Col>
                    <Col flex="10px">{item.id == current ? <CheckOutlined/> : ''}</Col>
                </Row>
            </List.Item>
        )
    }

    return (
        <Card className={style.ledger} bodyStyle={{padding: 0}} size={"small"}>
            <List split={false} size={'small'} dataSource={store.ledgers} renderItem={item => renderItem(item)}/>
            <Divider style={{margin: '5px 0px'}}/>
            <Space>
                <Button type="text" icon={<AppstoreOutlined/>}>账本模版</Button>
                <Button type="text" icon={<BookOutlined/>}>账本管理</Button>
            </Space>
        </Card>
    );
};

export default inject('store')(observer(ledgerDropdown))
