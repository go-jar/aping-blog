<!-- 首页 -->
<template>
<div>
    <div class="container">
        <div class="content">
            <div class="filter-container" style="margin: 3px 0 10px 6px;">
                <el-button class="filter-item" type="primary" @click="handleCreateCategory" icon="el-icon-edit">添加</el-button>
            </div>

            <!-- 列出类别 -->
            <el-row>
                <el-col
                    v-for="item in this.categoryObjs"
                    :key="item.Id"
                    style="padding: 6px"
                    :xs="24"
                    :sm="12"
                    :md="12"
                    :lg="6"
                    :xl="4"
                >

                    <el-card
                        :body-style="{padding: '0px', textAlign: 'right'}"
                        style="position: relative"
                        shadow="always"
                    >
                        <div class="categoryName">{{item.CategoryName}}</div>
                       
                        <div style="height: 25px; margin-top: 14px; margin-right: 6px">
                            <el-button-group>
                                <el-tooltip class="item" effect="dark" content="类别" placement="bottom-start" style="margin-right: 2px">
                                    <el-button
                                        size="mini"
                                        icon="el-icon-copy-document"
                                        @click="goSubjectItem(item)"
                                    />`
                                    </el-tooltip>

                                <el-tooltip class="item" effect="dark" content="编辑" placement="bottom-start" style="margin-right: 2px">
                                    <el-button
                                        type="primary"
                                        size="mini"
                                        icon="el-icon-document-copy"
                                        @click="handleUpdateCategory(item)"
                                    >
                                    </el-button>
                                </el-tooltip>

                                <el-tooltip class="item" effect="dark" content="删除" placement="bottom-start" style="margin-right: 2px">
                                    <el-button
                                        type="danger"
                                        size="mini"
                                        icon="el-icon-delete"
                                        @click="handleDeleteCategory(item)"
                                    />
                                </el-tooltip>
                            </el-button-group>
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

            <!-- 添加或修改对话框 -->
            <el-dialog
                :title="title"
                :visible.sync="dialogVisible"
                :before-close="closeDialog"
                fullscreen
                v-focus
            >
                <el-form :model="category" :rules="categoryRules" ref="category">
                    <el-row>
                        <el-col :span="22">
                            <el-form-item label="类别" :label-width="formLabelWidth">
                                <el-input v-model="category.CategoryName" auto-complete="off" @input="contentChange" ref="categoryName" v-focus></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>

                    <el-form-item style="float: right; margin-right: 66px;">
                        <el-button @click="dialogVisible = false">取 消</el-button>
                        <el-button type="primary" @click="submitCategory">确 定</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>
        </div>
    </div>
</div>
</template>

<script>
import {createCategory, deleteCategory, updateCategory, describeCategories} from "@/api/category";
import {Code} from '@/const/code.js'
import {message} from '@/utils/common'

export default {
    data() {  // 选项 / 数据
        return {
            title: null,
            dialogVisible: false,  // 控制弹出框
            isChange: false,  // 内容是否改变
            changeCount: 0,  // 改变计数器
            isUpdateCategory: false,
            categoryObjs: [],
            formLabelWidth: "120px",
            currentPage: 1,
            total: null,
            pageSize: 40,
            category: {
                Id: null,
                CategoryName: null,
            },
            categoryRules: {
                CategoryName: [{required: true, message: '类别不能为空', trigger: 'blur'}],
            }
        }
    },
    components: {  // 定义组件
    },
    created() {  // 生命周期函数
        this.handleListCategories();
        this.directives();
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
        handleCreateCategory: function() {
            this.title = "创建标签"
            let that = this;
            that.dialogVisible = true;

            if (that.category != null && that.category.CategoryName != null && that.category.CategoryName != "") {
                if(that.category.Id) {
                    that.isUpdateCategory = true;
                } else {
                    that.isUpdateCategory = false;
                }
            } else {
                that.category = this.getInitCategoryObject();
                that.isUpdateCategory = false;
            }

            that.$nextTick(() => {
                that.$refs.categoryName.focus()
            })
        },
        getInitCategoryObject: function() {
            var categoryObject = {
                Id: null,
                CategoryName: null,
            };
            return categoryObject;
        },
        handleDeleteCategory: function(row) {
            var that = this;
            this.$confirm("此操作将把类别删除, 是否继续?", "提示", {
                confirmButtonText: "确定",
                cancelButtonText: "取消",
                type: "warning"
            })
            .then(() => {
                deleteCategory(row.Id).then(response => {
                    if (response.Code == Code.SUCCESS) {
                        message.success(response.message)
                    } else {
                        message.error(response.message)
                    }
                    this.handleListCategories();
                });
            })
            .catch(() => {
                message.info("已取消删除")
            });
        },
        handleUpdateCategory: function(row) {
            var that = this;
            that.title = "编辑类别";
            that.category = row;

            that.dialogVisible = true;
            that.isUpdateCategory = true;

            that.$nextTick(() => {
                that.$refs.categoryName.focus()
            })
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
        submitCategory: function() {
            this.$refs.category.validate((valid) => {
                if(!valid) {
                    console.log("校验出错")
                    return
                }
                
                if (this.isUpdateCategory) {
                    updateCategory(this.category.Id, this.category.CategoryName).then(response => {
                        if (response.Code == Code.SUCCESS) {
                            message.success(response.Message)
                            this.dialogVisible = false;
                            this.handleListCategories();
                        } else {
                            message.error(response.Message)
                        }
                    });
                } else {
                    createCategory(this.category.CategoryName).then(response => {
                        if (response.Code == Code.SUCCESS) {
                            message.success(response.Message)
                            this.dialogVisible = false;
                            this.handleListCategories();
                        } else {
                            message.error(response.Message)
                        }
                    });
                }
            })
        },
        /**
         * 让打开页面先获取焦点
         */
        directives: {
            'focus': {
                inserted (el, binding, vnode) {
                    // 聚焦元素
                    el.focus()
                }
            }
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
.categoryName {
    position: absolute;
    left: 10px;
    top: 6px;
    z-index: 2;
    /*top: 15px;*/
    background: rgba(58, 59, 59, 0.85);
    color: #FFF;
    padding: 3px 8px;
    font-size: 14px;
    border-radius: 3px;
    width: 70%;
    text-align: center;
}
</style>
