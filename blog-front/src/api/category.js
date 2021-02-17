import qs from 'qs';
import request from '../utils/request'

export function createCategory(category) {
    return request({
        url: process.env.WEB_API + '/category/create',
        method: 'post',
        data: qs.stringify({
            "CategoryName": category.CategoryName,
            "CategoryIndex": category.CategoryIndex,
        })
    })
}

export function deleteCategory(categoryId) {
    return request({
        url: process.env.WEB_API + '/category/delete',
        method: 'post',
        data: qs.stringify({
            "Id": categoryId,
        })
    })
}

export function updateCategory(category) {
    return request({
        url: process.env.WEB_API + '/category/update',
        method: 'post',
        data: qs.stringify({
            "Id": category.Id, 
            "CategoryName": category.CategoryName,
            "CategoryIndex": category.CategoryIndex,
        })
    })
}

export function describeCategories(categoryId, offset, limit) {
    var params;
    if (categoryId != null) {
        params = qs.stringify({
            "Id": categoryId,
        })
    } else {
        params = qs.stringify({
            "Offset": offset,
            "Limit": limit
        })
    }
    
    return request({
        url: process.env.WEB_API + '/category/describe',
        method: 'post',
        data: params,
    })
}
