import qs from 'qs';
import request from '../utils/request'

// 管理员登录
export function AdminLogin(username, password) {
    return request({
        url: process.env.WEB_API + '/user/login',
        method: 'post',
        data: qs.stringify({
            "username": username, 
            "password": password
        })
    })
}
