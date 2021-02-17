<!-- 首页 -->
<template>
<div>
    <div class="container">
        <div class="editor">
            <el-form :model="article" :rules="articleRules" ref="article">
                <el-row>
                    <el-col :span="16">
                        <el-form-item label="标题" :label-width="formLabelWidth" prop="title">
                            <el-input v-model="article.title" auto-complete="off" @input="contentChange"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-row>
                    <el-col :span="6.5">
                        <el-form-item label="分类" :label-width="formLabelWidth" prop="category">
                        <el-select
                            v-model="article.category"
                            size="small"
                            placeholder="请选择"
                            style="width:150px"
                        >
                            <el-option
                                v-for="item in categories"
                                :key="item.id"
                                :label="item.categoryName"
                                :value="item.id"
                            ></el-option>
                        </el-select>
                        </el-form-item>
                    </el-col>

                    <el-col :span="6.5">
                        <el-form-item label="标签" label-width="80px">
                        <el-select
                            v-model="tagIds"
                            multiple
                            size="small"
                            placeholder="请选择"
                            style="width:210px"
                            filterable
                        >
                            <el-option
                                v-for="item in tags"
                                :key="item.id"
                                :label="item.tagName"
                                :value="item.id"
                            ></el-option>
                        </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-form-item label="内容" :label-width="formLabelWidth" prop="content">
                    <MarkdownEditor content="article.content" ref="editor" :height="300" @contentChange="contentChange"></MarkdownEditor>
                </el-form-item>

                <el-form-item style="float: right; margin-right: 20px;">
                    <el-button @click="dialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="submitArticle">确 定</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</div>
</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor'
import {Code} from '@/const/code.js'
import {setCookie, getCookie, delCookie} from "@/utils/cookie";
import {formatData} from "@/utils/web";

export default {
    data() {  // 选项 / 数据
        return {
            isChange: false,  // 文章内容是否改变
            categories: [],
            isEditArticle: false,
            changeCount: 0,  // 改变计数器
            formLabelWidth: "120px",
            title: "创建博客",
            tags: [],  // 标签数据
            tagIds: [],  // 保存选中标签 id (编辑时)
            article: {
                id: null,
                title: null,
                category: null,
                tags: null,
                content: null,
                isPublish: null,
                clickCount: 0,
            },
            articleRules: {
                title: [{required: true, message: '标题不能为空', trigger: 'blur'}],
                category: [{required: true, message: '类别不能为空', trigger: 'blur'}],
                tags: [{required: true, message: '标签不能为空', trigger: 'blur'}],
                content: [{required: true, message: '内容不能为空', trigger: 'blur'}],
            }
        }
    },
    components: {  // 定义组件
        MarkdownEditor,
    },
    created() {  // 生命周期函数
    },
    methods: {  // 事件处理器
        closeDialog(done) {
            if(this.isChange) {
                this.$confirm("是否关闭博客编辑窗口", "提示", {
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
                    this.$commonUtil.message.info("已取消")
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
                // 存放到cookie中，时间10天
                that.article.content = that.$refs.editor.getData(); //获取CKEditor中的内容
                that.article.tags = that.tags.join(",");
                setCookie("article", JSON.stringify(that.article), 10);
            }
            this.changeCount = this.changeCount + 1;
        },
        createArticle: function() {
            this.title = "增加博客"
            let that = this;
            
            let tempArticle = JSON.parse(getCookie("article"));
            if (tempArticle != null && tempArticle.title != null && tempArticle.title != "") {
                this.$confirm("还有上次未完成的博客编辑，是否继续编辑?", "提示", {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning"
                })
                .then(() => {
                    that.dialogVisible = true;
                    that.tags = [];
                    that.article = JSON.parse(getCookie("article"));
                    
                    var tags = that.article.tags.split(",");
                    for (var i = 0; i < tags.length; i++) {
                        if (tags[i] != null && tags[i] != "") {
                            that.tags.push(tags[i]);
                        }
                    }
                    
                    if(that.article.id) {
                        that.title = "编辑博客";
                        that.isEditArticle = true;
                    } else {
                        that.title = "新增博客";
                        that.isEditArticle = false;
                    }
                })
                .catch(() => {
                    that.dialogVisible = true;
                    that.article = that.getArticleObject();
                    that.$nextTick(() => {
                        //DOM现在更新了
                        that.$refs.editor.setData(that.article.content);  //设置富文本内容
                    });
                    that.tagValue = [];
                    that.isEditArticle = false;
                    that.title = "新增博客";
                    delCookie("article");
                });
            } else {
                that.dialogVisible = true;
                that.article = this.getArticleObject();

                that.$nextTick(() => {
                    //初始化内容
                    that.$refs.editor.initData();
                });

                that.tags = [];
                that.isEditArticle = false;
                that.formBak();
            }
        },
        getArticleObject: function() {
            var articleObject = {
                id: null,
                title: null,
                category: null,
                tags: null,
                content: null,
                isPublish: null,
                clickCount: 0,
            };
            return articleObject;
        },
        submitArticle: function() {
            if(this.tags.length <= 0) {
                this.$commonUtil.message.error("标签不能为空!")
                return;
            }

            this.$refs.article.validate((valid) => {
                if(!valid) {

                } else {
                    this.article.content = this.$refs.editor.getData();
                    this.article.tagIds = this.tagIds.join(",");
                    var params = formatData(this.article);
                    if (this.isEditArticle) {
                        editArticle(this.article).then(response => {
                            if (response.code == Code.SUCCESS) {
                                this.$commonUtil.message.success(response.message)
                                delCookie("article");
                                this.dialogFormVisible = false;
                                this.blogList();
                            } else {
                                this.$commonUtil.message.error(response.message)
                            }
                        });
                    } else {
                        createArticle(this.article).then(response => {
                            if (response.code == this.$ECode.SUCCESS) {
                                this.$commonUtil.message.success(response.message)
                                delCookie("article");
                                this.dialogVisible = false;
                                this.blogList();
                            } else {
                                this.$commonUtil.message.error(response.message)
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

<style>
.editor {
  position: relative;
  border-radius: 5px;
  height: 93.3%;
  margin-top: 38px;
}
</style>
