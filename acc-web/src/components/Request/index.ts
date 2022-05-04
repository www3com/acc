import axios from "axios";

export const get = async (url: string, params?: object, options?: object) => {
  try {
    let res = await axios({
      method: 'get',
      url: url,
      params: params,
    })
    return res.data
  } catch (e) {

  }

}

export default {
  get: get
}
