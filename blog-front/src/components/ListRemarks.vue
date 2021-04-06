<!-- 首页 -->
<template>
<div>
    <div class="container">
        <div class="content">
            <!-- 列出类评论-->
            <el-row>
                <el-col
                    v-for="item in this.remarkSet"
                    :key="item.Remark.RemarkId"
                    style="padding: 3px"
                >
          
                    <el-card class="back"
                         :body-style="{height: '50px', padding: '0px', textAlign: 'center'}"
                         style="position: relative"
                         shadow="always"
                    >

                        <div class="remark-content">
                            <div class="in-remark-content">
                                {{formatTime(item.Remark.CreatedTime)}}&#12288;
                                {{item.Remark.Nickname}}&#12288;
                                {{item.Remark.Content}}&#12288;
                                <a :href="'#/article?id='+item.Remark.ArticleId">文章</a>
                                <a :href="'#'" @click="handleReplyRemark(item.Remark)">回复</a>
                                <el-link
                                    type="danger"
                                    size="mini"
                                    icon="el-icon-edit"
                                    @click="handleUpdateRemark(item.Remark)"
                                    v-show="adminLogin"
                                />
                                <el-link
                                    type="danger"
                                    size="mini"
                                    icon="el-icon-delete"
                                    @click="handleDeleteRemark(item.Remark.RemarkId)" 
                                    v-show="adminLogin"
                                />
                            </div>
                        </div>
                        
                         <!-- 列出类回复-->
                        <el-row>
                            <el-col class="back"
                                v-for="reply in item.ReplySet"
                                :key="reply.RemarkId"
                                style="padding-left: 20px"
                            >

                                <el-card
                                    :body-style="{height: '50px', padding: '0px', textAlign: 'right'}"
                                    style="position: relative"
                                    shadow="always"
                                >
                                    <div class="remark-content">
                                        <div class="in-remark-content">
                                            {{formatTime(reply.CreatedTime)}}&#12288;
                                            {{reply.Nickname}}@{{reply.NicknameReplied}}&#12288;
                                            {{reply.Content}}&#12288;
                                            <a :href="'#'" @click="handleReplyRemark(reply)">回复</a>
                                            <el-link
                                                type="danger"
                                                size="mini"
                                                icon="el-icon-edit"
                                                @click="handleUpdateRemark(reply)" 
                                                v-show="adminLogin"
                                            />
                                            <el-link
                                                type="danger"
                                                size="mini"
                                                icon="el-icon-delete"
                                                @click="handleDeleteRemark(reply.RemarkId)" 
                                                v-show="adminLogin"
                                            />
                                        </div>
                                    </div>
                                </el-card>
                            </el-col>
                        </el-row>

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
                style="height: 100%;"
            >
                <el-form :model="remark" :rules="remarkRules" ref="remark">
                    <el-row>
                        <el-col :span="22">
                            <el-form-item label="您的昵称" :label-width="formLabelWidth">
                                <el-input v-model="remark.Nickname" auto-complete="off" @input="contentChange" ref="nickname" prop="Nickname"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="22">
                            <el-form-item label="评论内容" :label-width="formLabelWidth">
                                <el-input v-model="remark.Content" auto-complete="off" @input="contentChange" ref="content" prop="Content"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>

                    <div style="float:bottome; margin-bottom:15px;"> 
                        <el-form-item style="float: right; margin-right: 66px;">
                            <el-button @click="dialogVisible=false">取 消</el-button>
                            <el-button type="primary" @click="submitRemark">确 定</el-button>
                        </el-form-item>
                    </div>
                </el-form>
            </el-dialog>
        </div>
    </div>
</div>
</template>

<script>
import {describeArticles} from "@/api/article";
import {createRemark, deleteRemark, modifyRemark, describeRemarks} from "@/api/remark";
import {Code} from '@/const/code.js'
import {getToken} from '@/utils/auth.js'
import {message} from '@/utils/common'

