import qs from 'qs';
import request from '../utils/request'

export function createTag(tag) {
    return request({
        url: process.env.WEB_API + '/tag/create',
        method: 'post',
        data: qs.stringify({
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
            "Id": tagId,
        })
    })
}

export function updateTag(tag) {
    return request({
        url: process.env.WEB_API + '/tag/update',
        method: 'post',
        data: qs.stringify({
            "Id": tag.Id, 
            "TagName": tag.TagName,
            "TagIndex": tag.TagIndex,
        })
    })
}

export function describeTags(tagId, offset, limit) {
    var params;
    if (tagId != null) {
        params = qs.stringify({
            "Id": tagId,
        })
    } else {
        params = qs.stringify({
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
