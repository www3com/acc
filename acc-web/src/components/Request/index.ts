import axios from "axios";
import {message} from "antd";
import {getToken} from "@/components/session";

const json = "application/json;charset=UTF-8"

const get = async (url: string, params?: object) => {
    try {
        let res = await axios({
            method: 'get',
            url: url,
            params: params,
            headers: {
                "ACC-TOKEN": getToken()
            },
        })
        return res.data
    } catch (err) {
        fail(err)
        throw err
    }
}

const post = async (url: string, data?: object) => {
    try {
        let res = await axios({
            method: 'post',
            url: url,
            data: data,
            headers: {
                'Content-Type': json,
                "ACC-TOKEN": getToken()
            },
        })
        return res.data
    } catch (err) {
        fail(err)
        throw err
    }
}

const put = async (url: string, data?: object) => {
    try {
        let res = await axios({
            method: 'put',
            url: url,
            data: data,
            headers: {
                'Content-Type': json,
                "ACC-TOKEN": getToken()
            },
        })
        return res.data
    } catch (err) {
        fail(err)
        throw err
    }
}

const del = async (url: string, params?: object) => {
    try {
        let res = await axios({
            method: 'delete',
            url: url,
            params: params,
            headers: {
                "ACC-TOKEN": getToken()
            },
        })
        return res.data
    } catch (err) {
        fail(err)
        throw err
    }
}

const fail = (err: any) => {
    switch (err.response.status) {
        case UNAUTHORIZED:
            message.error(err.response.data.message)
            location.href = "/sign-in"
            break
        case INVALID_PARAMS:
            message.error(err.response.data.message)
            break
        case FORBIDDEN:
            message.error(err.response.data.message)
            break
        case NOT_FOUND:
            message.error(err.response.data.message)
            break
        case ERROR:
            message.error(err.response.data.message)
            return
        default:
            message.info("请求数据发生错误！")
    }
}

export const OK = 200
export const ERROR = 500
export const INVALID_PARAMS = 400
export const UNAUTHORIZED = 401
export const FORBIDDEN = 403
export const NOT_FOUND = 404

export default {
    get: get,
    post: post,
    put: put,
    delete: del
}
