import qs from 'qs';
import request from '../utils/request'

export function createCategory(categoryName) {
    return request({
        url: process.env.WEB_API + '/category/create',
        method: 'post',
        data: qs.stringify({
            "CategoryName": categoryName,
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

export function updateCategory(categoryId, newCategoryName) {
    return request({
        url: process.env.WEB_API + '/category/update',
        method: 'post',
        data: qs.stringify({
            "Id": categoryId, 
            "CategoryName": newCategoryName
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
