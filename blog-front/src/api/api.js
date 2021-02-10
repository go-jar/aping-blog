import Vue from 'vue'
import axios from 'axios'

// 全局设置
axios.defaults.baseURL = 'https://aping-dev.com/back';
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
let blogUrl = "https://aping-dev.com/back"

// 文章分类查询
const ArticleClassData = (callback) => {
    if(sessionStorage.getItem('classList')){
        var data = JSON.parse(sessionStorage.getItem('classList'));
        callback && callback(data)
    }else{
        let url = blogUrl + 'article/ArticleClassData';
        axios.get(url).then(num => {
            // console.log(num);
            if(num.data.code==1001){
                sessionStorage.setItem('classList',JSON.stringify(num.data.data));
                callback && callback(num.data.data)
            }else{
                alert("查询失败")
            }
        })
    }
}

// 查询文章列表
const ShowArticleAll = (artId,cateId,articleName,level,callback) =>{
    if(level == 1){
        var url = blogUrl + 'nav/ActiveClassAllData?art_id='+artId+'&cate_id='+cateId+'&article_name='+articleName;
    }else{
        var url = blogUrl + 'article/ShowArticleAll?art_id='+artId+'&cate_id='+cateId+'&article_name='+articleName;
    }
    axios.get(url).then(num => {
            callback && callback(num.data);
    })
}

// 查询文章详情
const getArticleInfo = (artId,userId,callback) =>{
    let url = blogUrl + 'article/getArticleInfo?art_id='+artId+'&user_id='+userId;
    axios.get(url).then(num => {
        if(num.data.code==1001){
            callback && callback(num.data.data);
        }else{
            alert("查询失败");
        }
    })
}

// 查询浏览量最多的10篇文章数据
const ShowBrowseCount = (callback) =>{
    let url = blogUrl + 'article/ShowBrowseCount';
    axios.get(url).then(num => {
        if(num.data.code==1001){
            callback && callback(num.data.data);
        }else if(num.data.code==1005){
            return;
        }else{
            alert("查询失败");
        }
    })
}

// 查询文章评论量最大的10篇文章
const ShowArtCommentCount = (callback) =>{
    let url = blogUrl + 'article/ShowArtCommentCount';
    axios.get(url).then(num => {
        if(num.data.code==1001){
            callback && callback(num.data.data);
        }else if(num.data.code==1005){
            return;
        }else{
            alert("查询失败");
        }
    })
}
//查询文章评论数据
const ArticleComment = (artId,commentId,callback) =>{
    let url = blogUrl + 'comment/ArticleComment?art_id='+artId+'&comment_id='+commentId;
    axios.get(url).then(num => {
            callback && callback(num.data);
    })
}

//查询其他评论数据
const OtherComment = (leaveId,commentId,callback) =>{//分类类型ID（1：赞赏 2：友情链接 3：留言板 4：关于我）
    let url = blogUrl + 'comment/OtherComment?leave_id='+leaveId+'&comment_id='+commentId;
    axios.get(url).then(num => {
        callback && callback(num.data);
    })
}

//文章评论
const setArticleComment = (content,user_id,article_id,leave_pid,pid,callback) =>{
    let url = blogUrl + 'comment/setArticleComment?content='+content+'&user_id='+user_id+'&article_id='+article_id+'&leave_pid='+leave_pid+'&pid='+pid;
    axios.get(url).then(num => {
            callback && callback(num.data);
    })
}

//其他评论
const setOuthComment = (content,user_id,article_id,leave_id,leave_pid,pid,callback) =>{
    let url = blogUrl + 'comment/setOuthComment?content='+content+'&user_id='+user_id+'&article_id='+article_id+'&leave_id='+leave_id+'&leave_pid='+leave_pid+'&pid='+pid;
    axios.get(url).then(num => {
            callback && callback(num.data);
    })
}

//查询网址点赞总数
const showLikeData = (callback) =>{
    let url = blogUrl + 'outh/showLikeData';
    axios.get(url).then(num => {
        if(num.data.code==1001){
            // console.log(num.data,parseInt(num.data));
            callback && callback(num.data.data);
        }else{
            alert("查询失败");
        }
    })
}

//点赞功能修改
const GetLike = (like_num,callback) =>{
    let url = blogUrl + 'outh/GetLike?like_num='+like_num;
    axios.get(url).then(num => {
        if(num.data.code==1001){
            callback && callback(num.data.msg);
        }else{
            alert("点赞失败");
        }
    })
}

//查询赞赏数据
const AdmireData = (callback) => {
    let url = blogUrl + 'outh/AdmireData';
    axios.get(url).then(num => {
        if(num.data.code==1001){
            callback && callback(num.data);
        }else{
            alert("查询失败");
        }
    })
}

//查询用户喜欢列表,查询用户收藏列表
const getLikeCollectList = (userId,artId,articleName,islike,callback)=>{
    var url = '';
    if(islike==1){
         url = blogUrl + 'article/getLikeList?user_id='+userId+'&art_id='+artId+'&article_name='+articleName;
    }else{
         url = blogUrl + 'article/getCollectList?user_id='+userId+'&art_id='+artId+'&article_name='+articleName;
    }
    axios.get(url).then(num => {
            callback && callback(num.data);
    })
}

//初始化时间
const initDate = (oldDate,full) => {
    var odate = new Date(oldDate);
    var year =  odate.getFullYear();
    var month = odate.getMonth()<9? '0' + (odate.getMonth()+1) : odate.getMonth()+1;
    var date = odate.getDate()<10? '0'+odate.getDate() : odate.getDate();
    if(full=='all'){
        var t = oldDate.split(" ")[0];
        // console.log(oldDate,t.split('-')[0],t.split('-')[1],t.split('-')[2]);
        return t.split('-')[0]+'年'+t.split('-')[1]+'月'+t.split('-')[2]+'日';
    }else if(full=='year'){
        return year
    }else if(full== 'month'){
        return odate.getMonth()+1
    }else if(full == 'date'){
        return date
    }else if(full== 'newDate'){
        return year+'年'+month+'月'+date+'日';
    }
}

//获取主题信息
const changeTheme = (callback) => {
    if(sessionStorage.getItem('changeThemeObj')){
        var data = JSON.parse(sessionStorage.getItem('changeThemeObj'));
        callback && callback(data)
    }else{
        let url = blogUrl + 'outh/ThemeMy';
        axios.get(url).then(num => {
            if(num.data.code==1001){
                sessionStorage.setItem('changeThemeObj',JSON.stringify(num.data.data))
                callback && callback(num.data.data);
            }else{
                alert("查询失败");
            }
        })
    }
}

export {
    ArticleClassData, // 分类
    ShowArticleAll, // 查询文章列表
    getArticleInfo, // 文章详情
    ShowBrowseCount, // 流量量做多的文章
    ShowArtCommentCount, // 评论最多的文章
    ArticleComment, // 文章评论列表
    OtherComment, // 其他评论列表
    setArticleComment, // 设置文章评论
    setOuthComment, // 设置其他评论
    showLikeData, // do you like me
    GetLike, // 设置 do you like me
    AdmireData, // 赞赏数据
    getLikeCollectList, // 用户收藏喜欢列表
    initDate, // 设置时间
    changeTheme, // 获取主题信息
}
