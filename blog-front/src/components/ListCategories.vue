<!-- 首页 -->
<template>
<div>
    <div class="container">
        <div class="content">
            <!-- 列出类别 -->
            <el-row>
                <el-col
                    v-for="item in this.categoryObjs"
                    :key="item.Category.CategoryId"
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
                        <div class="categoryName">
                            <div class="in-title">
                                <a :href="'#/category?id='+item.Category.CategoryId" style="float: left; margin-left: 20px;">
                                    {{item.Category.CategoryName}}
                                </a>&#12288;
                                <div style="float:right; margin-right: 20px;">
                                    {{item.ArticleCount}}
                                </div>
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
    </div>
</div>
</template>

<script>
import {describeCategories} from "@/api/category";
import {Code} from '@/const/code.js'
import {message} from '@/utils/common'

export default {
    data() {  // 选项 / 数据
        return {
            title: null,
            isChange: false,  // 内容是否改变
            categoryObjs: [],
            formLabelWidth: "120px",
            currentPage: 1,
            total: null,
            pageSize: window.innerWidth <= 700? 8: 12,
            category: {
                CategoryId: null,
                CategoryName: null,
            },
        }
    },
    components: {  // 定义组件
    },
    created() {  // 生命周期函数
        this.handleListCategories();
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
        // 内容改变，触发监听
        contentChange: function() {
            var that = this;
            if(this.changeCount > 0) {
                that.isChange = true;
            }
            this.changeCount = this.changeCount + 1;
        },
        // 改变页码
        handleCurrentChange(val) {
            var that = this;
            this.currentPage = val; // 改变当前所指向的页数
            this.handleListCategories();
        },
        handleListCategories: function() {
            var offset = (this.currentPage - 1) * this.pageSize;
            describeCategories(null, offset, this.pageSize).then(response => {
                if(response.Code == Code.SUCCESS) {
                    this.categoryObjs = response.Data.CategorySet;
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
  height: 93.1%;
  margin-top: 38px;
  padding-top: 10px;
  background: rgba(230, 244, 249, 0.85);
  opacity: 0.98;
}
a {
  color: #fff; 
}
.categoryName {
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
    padding-top: 6px;
}
</style>
