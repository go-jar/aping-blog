import axios from 'axios'
import {Code} from '../const/code.js'
import {getToken, removeToken} from '../utils/auth.js'

// 创建 axios 实例
const service = axios.create({
  baseURL: '', // api 的 base_url
  timeout: 5000 // 请求超时时间
})

// request 拦截器
service.interceptors.request.use(
  config => {
    if (getToken() != '') {
        config.headers['X-Token'] = getToken() // 让每个请求携带自定义 token 请根据实际情况自行修改
    }
    return config
  },
  error => {
    // Do something with request error
    console.log(error) // for debug
    Promise.reject(error)
  }
)

// response 拦截器
service.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code != Code.LOGIN_EXPIRED) {
      return response.data
    } else {
      MessageBox.confirm(
        '登录已过期，可以取消继续留在该页面，或者重新登录',
        '确定登出',
        {
          confirmButtonText: '重新登录',
          cancelButtonText: '取消',
          type: 'warning'
        }
      ).then(() => {
        removeToken();
        location.reload() // 为了重新实例化 vue-router 对象 避免 bug
      })
      return Promise.reject('error')
    }
  },
  error => {
    // 出现网络超时
    router.push('500')
    return Promise.reject(error)
  }
)

export default service
