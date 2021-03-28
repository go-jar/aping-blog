import qs from 'qs';
import request from '../utils/request'

export function createTag(tag) {
    return request({
        url: process.env.WEB_API + '/tag/create',
        method: 'post',
        data: qs.stringify({
            "Action": "CreateTag",
            "TagName": tag.TagName,
            "TagIndex": tag.TagIndex,
        })
    })
}

export function deleteTag(tagId) {
    return request({
        url: process.env.WEB_API + '/tag/delete',
        method: 'post',
        data: qs.stringify({
            "Action": "DeleteTag",
            "TagId": tagId,
        })
    })
}

export function modifyTag(tag) {
    return request({
        url: process.env.WEB_API + '/tag/modify',
        method: 'post',
        data: qs.stringify({
            "Action": "ModifyTag",
            "TagId": tag.TagId, 
            "TagName": tag.TagName,
            "TagIndex": tag.TagIndex,
        })
    })
}

export function describeTags(tagId, offset, limit) {
    var params;
    if (tagId != null) {
        params = qs.stringify({
            "Action": "DescribeTags",
            "TagId": tagId,
        })
    } else {
        params = qs.stringify({
            "Action": "DescribeTags",
            "Offset": offset,
            "Limit": limit
        })
    }
    
    return request({
        url: process.env.WEB_API + '/tag/describe',
        method: 'post',
        data: params,
    })
}
