<!-- 头部公用 -->
<template>
<div class="">
	<div class="headBack">
		<el-row class="container">
			<el-col :span="24">
				<!-- pc端导航 -->
				<div class="headBox">
					<el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect" :router="true">
						<el-menu-item index="/home" class="el-icon-house"> 首页</el-menu-item>
						<el-menu-item index="/list-categories" class="el-icon-notebook-2"> 类别</el-menu-item>
						<el-menu-item index="/list-tags" class="el-icon-collection-tag"> 标签</el-menu-item>
						<el-menu-item index="/about" class="el-icon-user"> 关于</el-menu-item>
						<el-submenu v-show="adminLogin" index="4">
							<template slot="title" class="el-icon-s-tools"> 管理</template>
							<el-menu-item index="/manage-cagegory">类别</el-menu-item>
							<el-menu-item index="/manage-tag">标签</el-menu-item>
							<el-menu-item index="/list-remarks">评论</el-menu-item>
						</el-submenu>
						<div class="createArticle">
							<el-menu-item v-show="adminLogin" index="/create-article" class="el-icon-edit"><i class="fa fa-fw fa-user-circle userImg"></i> 创作</el-menu-item>
						</div>
						<div index="" class="pcsearchbox">
							<i class="el-icon-search pcsearchicon"></i>
							<div class="pcsearchinput" :class="input?'hasSearched':''">
								<el-input placeholder="搜索" icon="search" v-model="input" :on-icon-click="searchEnterFun" @keyup.enter.native="searchEnterFun" @change="searchChangeFun">
								</el-input>
							</div>
						</div>
					</el-menu>
				</div>
				<!-- 移动端导航 -->
				<div class="mobileBox">
					<div class="hideMenu">
						<i @click="pMenu=!pMenu" class="el-icon-menu"></i>
						<el-collapse-transition>
							<el-menu :default-active="activeIndex" class="mlistmenu" v-show="!pMenu" theme="dark" @open="handleOpen" @close="handleClose" :unique-opened="true" :router="true">
								<el-menu-item index="/home"> 首页</el-menu-item>
								<el-menu-item index="/list-categories"> 类别</el-menu-item>
								<el-menu-item index="/list-tags"> 标签</el-menu-item>
								<el-menu-item index="/about"> 关于</el-menu-item>
							</el-menu>
						</el-collapse-transition>
						<div class="searchBox">
							<el-input placeholder="" icon="search" v-model="input" @keyup.enter.native="searchEnterFun" :on-icon-click="searchEnterFun" @change="searchChangeFun">
							</el-input>
						</div>
					</div>
				</div>
			</el-col>
		</el-row>
	</div>
</div>
</template>

