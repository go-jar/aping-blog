import Vue from 'vue'
import axios from 'axios'
import qs from 'qs';

// 管理
const AdminLogin =  (username, password, callback) =>{
    axios.post('/user/login', qs.stringify({
        "username": username, 
        "password": password
    }))
    .then(response => {
        callback && callback(response)
    })
    .catch(function (error) {
        console.log(error);
    });
}

export {
    AdminLogin, // 管理
}