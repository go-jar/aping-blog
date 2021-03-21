<!-- 登录注册 -->
<template>
<div>
    <div class="container">
        <div class="loginBox">
            <div class="lr-title">
                <h1>登录</h1>
            </div>
            <el-alert v-show="loginErr" :title="loginTitle" type="error" show-icon :closable="false">
            </el-alert>
            <el-input placeholder="用户名" v-model="username">
            </el-input>
            <el-alert v-show="usernameErr" title="用户名" type="error" show-icon :closable="false">
            </el-alert>
            <el-input type="password" placeholder="密码" @keyup.enter.native="loginEnterFun" v-model="password">
            </el-input>
            <el-alert v-show="passwordErr" title="请输入密码" type="error" show-icon :closable="false">
            </el-alert>
            <div class="lr-btn tcolors-bg" @click="login">登录</div>
        </div>
    </div>
</div>
</template>

<script>
import {adminLogin} from '@/api/user.js'
import {Code} from '@/const/code.js'
import {LoginKey} from '@/const/login.js'
import {getToken, setToken} from '@/utils/auth.js'

export default {
    data() { // 选项 / 数据
        return {
            username: '', // 用户名
            password: '', // 密码
            isLogin: 0,   // 是否已经登录
            usernameErr: false, // 用户名错误
            passwordErr: false, // 密码错误
            loginErr: false, // 登录错误
            loginTitle: '用户名或密码错误',
        }
    },
    methods: { // 事件处理器
        loginEnterFun: function (e) {
            var keyCode = window.event ? e.keyCode : e.which;
            if (keyCode == 13) {
                this.login();
            }
        },
        login: function () { // 用户登录
            var that = this;
            var reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/;
            var preg = /^([a-zA-Z0-9?]){6,12}/;
            
            that.usernameErr = that.username? false: true;
            that.passwordErr = that.password && preg.test(that.password)? false: true;

            if (!that.usernameErr && !that.passwordErr) {
                adminLogin(that.username, that.password).then(response => {
                    if (response.Code == Code.SUCCESS) { // 登录成功
                        setToken(response.Data.Token);
                        that.$router.push({
                            path: '/'
                        });
                    } else {
                        that.loginErr = true;
                        that.loginTitle = '登录失败';
                    }
                })
            }
        },
    },
    components: { // 定义组件
    },
    watch: {
    },
    created() { // 生命周期函数
    }
}
</script>

<style>
.loginTitle {
    text-align: center;
    font-size: 26px;
    padding-top: 50px;
    margin-bottom: 20px;
}

.loginBox {
    background: #fff;
    padding: 35px 40px;
    max-width: 320px;
    margin: 22.7% auto 36%;
}

.lr-title {
    position: relative;
    height: 32px;
    line-height: 32px;
    margin-bottom: 20px;
}

.lr-title h1 {
    font-size: 24px;
    color: #666;
    font-weight: bold;
    /*width:50%;*/
}

.lr-title p {
    font-size: 12px;
    color: #999;
    position: absolute;
    right: 0;
    top: 0;
}

.lr-btn {
    color: #fff;
    text-align: center;
    letter-spacing: 5px;
    padding: 8px;
    border-radius: 5px;
    cursor: pointer;
    margin-bottom: 30px;
}

.loginBox .el-input,
.registerBox .el-input {
    margin-bottom: 20px;
}

.loginBox .el-alert,
.registerBox .el-alert {
    top: -18px;
    background-color: #888;
}

.loginBox .el-input input,
.registerBox .el-input input {
    border-radius: 4px;
}

.loginBox h3,
.registerBox h3 {
    text-align: right;
    margin-bottom: 20px;
}

.loginBox h3 a,
.registerBox h3 a {
    font-size: 13px;
    color: #999;
}

.loginBox .otherLogin {
    max-width: 320px;
    padding: 30px 40px;
    background: #ddd;
    text-align: center;
    margin-left: -40px;
    margin-right: -40px;
    visibility: hidden;
}

.loginBox .otherLogin p {
    margin-bottom: 20px;
    font-size: 16px;
}

.loginBox .otherLogin a i {
    display: inline-block;
    width: 42px;
    height: 42px;
    line-height: 42px;
    font-size: 18px;
    border-radius: 50%;
    color: #fff;
    margin: 0 10px;
}

.loginBox .otherLogin a i.fa-wechat {
    background: #7bc549;
}

.loginBox .otherLogin a i.fa-qq {
    background: #56b6e7;
}

.loginBox .otherLogin a i.fa-weibo {
    background: #ff763b;
}
</style>
