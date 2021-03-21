<!-- 首页 -->
<template>
<div>
<el-row>
        <el-col
            v-for="item in this.articleObjs"
            :key="item.Id"
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
                            <el-col :span=4>&#12288;</el-col>
                            <el-col :span=9>
                                <a :href="'#/article?id='+item.Article.Id">
                                    {{item.Article.Title}}
                                </a>
                            </el-col>
                            <el-col :span=11>
                                <span >{{formatTime(item.Article.CreatedTime)}}</span>&#12288;
                                <a :href="'#/category?id='+item.Article.CategoryId">
                                    {{item.Category.CategoryName}}
                                </a>&#12288;
                                <i v-for="tag in item.TagSet" :key="tag.Id">
                                    <a :href="'#/tag?id='+tag.Id">{{tag.TagName}}</a>&#12288;
                                </i>
                                <span >{{item.Article.ReadCount}}</span>
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
            :page-size="pageSize"
            layout="total, prev, pager, next, jumper"
            :total="total">
        </el-pagination>
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
            isChange: false,  // 内容是否改变
            articleObjs: [],
            formLabelWidth: "120px",
            currentPage: 1,
            total: null,
            pageSize: 40,
            tag: {
                Id: null,
                TagName: null,
            },
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
        // 关闭窗口
        closeDialog(done) {
            if(this.isChange) {
                this.$confirm("是否关闭类别编辑窗口", "提示", {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning"
                })
                .then(() => {
                    this.isChange = false;
                    this.changeCount = 0
                    done();
                })
                .catch(() => {
                    message.info("已取消")
                });
            } else {
                this.isChange = false;
                this.changeCount = 0
                done();
            }
        },
        // 改变页码
        handleCurrentChange(val) {
            var that = this;
            this.currentPage = val; // 改变当前所指向的页数
            this.handleListTags();
        },
        handleListArticles: function() {
            var offset = (this.currentPage - 1) * this.pageSize;
            describeArticles(null, null, null, offset, this.pageSize).then(response => {
                if(response.Code == Code.SUCCESS) {
                    this.articleObjs = response.Data.ArticleSet;
                    this.total = response.Data.Total;
                }
            });
        },
    }
}
</script>

<style scoped>
.content {
  position: relative;
  border-radius: 5px;
  height: 92.1%;
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
    font-size: 16px;
    border-radius: 3px;
    width: 100%;
    text-align: center;
}
.in-title {
    padding-top: 12px;
}
</style>
