import qs from 'qs';
import request from '../utils/request'

// 管理员登录
export function adminLogin(username, password) {
    return request({
        url: process.env.WEB_API + '/user/login',
        method: 'post',
        data: qs.stringify({
            "Action": "AdminLogin",
            "Username": username, 
            "Password": password
        })
    })
}
