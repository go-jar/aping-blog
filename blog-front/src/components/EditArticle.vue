<!-- 首页 -->
<template>
<div>
    <!-- <div class="container"> -->
        <div class="content">
            <el-form :model="article" :rules="articleRules" ref="article">
                <el-row>
                    <el-col :span="22">
                        <el-form-item label="标题" :label-width="formLabelWidth" prop="Title">
                            <el-input v-model="article.Article.Title" auto-complete="off" @input="handleContentChange"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-row>
                    <el-col :span="10">
                        <el-form-item label="分类" :label-width="formLabelWidth" prop="CategoryId">
                            <el-select v-model="article.Article.CategoryId" size="small" placeholder="请选择" filterable style="width: 350px;">
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
                            <MarkdownEditor ref="editor" :height="730" :content="article.Article.Content"></MarkdownEditor>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
        </div>
    </div>
<!-- </div> -->
</template>

<script>
import {describeArticles, createArticle, modifyArticle} from "@/api/article"
import {describeCategories} from "@/api/category"
import {describeTags} from "@/api/tag"
import MarkdownEditor from '@/components/MarkdownEditor'
import {Code} from '@/const/code.js'
import {setCookie, getCookie, delCookie} from "@/utils/cookie"
import {message} from '@/utils/common'
import {formatData} from "@/utils/web"


export default {
    name: 'EditArticle',
    data() { // 选项 / 数据
        return {
            categoryObjs: [], // 文章类别
            tagIds: [], // 保存选中标签 id (编辑时)
            tagObjs: [], // 标签数据
            formLabelWidth: "120px",
            article: null,
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
                that.isUpdateArticle = true;

                describeArticles(that.articleId, null, null, null, null).then(response => {
                    if (response.Data.ArticleSet == null || response.Data.ArticleSet.length == 0) {
                        return;
                    }

                    var editArticle = response.Data.ArticleSet[0];
                    if (editArticle != null) {
                        var tagIds = [];
                        if (editArticle.TagSet != null) {
                            for (let i =  0; i < editArticle.TagSet.length; i++) {
                                tagIds.push(editArticle.TagSet[i].TagId);
                            }
                        }

                        that.tagIds = tagIds;
                        that.article = editArticle;
                    }
                });
            } 
        },
        handleSubmitArticle: function () {
            this.article.Article.Content = this.$refs.editor.getData();
            this.article.Article.TagIds = this.tagIds.join(",");
                    
            modifyArticle(this.article.Article).then(response => {
                if (response.Code === Code.SUCCESS) {
                    message.success(response.Msg)
                    delCookie("article");
                    this.dialogFormVisible = false;
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
    height: 93.1%;
    margin-top: 38px;
    padding-top: 10px;
    background: rgba(230, 244, 249, 0.85);
    opacity: 0.98;
}
</style>
