import qs from 'qs';
import request from '../utils/request'
import {getToken} from '@/utils/auth'

export function createArticle(article) {
    return request({
        url: process.env.WEB_API + '/article/create',
        method: 'post',
        data: qs.stringify({
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
            "Id": articleId,
        })
    })
}

export function updateArticle(article) {
    return request({
        url: process.env.WEB_API + '/article/update',
        method: 'post',
        data: qs.stringify({
            "Id": article.Id, 
            "Title": article.Title,
            "CategoryId": article.CategoryId,
            "Content": article.Content,
            "TagIds": article.TagIds,
        })
    })
}

export function describeArticles(articleId, categoryId, tagId, offset, limit) {
    var params = new Array();

    if (articleId != null) {
        params["Id"] = articleId;
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
    
    return request({
        url: process.env.WEB_API + '/article/describe',
        method: 'post',
        data: qs.stringify(params),
    })
}
