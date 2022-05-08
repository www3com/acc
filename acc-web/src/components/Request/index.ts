import axios from "axios";
import {message, notification} from "antd";

const get = async (url: string, params?: object, options?: object) => {
  try {
    let res = await axios({
      method: 'get',
      url: url,
      params: params,
    })
    return res.data
  } catch (err) {
    fail(err)
  }

}

const post = async (url: string, data?: object) => {
  try {
    let res = await axios({
      method: 'post',
      url: url,
      data: data,
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
    })
    debugger
    return res.data
  } catch (err) {
    fail(err)
    throw err
  }

}

const fail = (err: any) => {
  switch (err.response.status) {
    case 400:
      message.error("参数不合法!")
      return
    case 401:
      message.error("请求要求身份验证!")
      return
    case 403:
      message.error("服务器拒绝请求!")
      return
    case 404:
      message.error("请求资源不存在!")
      return
    case 500:
      message.error("服务器内部错误!")
      return
    default:
      message.info("请求数据发生错误！")
  }

}

export default {
  get: get,
  post: post
}
