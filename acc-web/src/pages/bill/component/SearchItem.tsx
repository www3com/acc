import {Button, Divider, Space} from 'antd';
import React, {useEffect, useRef, useState} from "react";
import {DownOutlined} from "@ant-design/icons";
import styles from "@/pages/bill/style.less";
import classnames from 'classnames';

interface ItemProps {
    title: string,
    children?: any,
    onOk?: Function,
    bodyStyle?: any
}

export default ({title, children, onOk, bodyStyle}: ItemProps) => {
    const ref = useRef<HTMLDivElement | null>(null);
    const [visible, setVisible] = useState(false)
    const handleClick = (e: any) => {
        if (ref.current && !ref.current.contains(e.target)) {
            setVisible(false)
        }
    };

    const okHandler = () => {
        setVisible(false)
        if (onOk) {
            onOk()
        }
    }

    useEffect(() => {
        document.addEventListener('click', handleClick);
        return () => {
            document.removeEventListener('click', handleClick);
        };
    }, [])

    return <div ref={ref} className={styles.searchItem}>
        <Space size={5} onClick={event => setVisible(true)}>{title}<DownOutlined/></Space>
        <div className={visible ? classnames(styles.active, styles.innerItem) : styles.innerItem} style={bodyStyle}>
            {children}
            <Divider style={{margin: '8px 0px 8px 0px'}}/>
            <Space style={{width: '100%', direction: 'rtl', paddingRight: 10}}>
                <Button type='primary' size='small' onClick={okHandler}>确定</Button>
            </Space>
        </div>
    </div>;
};