export default {
    data() {  // 选项 / 数据
        return {
            articleId: null,
            adminLogin: false,
            title: null,
            dialogVisible: false,  // 控制弹出框
            isChange: false,  // 内容是否改变
            changeCount: 0,  // 改变计数器
            isUpdateRemark: false,
            article: {
                Title: "",
                ArticleId: 0,
            },
            remarkSet: [],
            formLabelWidth: "120px",
            currentPage: 1,
            total: null,
           pageSize: window.innerWidth <= 700? 7: 12,
            remark: {
                RemarkId: null,
                ArticleId: null,
                Nickname: null,
                Content: null,
                InitRemarkId: null,
                NicknameReplied: null,
            },
            remarkRules: {
                content: [{required: true, message: '类别不能为空', trigger: 'blur'}],
                remarkIndex: [{required: true, message: '序号不能为空', trigger: 'blur'}],
            }
        }
    },
    components: {  // 定义组件
    },
    watch: {
        // 如果路由有变化，会再次执行该方法
        '$route':'routeChange'
    },
    created() {  // 生命周期函数
        this.routeChange();

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
    methods: {  // 事件处理器
        // 对时间进行格式化
        formatTime: function(createdTime) {
            if (createdTime) {
                const dt = new Date(createdTime)
                return dt.format("yyyy-MM-dd");
            }
            return '';
        },
        routeChange: function(){
            var that = this;

            if (getToken()) {
				that.adminLogin = true;
			} else {
				that.adminLogin = false;
			}

            that.articleId = that.$route.query.articleid == undefined? 0: parseInt(that.$route.query.articleid);
            
            //获取详情接口
            that.handleGetArticle();
            that.handleListRemarks();
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
        // 内容改变，触发监听
        contentChange: function() {
            var that = this;
            if(this.changeCount > 0) {
                that.isChange = true;
            }
            this.changeCount = this.changeCount + 1;
        },
        getInitRemarkObject: function() {
            var remarkObject = {
                RemarkId: null,
                ArticleId: null,
                Nickname: null,
                Content: null,
                InitRemarkId: null,
                NicknameReplied: null,
            };
            return remarkObject;
        },
        handleDeleteRemark: function(remarkId) {
            var that = this;
            this.$confirm("此操作将把评论删除, 是否继续?", "提示", {
                confirmButtonText: "确定",
                cancelButtonText: "取消",
                type: "warning"
            })
            .then(() => {
                deleteRemark(remarkId).then(response => {
                    if (response.Code == Code.SUCCESS) {
                        message.success(response.message)
                    } else {
                        message.error(response.message)
                    }

                    this.handleListRemarks();

                    this.$router.push({
                        path: '/remark',
                        query: {
                            "articleid": this.articleId,
                        }
                    });
                });
            })
            .catch(() => {
                message.info("已取消删除")
            });
        },
        handleUpdateRemark: function(row) {
            var that = this;
            that.title = "编辑评论";
            that.remark = row;
            that.dialogVisible = true;
            that.isUpdateRemark = true;
        },
        handleReplyRemark: function(row) {
            var that = this;
            that.title = "回复评论";
            that.dialogVisible = true;

            that.remark.InitRemarkId = row.InitRemarkId == 0? row.RemarkId: row.InitRemarkId;
            that.remark.NicknameReplied = row.Nickname;
        },
        // 改变页码
        handleCurrentChange(val) {
            var that = this;
            this.currentPage = val; // 改变当前所指向的页数
            this.handleListRemarks();
        },
        handleListRemarks: function() {
            var offset = (this.currentPage - 1) * this.pageSize;
            describeRemarks(null, offset, this.pageSize).then(response => {
                if(response.Code == Code.SUCCESS) {
                    this.remarkSet = response.Data.RemarkSet;
                    this.total = response.Data.Total;
                }
            });
        },
        handleGetArticle: function() {
            describeArticles(this.articleId, null, null, null, null, null).then(response => {
                if(response.Code == Code.SUCCESS) {
                    if (response.Data.ArticleSet.length == 0) {
                        message.error("没有查到文章")
                    } else {
                        this.article = response.Data.ArticleSet[0].Article;
                    }
                }
            });
        },
        submitRemark: function() {
            this.$refs.remark.validate((valid) => {
                if(!valid) {
                    console.log("校验出错")
                    return
                }
                
                this.remark.ArticleId = this.articleId;

                if (this.isUpdateRemark) {
                    modifyRemark(this.remark).then(response => {
                        if (response.Code == Code.SUCCESS) {
                            message.success(response.Message)
                            this.dialogVisible = false;
                        } else {
                            message.error(response.Message)
                        }
                    });
                } else {
                    createRemark(this.remark).then(response => {
                        if (response.Code == Code.SUCCESS) {
                            message.success(response.Message)
                            this.dialogVisible = false;
                        } else {
                            message.error(response.Message)
                        }
                    });
                }

                this.handleListRemarks();

                this.$router.push({
                    path: '/remark',
                    query: {
                        "articleid": this.articleId,
                    }
                });

            })
        },
    }
}
</script>

<style scoped>
a {
  /* color: #000;  */
  font-size: 20px;
}
.in-remark-content a {
    font-size: 15px;
    color: rgb(113, 137, 243);
}
.back {
    background: rgba(230, 244, 249, 0.7);
}
.content {
  position: relative;
  border-radius: 5px;
  height: 93.1%;
  margin-top: 38px;
  padding-top: 10px;
  background: rgba(230, 244, 249, 0.85);
  opacity: 0.98;
}
.remark-content {
    z-index: 2;
    background: rgba(58, 59, 59);
    color: #FFF;
    font-size: 16px;
    border-radius: 3px;
    width: 100%;
    text-align: left;
    height: 100%;
}
.in-remark-content {
    padding-top: 12px;
    padding-left: 30px;
    padding-right: 30px;
}
@media screen and (max-width: 700px) {
    .in-remark-content {
        font-size: 12px;
    }
    .in-remark-content a {
        font-size: 12px;
    }
}
</style>
