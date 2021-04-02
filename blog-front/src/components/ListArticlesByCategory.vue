<!-- 首页 -->
<template>
<div class="container">
<div class="content">
<el-row>
        <el-col
            v-for="item in this.articleObjs"
            :key="item.ArticleId"
            style="padding: 3px"
        >

            <el-card
                :body-style="{height: '50px', padding: '0px', textAlign: 'center'}"
                style="position: relative"
                shadow="always"
            >
                <div class="title">
                    <div class="in-title">
                        <el-row style="textAlign: left;">
                            <el-col :span=1>&#12288;</el-col>
                            <el-col :span=10>
                                <a :href="'#/article?id='+item.Article.ArticleId">
                                    {{item.Article.Title}}
                                </a>
                            </el-col>
                            <el-col :span=13>
                                <span class="el-icon-date"> {{formatTime(item.Article.CreatedTime)}}</span>&#12288;
                                <a :href="'#/category?id='+item.Article.CategoryId" class="el-icon-notebook-2">
                                    {{item.Category.CategoryName}}
                                </a>&#12288;
                                <i v-for="tag in item.TagSet" :key="tag.TagId">
                                    <a :href="'#/tag?id='+tag.TagId" class="el-icon-collection-tag"> {{tag.TagName}}</a>&#12288;
                                </i>
                                <span class="el-icon-view"> {{item.Article.ReadCount}}</span>&#12288;
                                <span v-show="item.RemarkCount>0" class="el-icon-chat-dot-square"> {{item.RemarkCount}}</span>
                            </el-col>
                        </el-row>
                    </div>
                </div> 
            </el-card>
        </el-col>
    </el-row>

    <!-- 分页 -->
    <div class="block" style="margin: 3px 0 10px 6px;">
        <el-pagination
            @current-change="handleCurrentChange"
            :current-page.sync="currentPage"
            :pager-count="pagerCount"
            :page-size="pageSize"
            layout="total, prev, pager, next, jumper"
            :total="total">
        </el-pagination>
    </div>
</div>
</div>
</template>

<script>
import {describeArticles} from "@/api/article";
import {Code} from '@/const/code.js'
import {message} from '@/utils/common'

export default {
    data() {  // 选项 / 数据
        return {
            title: null,
            articleObjs: [],
            formLabelWidth: "120px",
            currentPage: 1,
            total: null,
            pagerCount: window.innerWidth <= 700? 2: 10,
            pageSize: window.innerWidth <= 700? 8: 13,
            tag: {
                TagId: null,
                TagName: null,
            },
            categoryId: null,
        }
    },
    components: {  // 定义组件
    },
    watch: {
        // 如果路由有变化，会再次执行该方法
        '$route':'handleListArticles'
    },
    created() {  // 生命周期函数
        this.handleListArticles();

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
    },
    methods: {  // 事件处理器
        // 对时间进行格式化
        formatTime: function(createdDate) {
            if (createdDate) {
                const dt = new Date(createdDate)
                return dt.format("yyyy-MM-dd");
            }
            return '';
        },
        // 改变页码
        handleCurrentChange(val) {
            var that = this;
            this.currentPage = val; // 改变当前所指向的页数
            this.handleListArticles();
        },
        handleListArticles: function(){
            var that = this;
            that.categoryId = that.$route.query.id==undefined?1:parseInt(that.$route.query.id);  // 获取传参的 id
            var offset = (this.currentPage - 1) * this.pageSize;

            //获取详情接口
            describeArticles(null, that.categoryId, null, null, offset, this.pageSize).then(response => {
                if(response.Code == Code.SUCCESS) {
                    this.articleObjs = response.Data.ArticleSet;
                    this.total = response.Data.Total;
                }
            })
        },
    }
}
</script>

<style scoped>
.content {
  position: relative;
  border-radius: 5px;
  height: 93.1%;
  margin-top: 38px;
  padding-top: 10px;
  background: rgba(230, 244, 249, 0.85);
  opacity: 0.98;
}
a {
  color: #fff; 
}
.title {
    height: 100%;
    z-index: 2;
    background: rgba(58, 59, 59);
    color: #FFF;
    font-size: 15px;
    border-radius: 3px;
    width: 100%;
    text-align: center;
}
.in-title {
    padding-top: 12px;
}
@media screen and (max-width: 700px) {
    .title {
        font-size: 12px;
    }
}
</style>
