<!-- 文章详情模块 -->
<template>
    <div class="detailBox tcommonBox">
        <header>
            <h1 style="padding-top:10px;">
                {{articleData.Article.Title}}
            </h1>
            <h2>
                <el-row>
                    <span class="el-icon-date"> {{formatTime}}</span>&#12288; 

                    <a :href="'#/category?id='+articleData.Article.CategoryId" class="el-icon-notebook-2"> {{articleData.Category.CategoryName}}</a>&#12288;

                    <i v-for="item in articleData.TagSet" :key="item.TagId">
                        <a :href="'#/tag?id='+item.TagId" class="el-icon-collection-tag"> {{item.TagName}}</a>&#12288;
                    </i>

                    <i class="el-icon-view"></i> {{articleData.Article.ReadCount}}&#12288;
                    <span v-show="articleData.RemarkCount>0" class="el-icon-chat-dot-square"> {{articleData.RemarkCount}}&#12288;</span>

                    <a :href="'#/remark?articleid='+articleData.Article.ArticleId" target="_blank" class="el-icon-edit-outline"> 评论&#12288;</a>
                    <a :href="'#/edit-article?id='+articleData.Article.ArticleId" target="_blank" v-show="adminLogin" class="el-icon-edit-outline"> 编辑&#12288;</a>
                    <!-- <a :href="'#'" @click="handleDeleteArticle(articleData.Article.ArticleId)" v-show="adminLogin" class="el-icon-remove-outline"> 删除&#12288;</a> -->
                    <el-link v-show="adminLogin" type="danger" size="mini" icon="el-icon-delete" @click="handleDeleteArticle(articleData.Article.ArticleId)"
                    />&#12288;
                    <span style="color: red; margin-right: 0px;" class="el-icon-warning-outline"> 转载请注明出处</span>
                </el-row>
            </h2>
        </header>
        <ViewMarkdown :height="780" :content="articleData.Article.Content"></ViewMarkdown>
        <!-- <div class="article-content" v-html="articleData.Article.Content"></div> -->
        <!-- <div class="donate">
            <div class="donate-word">
                <span @click="pdonate=!pdonate">赞赏</span>
            </div>
            <el-row :class="pdonate?'donate-body':'donate-body donate-body-show'" :gutter="30">
                <el-col :span="12" class="donate-item">
                    <div class="donate-tip">
                        <img :src="articleData.wechat_image?articleData.wechat_image: 'static/img/tou.jpg'" :onerror="$store.state.errorImg"/>
                        <span>微信扫一扫，向我赞赏</span>
                    </div>
                </el-col>
                <el-col :span="12" class="donate-item">
                    <div class="donate-tip">
                        <img :src="articleData.alipay_image?articleData.alipay_image:'static/img/tou.jpg'" :onerror="$store.state.errorImg"/>
                        <span>支付宝扫一扫，向我赞赏</span>
                    </div>
                </el-col>
            </el-row>
        </div> -->
    </div>
</template>

<script>
import {describeArticles, deleteArticle} from '@/api/article.js'
import {getToken} from '@/utils/auth.js'
import ViewMarkdown from '@/components/ViewMarkdown'

export default {
    name: 'ArticleDetail',
    data() { // 选项 / 数据
        return {
            articleId: '', // 文章 ID
            pdonate: true, // 打开赞赏控制,
            articleData: '', // 返回详情数据
            adminLogin: false,
        }
    },
    components: { // 定义组件
        ViewMarkdown,
    },
    watch: {
        // 如果路由有变化，会再次执行该方法
        '$route':'routeChange'
    },
    created() { //生命周期函数
        var that = this;
        this.routeChange();

        Date.prototype.format = function(fmt) { 
            var o = { 
                "M+" : this.getMonth()+1,                 //月份 
                "d+" : this.getDate(),                    //日 
                "h+" : this.getHours(),                   //小时 
                "m+" : this.getMinutes(),                 //分 
                "s+" : this.getSeconds(),                 //秒 
                "q+" : Math.floor((this.getMonth()+3)/3), //季度 
                "S"  : this.getMilliseconds()             //毫秒 
            }; 
            if(/(y+)/.test(fmt)) {
                    fmt=fmt.replace(RegExp.$1, (this.getFullYear()+"").substr(4 - RegExp.$1.length)); 
            }
            for(var k in o) {
                if(new RegExp("("+ k +")").test(fmt)){
                    fmt = fmt.replace(RegExp.$1, (RegExp.$1.length==1) ? (o[k]) : (("00"+ o[k]).substr((""+ o[k]).length)));
                }
            }
            return fmt; 
        }
    },
    // 计算属性
    computed: {
        // 对时间进行格式化
        formatTime: function() {
            if (this.articleData) {
                const dt = new Date(this.articleData.Article.CreatedTime)
                return dt.format("yyyy-MM-dd");
            }
            return '';
        },
    },
    methods: { // 事件处理器
        routeChange: function(){
            var that = this;

            if (getToken()) {
				that.adminLogin = true;
			} else {
				that.adminLogin = false;
			}

            that.articleId = that.$route.query.id==undefined?1:parseInt(that.$route.query.id);  // 获取传参的 id
            
            //获取详情接口
            describeArticles(that.articleId, null, null, null, null).then(response => {
                that.articleData = response.Data.ArticleSet[0];
            })
        },
        handleDeleteArticle: function(articleId) {
            var that = this;

            this.$confirm("此操作将把本文删除, 是否继续?", "提示", {
                confirmButtonText: "确定",
                cancelButtonText: "取消",
                type: "warning"
            }).then(() => {
                deleteArticle(articleId).then(response => {
                    if (response.Code == Code.SUCCESS) {
                        message.success(response.message)
                    } else {
                        message.error(response.message)
                    }
                });

                that.$router.push({
                    path: '/'
                });
            }).catch(() => {
                message.info("已取消删除")
            });
        },
    }, 
}
</script>