<script>
import {Typeit} from '@/utils/plug'
import {Const,} from '@/const/common'
import {getToken} from '@/utils/auth.js'
export default {
	data() { //选项 / 数据
		return {
			adminLogin: false, // 是否已登录
			articleClassListObj: '', // 文章分类
			activeIndex: '/', // 当前选择的路由模块
			state: '', // icon点击状态
			pMenu: true, // 手机端菜单打开
			input: '', // input输入内容
		}
	},
	watch: {

	},
	methods: { // 事件处理器
		handleOpen(key, keyPath) { // 分组菜单打开
			// console.log(key, keyPath);
		},
		handleClose(key, keyPath) { // 分组菜单关闭
			// console.log(key, keyPath);
		},
		searchChangeFun(e) { // input change 事件
			// console.log(e)
			if (this.input == '') {
				this.$store.state.keywords = '';
				this.$router.push({path:'/'});
			}
		},
		searchEnterFun: function(e) { // input 输入 enter
			 var keyCode = window.event? e.keyCode:e.which;
             if(this.input){
				 this.$store.state.keywords = this.input;
                 this.$router.push({path:'/articles?keyword='+this.input});
             }
		},
		handleSelect(key, keyPath) { // pc 菜单选择
			// console.log(key, keyPath);
		},
		routeChange: function() {
			var that = this;
			that.pMenu = true
			this.activeIndex = this.$route.path == '/' ? '/Home' : this.$route.path;
			if (getToken()) { // 存储用户信息
				that.adminLogin = true;
			} else {
				that.adminLogin = false;
			}
		}
	},
	components: { // 定义组件

	},
	watch: {
		// 如果路由有变化，会再次执行该方法
		'$route': 'routeChange'
	},
	created() { // 生命周期函数
		// 判断当前页面是否被隐藏
		var that = this;
		document.title = Const.TITLE
		var hiddenProperty = 'hidden' in document ? 'hidden' :
			'webkitHidden' in document ? 'webkitHidden' :
			'mozHidden' in document ? 'mozHidden' :
			null;
		var visibilityChangeEvent = hiddenProperty.replace(/hidden/i, 'visibilitychange');
		var onVisibilityChange = function() {
			if (!document[hiddenProperty]) { // 被打开
				if (getToken()) {
					that.adminLogin = true;
				} else {
					that.adminLogin = false;
				}
			}
		}
		document.addEventListener(visibilityChangeEvent, onVisibilityChange);
		this.routeChange();
	},
	mounted() { // 页面元素加载完成
		var that = this;
		var timer = setTimeout(function() {
			Typeit("#luke"); // 打字机效果
			clearTimeout(timer);
		}, 500);
	}
}
</script>

<style>
/*********头部导航栏********/

/*头部导航栏盒子*/

.headBack {
	width: 100%;
	background: rgba(40, 42, 44, 0.6);
	/*margin-bottom:30px;*/
	box-shadow: 0 2px 4px 0 rgba(0, 0, 0, .12), 0 0 6px 0 rgba(0, 0, 0, .04);
	position: fixed;
	left: 0;
	top: 0;
	right: 0;
	z-index: 100;
}

.headBox li.is-active {
	/*background: #48456C;*/
	background: rgba(73, 69, 107, 0.7);
}

.el-menu--horizontal>.el-submenu.is-active .el-submenu__title {
	border-bottom: none!important;
}

.headBox .el-menu {
	background: transparent;
	border-bottom: none!important;
}

.headBox .el-menu-demo li.el-menu-item,
.headBox .el-menu--horizontal .el-submenu .el-submenu__title {
	height: 38px;
	line-height: 38px;
	border-bottom: none!important;

}

.headBox .el-submenu li.el-menu-item {
	height: 38px;
	line-height: 38px;
}

.headBox li .fa-wa {
	vertical-align: baseline;
}

.headBox ul li.el-menu-item,
.headBox ul li.el-menu-item.is-active,
.headBox ul li.el-menu-item:hover,
.headBox .el-submenu div.el-submenu__title,
.headBox .el-submenu__title i.el-submenu__icon-arrow {
	color: #fff;
}

.headBox .el-menu--horizontal .el-submenu>.el-menu {
	top: 38px;
	border: none;
	padding: 0;
}

.headBox>ul li.el-menu-item:hover,
.headBox>ul li.el-submenu:hover .el-submenu__title {
	background: #48456C;
	border-bottom: none;
}

.headBox>ul .el-submenu .el-menu,
.headBox>ul .el-submenu .el-menu .el-menu-item {
	background: #48456C;
}

.headBox>ul .el-submenu .el-menu .el-menu-item {
	min-width: 0;
}

.headBox>ul .el-submenu .el-menu .el-menu-item:hover {
	background: #64609E;
}

/*pc搜索框*/

.headBox .pcsearchbox {
	padding: 11px 0px;
	max-width: 170px;
	/*min-width: 30px;*/
	height: 100%;
	line-height: 38px;
	position: absolute;
	top: 0;
	right: 0;
	cursor: pointer;
}

.headBox .pcsearchbox:hover .pcsearchinput {
	opacity: 1;
	/*transform: scaleX(1);*/
	visibility: visible;
}

