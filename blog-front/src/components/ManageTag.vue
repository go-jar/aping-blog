<!-- 首页 -->
<template>
<div>
    <div class="container">
        <div class="content">
            <div class="filter-container" style="margin: 3px 0 10px 6px;">
                <el-button class="filter-item" type="primary" @click="handleCreateTag" icon="el-icon-edit">添加</el-button>
            </div>

            <!-- 列出标签 -->
            <el-row>
                <el-col
                    v-for="item in this.tagObjs"
                    :key="item.Tag.TagId"
                    style="padding: 6px"
                    :xs="24"
                    :sm="52"
                    :md="12"
                    :lg="66"
                    :xl="4"
                >

                    <el-card
                        :body-style="{height: '35px', padding: '0px', textAlign: 'right'}"
                        style="position: relative"
                        shadow="always"
                    >
                        <div class="TagName">
                            <a :href="'#/tag?id='+item.Tag.TagId" style="float: left; margin-left: 30px; color: #fff;">
                                {{item.Tag.TagName}}
                            </a>&#12288;
                            <div style="float:right; margin-right: 20px;">{{item.ArticleCount}}</div>
                        </div>
                       
                        <div style="height: 25px; margin-top: 5px; margin-right: 6px">
                            <el-button-group>
                                <el-tooltip class="item" effect="dark" content="编辑" placement="bottom-start" style="margin-right: 2px">
                                    <el-button
                                        type="primary"
                                        size="mini"
                                        icon="el-icon-document-copy"
                                        @click="handleUpdateTag(item.Tag)"
                                    >
                                    </el-button>
                                </el-tooltip>

                                <el-tooltip class="item" effect="dark" content="删除" placement="bottom-start" style="margin-right: 2px">
                                    <el-button
                                        type="danger"
                                        size="mini"
                                        icon="el-icon-delete"
                                        @click="handleDeleteTag(item.Tag)"
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
            >
                <el-form :model="tag" :rules="tagRules" ref="tag">
                    <el-row>
                        <el-col :span="22">
                            <el-form-item label="标签" :label-width="formLabelWidth">
                                <el-input v-model="tag.TagName" auto-complete="off" @input="contentChange" ref="tagName" prop="tagName"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="22">
                            <el-form-item label="序号" :label-width="formLabelWidth">
                                <el-input v-model="tag.TagIndex" auto-complete="off" @input="contentChange" ref="tagIndex" prop="tagIndex"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>

                    <el-form-item style="float: right; margin-right: 66px;">
                        <el-button @click="dialogVisible = false">取 消</el-button>
                        <el-button type="primary" @click="submitTag">确 定</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>
        </div>
    </div>
</div>
</template>

<script>
import {createTag, deleteTag, modifyTag, describeTags} from "@/api/tag";
import {Code} from '@/const/code.js'
import {message} from '@/utils/common'

export default {
    data() {  // 选项 / 数据
        return {
            title: null,
            dialogVisible: false,  // 控制弹出框
            isChange: false,  // 内容是否改变
            changeCount: 0,  // 改变计数器
            isUpdateTag: false,
            tagObjs: [],
            formLabelWidth: "120px",
            currentPage: 1,
            total: null,
            pageSize: 40,
            tag: {
                TagId: null,
                TagName: null,
                TagIndex: null,
            },
            tagRules: {
                tagName: [{required: true, message: '标签不能为空', trigger: 'blur'}],
                tagIndex: [{required: true, message: '序号不能为空', trigger: 'blur'}],
            }
        }
    },
    components: {  // 定义组件
    },
    created() {  // 生命周期函数
        this.handleListTags();
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
        handleCreateTag: function() {
            this.title = "创建标签"
            let that = this;
            that.dialogVisible = true;
            that.isUpdateTag = false;

            if (that.tag != null && that.tag.TagName != null && that.tag.TagName != "") {
                if (that.tag.TagId) {
                    that.isUpdateTag = true;
                } else {
                    that.isUpdateTag = false;
                }
            } else {
                that.tag = this.getInitTagObject();
                that.isUpdateTag = false;
            }
        },
        getInitTagObject: function() {
            var tagObject = {
                TagId: null,
                TagName: null,
                TagIndex: null,
            };
            return tagObject;
        },
        handleDeleteTag: function(row) {
            var that = this;
            this.$confirm("此操作将把标签删除, 是否继续?", "提示", {
                confirmButtonText: "确定",
                cancelButtonText: "取消",
                type: "warning"
            })
            .then(() => {
                deleteTag(row.TagId).then(response => {
                    if (response.Code == Code.SUCCESS) {
                        message.success(response.message)
                    } else {
                        message.error(response.message)
                    }
                    this.handleListTags();
                });
            })
            .catch(() => {
                message.info("已取消删除")
            });
        },
        handleUpdateTag: function(row) {
            var that = this;
            that.title = "编辑标签";
            that.tag = row;

            that.dialogVisible = true;
            that.isUpdateTag = true;
        },
        // 改变页码
        handleCurrentChange(val) {
            var that = this;
            this.currentPage = val; // 改变当前所指向的页数
            this.handleListTags();
        },
        handleListTags: function() {
            var offset = (this.currentPage - 1) * this.pageSize;
            describeTags(null, offset, this.pageSize).then(response => {
                if(response.Code == Code.SUCCESS) {
                    this.tagObjs = response.Data.TagSet;
                    this.total = response.Data.Total;
                }
            });
        },
        submitTag: function() {
            this.$refs.tag.validate((valid) => {
                if(!valid) {
                    console.log("校验出错")
                    return
                }
                
                if (this.isUpdateTag) {
                    modifyTag(this.tag).then(response => {
                        if (response.Code == Code.SUCCESS) {
                            message.success(response.Message)
                            this.dialogVisible = false;
                            this.handleListTags();
                        } else {
                            message.error(response.Message)
                        }
                    });
                } else {
                    createTag(this.tag).then(response => {
                        if (response.Code == Code.SUCCESS) {
                            message.success(response.Message)
                            this.dialogVisible = false;
                            this.handleListTags();
                        } else {
                            message.error(response.Message)
                        }
                    });
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
  height: 92.1%;
  margin-top: 38px;
  padding-top: 10px;
  background: rgba(230, 244, 249, 0.85);
  opacity: 0.98;
}
.tagName {
    position: absolute;
    left: 10px;
    top: 6px;
    z-index: 2;
    background: rgba(58, 59, 59, 0.85);
    color: #FFF;
    padding: 3px 8px;
    font-size: 15px;
    border-radius: 3px;
    width: 75%;
    text-align: center;
}
</style>
