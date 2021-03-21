import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)

export const constantRouterMap = [
	{ path: '/', component: () => import('@/pages/Home') },
	{ path: '/home', component: () => import('@/pages/Home') },
	{ path: '/admin-login', component: () => import('@/pages/AdminLogin') },
	{ path: '/about', component: () => import('@/pages/About') },
	{ path: '/edit-article', component: () => import('@/pages/EditArticle') },
	{ path: '/article', component: () => import('@/pages/ArticleDetail') },
	{ path: '/manage-cagegory', component: () => import('@/pages/ManageCategory') },
	{ path: '/manage-tag', component: () => import('@/pages/ManageTag') },
	{ path: '/list-categories', component: () => import('@/pages/ListCategories') },
	{ path: '/list-tags', component: () => import('@/pages/ListTags') },
	{ path: '/category', component: () => import('@/pages/ListArticlesByCategory') },
	{ path: '/tag', component: () => import('@/pages/ListArticlesByTag') },
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