.headBox .pcsearchbox i.pcsearchicon {
	color: #fff;
	padding-left: 10px;
}

.headBox .pcsearchbox .pcsearchinput {
	width: 180px;
	padding: 10px 20px 10px 20px;
	background: rgba(40, 42, 44, 0.6);
	border-radius: 0 0 2px 2px;
	position: absolute;
	right: 0;
	top: 38px;
	opacity: 0;
	visibility: hidden;
	/*transform: scaleX(0);*/
	transform-origin: right;
	transition: all 0.3s ease-out;
}

.headBox .pcsearchbox .hasSearched {
	opacity: 1;
	/*transform: scaleX(1);*/
	visibility: visible;
}

.headBox .pcsearchbox .el-input {
	width: 100%;
}

.headBox .el-input__inner {
	height: 30px;
	border: none;
	background: #fff;
	/*border: 1px solid #333;*/
	border-radius: 2px;
	padding-right: 10px;
}

.headBox .createArticle {
	height: 100%;
	line-height: 38px;
	position: absolute;
	right: 30px;
	top: 0;
	color: #fff;
}

.headBox .createArticle a {
	color: #fff;
	font-size: 13px;
	transition: all 0.2s ease-out;
}

.headBox .createArticle a:hover {
	color: #48456C;
}

.headBox .nologin {
	text-align: right;
}

.headBox .haslogin {
	text-align: right;
	position: relative;
	min-width: 80px;
	cursor: pointer;
}

.headBox .haslogin:hover ul {
	visibility: visible;
	opacity: 1;
}

.headBox .haslogin ul {
	background: rgba(40, 42, 44, 0.6);
	padding: 5px 10px;
	position: absolute;
	right: 0;
	visibility: hidden;
	opacity: 0;
	transition: all 0.3s ease-out;
}

.headBox .haslogin ul li {
	border-bottom: 1px solid #48456C;
}

.headBox .haslogin ul li:last-child {
	border-bottom: 1px solid transparent;
}

/*******移动端*******/

.mobileBox {
	position: relative;
	height: 38px;
	line-height: 38px;
	color: #fff;
}

.hideMenu {
	position: relative;
	width: 100%;
	height: 100%;
	line-height: 38px;
}

.hideMenu ul.mlistmenu {
	width: 100%;
	position: absolute;
	left: 0;
	top: 100%;
	box-sizing: border-box;
	z-index: 999;
	box-shadow: 0 2px 6px 0 rgba(0, 0, 0, .12), 0 0 8px 0 rgba(0, 0, 0, .04);
	background: #48456C;
	color: #fff;
	border-right: none;
}

.hideMenu .el-submenu .el-menu {
	background: #64609E;
}

.hideMenu .el-menu-item,
.hideMenu .el-submenu__title {
	color: #fff;
}

.hideMenu>i {
	position: absolute;
	left: 10px;
	top: 12px;
	width: 30px;
	height: 30px;
	font-size: 16px;
	color: #fff;
	cursor: pointer;
}

.hideMenu .el-menu-item,
.el-submenu__title {
	height: 40px;
	line-height: 40px;
}

.mobileBox .searchBox {
	padding-left: 40px;
	width: 100%;
	box-sizing: border-box;
}

.mobileBox .searchBox .el-input__inner {
	display: block;
	border-radius: 2px;
	border: none;
	height: 25px;
}

.hideMenu ul.mlistmenu.pshow {
	display: block;
}

.hideMenu ul.mlistmenu .el-submenu__icon-arrow,
.mobileBox li.el-menu-item a {
	color: #fff;
}

.hideMenu>ul li.el-menu-item:hover,
.hideMenu>ul li.el-menu-item.is-active {
	background: #48576a;
}
.el-input__inner {
  background-color: rgba(230, 244, 249, 0.85);
  border-radius: 1px;
}
</style>