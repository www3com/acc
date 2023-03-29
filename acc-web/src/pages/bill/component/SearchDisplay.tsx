import {Col, Row} from 'antd';


import React from "react";
import {CloseOutlined} from "@ant-design/icons";
import styles from "@/pages/bill/style.less";

interface ItemProps {
    title: string,
    titleWith?: string,
    data?: string[],
    onClose?: Function
}

export default ({title, titleWith, data, onClose}: ItemProps) => {
    if (!data || data.length == 0) {
        return <></>
    }
    return <Row wrap={false} align='middle' className={styles.searchDisplay}>
        <Col flex={titleWith ? titleWith : '100px'} className={styles.title}>
            {title}
        </Col>
        <Col flex="auto">
            {data.map((item) => {
                return <span style={{marginRight: 20, color:'#000000A6'}}>{item}</span>
            })}
        </Col>
        <Col flex="50px" className={styles.operator}>
            <a onClick={() => {
                if (onClose) {
                    onClose()
                }
            }}><CloseOutlined/></a>
        </Col>
    </Row>;
};