<style lang="less">
.detailBox .article-content{
    font-size: 15px;
    white-space: normal;
    word-wrap: break-word;
    word-break: break-all;
    overflow-x: hidden;
    text-align: center;
    height: 100%;
}
.detailBox .article-content p{
    margin:10px 0;
    line-height:24px;
    word-wrap: break-word;
    word-break: break-all;
    overflow-x: hidden;
}
.detailBox .article-content pre{
    word-wrap: break-word;
    word-break: break-all;
    overflow-x: hidden;
}
.detailBox .article-content img{
    max-width: 100%!important;
    height: auto!important;
    overflow-x: hidden;
}
.detailBox .article-content a{
    color:#df2050!important;
}
.detailBox .article-content a:hover{
    text-decoration: underline;
    color: #f00!important;
}
.detailBox .article-content i{
    font-style: italic;
}
.detailBox .article-content strong{
    font-weight: bold;
}
.detailBox .article-content ul{
    list-style-type: disc!important;
    list-style: disc!important;
    padding-left: 40px!important;
    li{
        list-style-type: disc!important;
        list-style: disc!important;
    }
}
.detailBox .article-content h1, .detailBox .article-content h2, .detailBox .article-content h3{
    font-size: 200%;
    font-weight: bold;
}
 .detailBox .article-content h4, .detailBox .article-content h5, .detailBox .article-content h6{
    font-size: 150%;
    font-weight: bold;
}
.detailBox .viewdetail{
    margin:10px 0 ;
    line-height: 24px;
    text-align: center;
}
/*分享图标*/
.dshareBox {
    margin-top:40px;
    position: relative;
}
.dshareBox a{
    position: relative;
    display: inline-block;
    width: 32px;
    height: 32px;
    font-size: 18px;
    border-radius: 50%;
    line-height: 32px;
    text-align: center;
    vertical-align: middle;
    margin: 4px;
    background: #fff;
    transition: background 0.6s ease-out;
}
.dshareBox .ds-weibo{
    border: 1px solid #ff763b;
    color: #ff763b;
}
.dshareBox .ds-weibo:hover{
    color: #fff;
    background: #ff763b;
}
.dshareBox .ds-qq{
    color: #56b6e7;
    border: 1px solid #56b6e7;
}
.dshareBox .ds-qq:hover{
    color: #fff;
    background: #56b6e7;
}
.dshareBox .ds-wechat{
    color: #7bc549;
    border: 1px solid #7bc549;
}
.dshareBox .ds-wechat:hover{
    color: #fff;
    background: #7bc549;
}
.dshareBox .ds-wechat:hover .wechatShare{
    opacity: 1;
    visibility: visible;
}
.detailBox .bdshare-button-style0-32 a{
    float:none;
    background-image: none;
    text-indent: inherit;
}

/*赞赏*/
.donate-word{
    margin:40px 0;
    text-align: center;
}
.donate-word span{
    display: inline-block;
    width:80px;
    height:34px;
    line-height: 34px;
    color:#fff;
    background: #e26d6d;
    margin:0 auto;
    border-radius: 4px;
    cursor: pointer;
}
.donate-body{
    display: none;
}
.donate-body-show{
    display: block;
}
.donate-item{
    text-align: right;
}
.donate-item:last-child{
    text-align: left;
}
.donate-item img{
    width:100%;
    display: block;
    height:auto;
}
.donate-item div{
    display: inline-block;
    width: 150px;
    padding: 0 6px;
    box-sizing: border-box;
    text-align: center;
}
.donate-item div span{
    display: inline-block;
    width:100%;
    margin: 10px 0;
    text-align: center;
}
.donate-body .donate-item:first-of-type div{
    color:#44b549;
}
.donate-body .donate-item:nth-of-type(2) div{
    color:#00a0e9;
}

.bd_weixin_popup{
    position: fixed!important;
}
</style>
