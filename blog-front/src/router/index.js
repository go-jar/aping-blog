import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)

export const constantRouterMap = [
	{ path: '/', component: () => import('@/pages/Home') },
	{ path: '/Home', component: () => import('@/pages/Home') },
	{ path: '/AdminLogin', component: () => import('@/pages/AdminLogin') },
	{ path: '/About', component: () => import('@/pages/About') },
  ]
  
const router = new Router({
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
	routes: constantRouterMap
})

export default router
