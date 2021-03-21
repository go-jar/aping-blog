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
                                <el-option v-for="item in categoryObjs" :key="item.Id" :label="item.CategoryName" :value="item.Id">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>

                    <el-col :span="10">
                        <el-form-item label="标签" :label-width="formLabelWidth">
                            <el-select v-model="tagIds" multiple size="small" placeholder="请选择" filterable style="width: 300px;">
                                <el-option v-for="item in tagObjs" :key="item.Id" :label="item.TagName" :value="item.Id">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>

                    <el-form-item style="float: right; margin-right: 90px;">
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
import {describeArticles, createArticle, updateArticle} from "@/api/article";
import {describeCategories} from "@/api/category";
import {describeTags} from "@/api/tag";
import MarkdownEditor from '@/components/MarkdownEditor'
import {Code} from '@/const/code.js'
import {setCookie, getCookie, delCookie} from "@/utils/cookie";
import {message} from '@/utils/common'
import {formatData} from "@/utils/web";

export default {
    name: 'EditArticle',
    data() { // 选项 / 数据
        return {
            isChange: false, // 文章内容是否改变
            categoryObjs: [], // 文章类别
            tagIds: [], // 保存选中标签 id (编辑时)
            tagObjs: [], // 标签数据
            isUpdateArticle: false,
            changeCount: 0, // 改变计数器
            formLabelWidth: "120px",
            interval: null, // 定义触发器
            article: {
                Id: null,
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
        this.handleRouteChange()
    },
    mounted() {
    },
    methods: { // 事件处理器
        handleRouteChange:function(){
            var that = this;
            that.articleId = that.$route.query.id==undefined? null: parseInt(that.$route.query.id);  // 获取传参的 id
            
            //获取详情接口
            if (that.articleId != null) {
                describeArticles(that.articleId, null, null, null, null).then(response => {
                    if (response.Data.ArticleSet == null || response.Data.ArticleSet.length == 0) {
                        return;
                    }

                    var editArticle = response.Data.ArticleSet[0];
                    if (editArticle != null) {
                        var tagIds = [];
                        if (editArticle.TagSet != null) {
                            for (let i=  0; i < editArticle.TagSet.length; i++) {
                                tagIds.push(editArticle.TagSet[i].Id);
                            }
                        }
                        that.tagIds = tagIds;

                        that.article = {
                            Id: editArticle.Article.Id,
                            Title: editArticle.Article.Title,
                            CategoryId: editArticle.Article.CategoryId,
                            Content: editArticle.Article.Content,
                        };
                    }
                });
                that.isUpdateArticle = true;
            } else {
                this.handleLoadBackUpArticle()
            }
        },
        handleLoadBackUpArticle: function() {
            let that = this;
            if (this.articleId != null) {
                that.isUpdateArticle = true;
                return;
            }

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

                    if (!that.article.Id) {
                        that.isUpdateArticle = false;
                    } else {
                        that.articleId = that.article.Id;
                    }
                })
                .catch(() => {
                    that.article = that.getInitArticleObject();
                    that.$nextTick(() => {
                        // DOM现在更新了
                        that.$refs.editor.setData(that.article.Content); // 设置富文本内容
                    });
                    that.tagIds = [];
                    that.isUpdateArticle = false;
                    delCookie("article");
                });
            } else {
                that.article = this.getInitArticleObject();
                that.tagIds = [];
                that.isUpdateArticle = false;
                that.backUpArticle();
            }
        },
        getInitArticleObject: function () {
            var articleObject = {
                Id: null,
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
            this.$refs.article.validate((valid) => {
                if (!valid) {

                } else {
                    this.article.Content = this.$refs.editor.getData();
                    this.article.TagIds = this.tagIds.join(",");
                    var params = formatData(this.article);
                    if (this.isUpdateArticle) {
                        updateArticle(this.article).then(response => {
                            if (response.Code === Code.SUCCESS) {
                                message.success(response.Msg)
                                delCookie("article");
                                this.dialogFormVisible = false;
                            } else {
                                message.error(response.Msg)
                            }
                        });
                    } else {
                        createArticle(this.article).then(response => {
                            if (response.Code === Code.SUCCESS) {
                                message.success(response.Msg)
                                delCookie("article");
                                this.dialogVisible = false;
                                this.articleId = response.Data.ArticleId;


                            } else {
                                message.error(response.Msg)
                            }
                        });
                    }
                }
            })

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
