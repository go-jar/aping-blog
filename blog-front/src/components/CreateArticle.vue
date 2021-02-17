<!-- 首页 -->
<template>
<div>
    <div class="container">
        <div class="content">
            <el-form :model="article" :rules="articleRules" ref="article">
                <el-row>
                    <el-col :span="22">
                        <el-form-item label="标题" :label-width="formLabelWidth" prop="title">
                            <el-input v-model="article.title" auto-complete="off" @input="contentChange"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-row>
                    <el-col :span="10">
                        <el-form-item label="分类" :label-width="formLabelWidth" prop="category">
                            <el-select
                                v-model="article.categoryId"
                                size="small"
                                placeholder="请选择"
                                filterable
                                style="width: 350px;"
                            >
                                <el-option
                                    v-for="item in categoryObjs"
                                    :key="item.id"
                                    :label="item.categoryName"
                                    :value="item.id">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>

                    <el-col :span="10">
                        <el-form-item label="标签" :label-width="formLabelWidth">
                            <el-select
                                v-model="tagIds"
                                multiple
                                size="small"
                                placeholder="请选择"
                                filterable
                                style="width: 300px;"
                            >
                                <el-option
                                    v-for="item in tagObjs"
                                    :key="item.id"
                                    :label="item.tagName"
                                    :value="item.id">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>

                    <el-form-item style="float: right; margin-right: 90px;">
                        <el-button type="primary" @click="submitArticle">发布</el-button>
                    </el-form-item>
                </el-row>

                <el-row>
                    <el-col :span="24">
                        <el-form-item prop="content">
                            <MarkdownEditor :height="700" ref="editor" @contentChange="contentChange"></MarkdownEditor>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
        </div>
    </div>
</div>
</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor'
import {Code} from '@/const/code.js'
import {setCookie, getCookie, delCookie} from "@/utils/cookie";
import {message} from '@/utils/common'
import {formatData} from "@/utils/web";

export default {
    data() {  // 选项 / 数据
        return {
            isChange: false,  // 文章内容是否改变
            categoryObjs: [],  // 文章类别
            tagIds: [],  // 保存选中标签 id (编辑时)
            tagObjs: [],  // 标签数据
            isEditArticle: false,
            changeCount: 0,  // 改变计数器
            formLabelWidth: "120px",
            interval: null,  // 定义触发器
            article: {
                id: null,
                title: null,
                categoryId: null,
                tagIds: null,
                content: null,
                readCount: 0,
            },
            articleRules: {
                title: [{required: true, message: '标题不能为空', trigger: 'blur'}],
                content: [{required: true, message: '内容不能为空', trigger: 'blur'}],
            }
        }
    },
    components: {  // 定义组件
        MarkdownEditor,
    },
    created() {  // 生命周期函数
        this.createArticle()
    },
    methods: {  // 事件处理器
        // 内容改变，触发监听
        contentChange: function() {
            var that = this;
            if(this.changeCount > 0) {
                that.isChange = true;
                // 存放到 cookie 中，时间 10 天
                that.article.content = that.$refs.editor.getData();
                that.article.tagNames = that.tagNames.join(",");
                setCookie("article", JSON.stringify(that.article), 10);
            }
            this.changeCount = this.changeCount + 1;
        },
        createArticle: function() {
            let that = this;
            that.tagIds = [];
            let tempArticle = JSON.parse(getCookie("article"));
            if (tempArticle != null && tempArticle.title != null && tempArticle.title != "") {
                this.$confirm("还有上次未完成的博客编辑，是否继续编辑?", "提示", {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning"
                })
                .then(() => {
                    that.article = JSON.parse(getCookie("article"));
                    var tagIds = that.article.tagIds.split(",");
                    for (var i = 0; i < tagIds.length; i++) {
                        if (tagIds[i] != null && tagIds[i] != "") {
                            that.tagIds.push(tagIds[i]);
                        }
                    }
                    if(that.article.id) {
                        that.isEditArticle = true;
                    } else {
                        that.isEditArticle = false;
                    }
                })
                .catch(() => {
                    that.article = that.getInitArticleObject();
                    that.$nextTick(() => {
                        // DOM现在更新了
                        that.$refs.editor.setData(that.article.content);  // 设置富文本内容
                    });
                    that.tagIds = [];
                    that.isEditArticle = false;
                    delCookie("article");
                });
            } else {
                that.article = this.getInitArticleObject();

                that.$nextTick(() => {
                    // 初始化内容
                    that.$refs.editor.initData();
                });

                that.tagIds = [];
                that.isEditArticle = false;
                that.backUpArticle();
            }
        },
        getInitArticleObject: function() {
            var articleObject = {
                id: null,
                title: null,
                categoryId: null,
                tagIds: null,
                content: null,
                readCount: 0,
            };
            return articleObject;
        },
        // 备份文章
        backUpArticle: function() {
            var that = this;
            that.interval = setInterval(function() {
                if (that.article.title != null && that.article.title != "") {
                    // 存放到 cookie 中，时间 10 天
                    that.article.content = that.$refs.editor.getData(); //获取CKEditor中的内容
                    that.article.tagIds = that.tagIds.join(",");
                    setCookie("article", JSON.stringify(that.article), 10);
                }
            }, 10000);
        },
        submitArticle: function() {
            this.$refs.article.validate((valid) => {
                if(!valid) {

                } else {
                    this.article.content = this.$refs.editor.getData();
                    this.article.tagIds = this.tagIds.join(",");
                    var params = formatData(this.article);
                    if (this.isEditArticle) {
                        editArticle(this.article).then(response => {
                            if (response.code == Code.SUCCESS) {
                                message.success(response.message)
                                delCookie("article");
                                this.dialogFormVisible = false;
                                this.blogList();
                            } else {
                                message.error(response.message)
                            }
                        });
                    } else {
                        createArticle(this.article).then(response => {
                            if (response.code == this.$ECode.SUCCESS) {
                                message.success(response.message)
                                delCookie("article");
                                this.dialogVisible = false;
                                this.blogList();
                            } else {
                                message.error(response.message)
                            }
                        });
                    }
                }
            })
        },
        articleCategoryList: function() {
            var articleCategoryParams = {};
            articleCategoryParams.pageSize = 100;
            articleCategoryParams.currentPage = 1;
            getArticleCategoryList(articleCategoryParams).then(response => {
                if(response.code == Code.SUCCESS) {
                    this.categories = response.data.records;
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
</style>
