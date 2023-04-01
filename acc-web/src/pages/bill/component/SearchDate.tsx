import {Calendar, Space} from 'antd';
import React, {forwardRef, ForwardRefRenderFunction, useImperativeHandle, useState} from "react";
import {ArrowRightOutlined} from "@ant-design/icons";

import SearchDropdown from "@/pages/bill/component/SearchDropdown";
import dayjs, {Dayjs} from "dayjs";

export interface SearchDateHandler {
    reset: () => void;
}

interface SearchDateProps {
    title: string,
    onOk?: Function,
    bodyStyle?: any
}

const searchDate: ForwardRefRenderFunction<SearchDateHandler, SearchDateProps> = ({
                                                                                      title,
                                                                                      onOk,
                                                                                      bodyStyle
                                                                                  }: SearchDateProps, ref: any) => {
    const [startTime, setStartTime] = useState<Dayjs>(dayjs());
    const [endTime, setEndTime] = useState<Dayjs>(dayjs())

    const okHandler = () => {
        if (onOk) {
            onOk(startTime, endTime);
        }
    }

    useImperativeHandle(ref, () => ({
        reset: () => {
            setStartTime(dayjs())
            setEndTime(dayjs())
        }
    }));
    return <SearchDropdown title={title} bodyStyle={bodyStyle} onOk={okHandler}>
        <Space>
            <Calendar fullscreen={false} value={startTime}
                      onSelect={(date: Dayjs) => setStartTime(date)}/>
            <ArrowRightOutlined/>
            <Calendar fullscreen={false} value={endTime}
                      onSelect={(date: Dayjs) => setEndTime(date)}/>
        </Space>
    </SearchDropdown>;
};
export default forwardRef(searchDate);
