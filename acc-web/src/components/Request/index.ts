import axios from "axios";
import {message, notification} from "antd";
import {getToken} from "@/components/token";

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
        'Content-Type': 'application/json;charset=UTF-8',
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
        'Content-Type': 'application/json;charset=UTF-8',
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
  debugger
  switch (err.response.status) {
    case 400:
    case 401:
    case 403:
    case 404:
    case 500:
      message.error(err.response.data.message)
      return
    default:
      message.info("请求数据发生错误！")
  }
}

export default {
  get: get,
  post: post,
  put: put,
  delete: del
}
