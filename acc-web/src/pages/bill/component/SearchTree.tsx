import {Button, Divider, Space, Tree, TreeDataNode} from 'antd';
import React, {useEffect, useImperativeHandle, useRef, useState, forwardRef, ForwardRefRenderFunction} from "react";
import {DownOutlined} from "@ant-design/icons";
import styles from "@/pages/bill/style.less";
import classnames from 'classnames';
import {AntTreeNode} from "antd/es/tree";
import SearchDropdown from "@/pages/bill/component/SearchDropdown";


export interface SearchTreeHandler {
    reset: () => void;
}

interface SearchTreeProps {
    title: string,
    treeData: any,
    fieldNames?: any,
    onOk?: Function,
    bodyStyle?: any
}

const searchTree: ForwardRefRenderFunction<SearchTreeHandler, SearchTreeProps> = ({
                                                                                      title,
                                                                                      treeData,
                                                                                      fieldNames,
                                                                                      onOk,
                                                                                      bodyStyle
                                                                                  }: SearchTreeProps, ref: any) => {


    const [checkNodes, setCheckNodes] = useState<any[]>([])

    const okHandler = () => {

        if (onOk) {
            onOk(checkNodes)
        }
    }

    useImperativeHandle(ref, () => ({
        reset: () => {
            setCheckNodes([])
        }
    }));

    return <SearchDropdown title={title} bodyStyle={bodyStyle} onOk={okHandler}>
        <Tree checkable fieldNames={fieldNames ? fieldNames : {title: 'name', key: 'id'}}
              treeData={treeData}
              checkedKeys={checkNodes.map((value: any) => value.id)}
              onCheck={(selectedKey: any, e) => setCheckNodes(e.checkedNodes)}/>
    </SearchDropdown>;
};

export default forwardRef(searchTree);
