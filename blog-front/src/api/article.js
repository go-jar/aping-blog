import qs from 'qs';
import request from '../utils/request'

export function createArticle(article) {
    var params = new Array();
    params["Action"] = "CreateArticle"
    params["Title"] = article.Title
    params["CategoryId"] = article.CategoryId
    params["Content"] = article.Content
    
    if (article.TagIds != null && article.TagIds != "") {
        params["TagIds"] = article.TagIds
    }

    return request({
        url: process.env.WEB_API + '/article/create',
        method: 'post',
        data: qs.stringify(params)
    })
}

export function deleteArticle(articleId) {
    return request({
        url: process.env.WEB_API + '/article/delete',
        method: 'post',
        data: qs.stringify({
            "Action": "DeleteArticle",
            "ArticleId": articleId,
        })
    })
}

export function modifyArticle(article) {
    var params = new Array();
    params["Action"] = "ModifyArticle"
    params["ArticleId"] = article.ArticleId
    params["Title"] = article.Title
    params["CategoryId"] = article.CategoryId
    params["Content"] = article.Content
    
    if (article.TagIds != null && article.TagIds != "") {
        params["TagIds"] = article.TagIds
    }
    
    return request({
        url: process.env.WEB_API + '/article/modify',
        method: 'post',
        data: qs.stringify(params)
    })
}

export function describeArticles(articleId, categoryId, tagId, keyword, offset, limit) {
    var params = new Array();
    params["Action"] = "DescribeArticles";

    if (articleId != null) {
        params["ArticleId"] = articleId;
    } 
    
    if (categoryId != null) {
        params["CategoryId"] = categoryId;
    }

    if (tagId != null) {
        params["TagId"] = tagId;
    }

    if (offset != null) {
        params["Offset"] = offset;
    }

    if (limit != null) {
        params["Limit"] = limit;
    }
    
    if (keyword != null) {
        params["Keyword"] = keyword
    }

    return request({
        url: process.env.WEB_API + '/article/describe',
        method: 'post',
        data: qs.stringify(params),
    })
}
