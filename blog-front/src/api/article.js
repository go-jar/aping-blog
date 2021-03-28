import qs from 'qs';
import request from '../utils/request'

export function createArticle(article) {
    return request({
        url: process.env.WEB_API + '/article/create',
        method: 'post',
        data: qs.stringify({
            "Action": "CreateArticle",
            "Title": article.Title,
            "CategoryId": article.CategoryId,
            "Content": article.Content,
            "TagIds": article.TagIds,
        })
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
    return request({
        url: process.env.WEB_API + '/article/modify',
        method: 'post',
        data: qs.stringify({
            "Action": "ModifyArticle",
            "ArticleId": article.ArticleId, 
            "Title": article.Title,
            "CategoryId": article.CategoryId,
            "Content": article.Content,
            "TagIds": article.TagIds,
        })
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
