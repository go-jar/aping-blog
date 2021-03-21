<!-- 首页 -->
<template>
<div>
    <el-row>
        <el-col
            v-for="item in this.articleObjs"
            :key="item.Id"
            style="padding: 6px"
            :xs="24"
            :sm="12"
            :md="12"
            :lg="6"
            :xl="4"
        >

            <el-card
                :body-style="{height: '37px', padding: '0px', textAlign: 'right'}"
                style="position: relative"
                shadow="always"
            >
                <div class="article-title">
                    <a :href="'#/article?id='+item.Article.Id">
                        {{item.Article.Title}}
                    </a>
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
    created() {  // 生命周期函数
        this.handleListArticles();
    },
    methods: {  // 事件处理器
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
            describeArticles(null, offset, this.pageSize).then(response => {
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
.article-title {
    position: absolute;
    left: 10px;
    top: 6px;
    z-index: 2;
    /*top: 15px;*/
    background: rgba(216, 236, 236, 0.85);
    color: #FFF;
    padding: 3px 8px;
    font-size: 14px;
    border-radius: 3px;
    width: 70%;
    text-align: center;
}
</style>
