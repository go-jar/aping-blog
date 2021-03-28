<!-- 首页 -->
<template>
<div>
    <!-- <div class="container"> -->
        <div class="content">
            <el-form :model="article" :rules="articleRules" ref="article">
                <el-row>
                    <el-col :span="22">
                        <el-form-item label="标题" :label-width="formLabelWidth" prop="Title">
                            <el-input v-model="article.Title" auto-complete="off" @input="handleContentChange"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-row>
                    <el-col :span="10">
                        <el-form-item label="分类" :label-width="formLabelWidth" prop="CategoryId">
                            <el-select v-model="article.CategoryId" size="small" placeholder="请选择" filterable style="width: 350px;">
                                <el-option v-for="item in categoryObjs" :key="item.Category.CategoryId" :label="item.Category.CategoryName" :value="item.Category.CategoryId">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>

                    <el-col :span="10">
                        <el-form-item label="标签" :label-width="formLabelWidth">
                            <el-select v-model="tagIds" multiple size="small" placeholder="请选择" filterable style="width: 300px;">
                                <el-option v-for="item in tagObjs" :key="item.Tag.TagId" :label="item.Tag.TagName" :value="item.Tag.TagId">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>

                    <el-form-item style="float: right; margin-right: 100px;">
                        <el-button type="primary" @click="handleSubmitArticle">发布</el-button>
                    </el-form-item>
                </el-row>

                <el-row>
                    <el-col :span="24">
                        <el-form-item prop="Content">
                            <MarkdownEditor ref="editor" :height="700" :content="article.Content" @handleContentChange="handleContentChange"></MarkdownEditor>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
        </div>
    </div>
<!-- </div> -->
</template>

<script>
import {describeArticles, createArticle, updateArticle} from "@/api/article"
import {describeCategories} from "@/api/category"
import {describeTags} from "@/api/tag"
import MarkdownEditor from '@/components/MarkdownEditor'
import {Code} from '@/const/code.js'
import {setCookie, getCookie, delCookie} from "@/utils/cookie"
import {message} from '@/utils/common'
import {formatData} from "@/utils/web"


export default {
    name: 'CreateArticle',
    data() { // 选项 / 数据
        return {
            isChange: false, // 文章内容是否改变
            categoryObjs: [], // 文章类别
            tagIds: [], // 保存选中标签 id (编辑时)
            tagObjs: [], // 标签数据
            changeCount: 0, // 改变计数器
            formLabelWidth: "120px",
            interval: null, // 定义触发器
            article: { 
                ArticleId: null,
                Title: null,
                CategoryId: null,
                Content: null,
                ReadCount: 0,
                TagIds: null,
            },
            articleRules: {
                Title: [{
                    required: true,
                    message: '标题不能为空',
                    trigger: 'blur',
                    type: 'string'
                }],
                CategoryId: [{
                    required: true,
                    message: '类别不能为空',
                    trigger: 'blur',
                    type: 'number'
                }],
            },
            articleId: null,
        }
    },
    components: { // 定义组件
        MarkdownEditor,
    },
    created() { // 生命周期函数
        this.handleListCategories()
        this.handleListTags()
        this.handleLoadBackUpArticle()
    },
    mounted() {
    },
    methods: { // 事件处理器
        handleLoadBackUpArticle: function() {
            let that = this;

            let tempArticle = JSON.parse(getCookie("article"));
            if (tempArticle != null && tempArticle.Title != null && tempArticle.Title != "") {
                this.$confirm("还有上次未完成的博客编辑，是否继续编辑?", "提示", {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning"
                })
                .then(() => {
                    that.article = JSON.parse(getCookie("article"));
                    var tagIds = that.article.TagIds.split(",");
                    for (var i = 0; i < tagIds.length; i++) {
                        if (tagIds[i] != null && tagIds[i] != "") {
                            that.tagIds.push(tagIds[i]);
                        }
                    }
                
                    that.$nextTick(() => {
                        // DOM现在更新了
                        that.$refs.editor.setData(that.article.Content); // 设置富文本内容
                    });

                    if (!that.article.Article.ArticleId) {
                        that.isUpdateArticle = false;
                    } else {
                        that.articleId = that.article.Article.ArticleId;
                    }
                })
                .catch(() => {
                    that.article = that.getInitArticleObject();
                    that.$nextTick(() => {
                        // DOM现在更新了
                        that.$refs.editor.setData(that.article.Content); // 设置富文本内容
                    });
                    that.tagIds = [];
                    delCookie("article");
                });
            } else {
                that.article = this.getInitArticleObject();
                that.tagIds = [];
                that.backUpArticle();
            }
        },
        getInitArticleObject: function () {
            var articleObject = {
                ArticleId: null,
                Title: null,
                CategoryId: null,
                Content: null,
                ReadCount: 0,
                TagIds: null,
            };
            return articleObject;
        },
        // 备份文章
        backUpArticle: function () {
            var that = this;
            that.interval = setInterval(function () {
                if (that.article.Title != null && that.article.Title != "") {
                    // 存放到 cookie 中，时间 10 天
                    that.article.Content = that.$refs.editor.getData();
                    that.article.TagIds = that.tagIds.join(",");
                    setCookie("article", JSON.stringify(that.article), 10);
                }
            }, 10000);
        },
        // 内容改变，触发监听
        handleContentChange: function () {
            var that = this;
            if (this.changeCount > 0) {
                that.isChange = true;
                // 存放到 cookie 中，时间 10 天
                that.article.Content = that.$refs.editor.getData();
                that.article.TagIds = that.tagIds.join(",");
                setCookie("article", JSON.stringify(that.article), 10);
            }
            this.changeCount = this.changeCount + 1;
        },
        handleSubmitArticle: function () {
            this.article.Content = this.$refs.editor.getData();
            this.article.TagIds = this.tagIds.join(",");

            createArticle(this.article).then(response => {
                if (response.Code === Code.SUCCESS) {
                    message.success(response.Msg)
                    delCookie("article");
                    this.dialogVisible = false;
                    this.articleId = response.Data.Id;
                } else {
                    message.error(response.Msg)
                }
            });

            if (this.articleId != null) {
                this.$router.push({
                    path: '/article',
                    query: {
                        "id": this.articleId,
                    }
                });
            } else {
                this.$router.push({
                    path: '/',
                });
            }
        },
        handleListCategories: function () {
            describeCategories(null, 0, 500).then(response => {
                if (response.Code == Code.SUCCESS) {
                    this.categoryObjs = response.Data.CategorySet;
                }
            });
        },
        handleListTags: function () {
            describeTags(null, 0, 500).then(response => {
                if (response.Code == Code.SUCCESS) {
                    this.tagObjs = response.Data.TagSet;
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
