import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)

export default new Router({
	scrollBehavior(to, from, savePosition) { // 在点击浏览器的“前进/后退”，或者切换导航的时候触发。
		if (savePosition) {
			return savePosition;
		} else {
            var top;
            if (window.innerWidth >= 700) {
                 top = 676
            } else {
                 top = 267
            }
			return {
				x: 0,
				y: top
			}
		}
	},
	routes: [
		// 首页
		{
		path: '/',
		name: 'Home',
				component: resolve => require(['../pages/Home.vue'], resolve),
				meta: {
					auth: true
				}
		}, 
		// 首页
		{
			path: '/Home',
			name: 'Home',
				component: resolve => require(['../pages/Home.vue'], resolve),
				meta: {
					auth: true
				}
		},
		// 博客列表
		{
			path: '/BlogList',
			name: 'BlogList',
				component: resolve => require(['../pages/BlogList.vue'], resolve),
				meta: {
					auth: true
				}
		},
		// 博客详情页
		{
			path: '/BlogDetail',
			name: 'BlogDetail',
				component: resolve => require(['../pages/BlogDetail.vue'], resolve),
				meta: {
					auth: true
				}
		},
		// 注册登录
		{
			path: '/Login',
			name: 'Login',
			component: resolve => require(['../pages/Login.vue'], resolve),
			meta: {
				auth: false
			}
		},
		// 注册登录
		{
			path: '/Admin',
			component: resolve => require(['../pages/Admin.vue'], resolve),
			meta: {
				auth: false
			}
		},
		// 关于
		{
			path: '/About',
			name: 'About',
			component: resolve => require(['../pages/About.vue'], resolve),
			meta: {
				auth: true
			}
		}, 
	]
})



