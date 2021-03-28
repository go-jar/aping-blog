import qs from 'qs';
import request from '../utils/request'

export function createRemark(remark) {
    var params;
    if (remark.InitRemarkId != null) {
        params = {
            "Action": "CreateRemark",
            "ArticleId": remark.ArticleId,
            "Nickname": remark.Nickname,
            "Content": remark.Content,
            "InitRemarkId": remark.InitRemarkId,
            "NicknameReplied": remark.NicknameReplied,
        }
    } else {
        params = {
            "Action": "CreateRemark",
            "ArticleId": remark.ArticleId,
            "Nickname": remark.Nickname,
            "Content": remark.Content,
        }
    }

    return request({
        url: process.env.WEB_API + '/remark/create',
        method: 'post',
        data: qs.stringify(params)
    })
}

export function deleteRemark(remarkId) {
    return request({
        url: process.env.WEB_API + '/remark/delete',
        method: 'post',
        data: qs.stringify({
            "Action": "DeleteRemark",
            "RemarkId": remarkId,
        })
    })
}

export function modifyRemark(remark) {
    return request({
        url: process.env.WEB_API + '/remark/modify',
        method: 'post',
        data: qs.stringify({
            "Action": "ModifyRemark",
            "RemarkId": remark.RemarkId, 
            "Content": remark.Content,
        })
    })
}

export function describeRemarks(articleId, offset, limit) {
    var params;
    if (articleId != null) {
        params = {
            "Action": "DescribeRemarks",
            "ArticleId": articleId,
        }
    } else {
        params = {
            "Action": "DescribeRemarks",
            "Offset": offset,
            "Limit": limit
        }
    }
    
    return request({
        url: process.env.WEB_API + '/remark/describe',
        method: 'post',
        data: qs.stringify(params),
    })
}
