import qs from 'qs';
import request from '../utils/request'

export function createCategory(category) {
    return request({
        url: process.env.WEB_API + '/category/create',
        method: 'post',
        data: qs.stringify({
            "Action": "CreateCategory",
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
            "Action": "DeleteCategory",
            "CategoryId": categoryId,
        })
    })
}

export function modifyCategory(category) {
    return request({
        url: process.env.WEB_API + '/category/modify',
        method: 'post',
        data: qs.stringify({
            "Action": "ModifyCategory",
            "CategoryId": category.CategoryId, 
            "CategoryName": category.CategoryName,
            "CategoryIndex": category.CategoryIndex,
        })
    })
}

export function describeCategories(categoryId, offset, limit) {
    var params;
    if (categoryId != null) {
        params = qs.stringify({
            "Action": "DescribeCategories",
            "CategoryId": categoryId,
        })
    } else {
        params = qs.stringify({
            "Action": "DescribeCategories",
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
