import {Message} from 'element-ui'

/**
 * 将Html转成Markdown文件
 * @param title：标题
 * @param text：正文
*/
export function htmlToMarkdownFile (title, text) {
    title = title || "默认标题"

    let turndownService = new TurndownService()

    let markdown = turndownService.turndown(text)

    //创建一个 blob 对象, file 的一种
    let blob = new Blob([markdown])

    let link = document.createElement('a')

    link.href = window.URL.createObjectURL(blob)

    //配置下载的文件名
    link.download = title + '.md'

    link.click()
}

export const message = {
    success: function(message) {
      Message({
        showClose: true,
        message: message || '成功',
        type: 'success'
      })
    },
    warning: function(message) {
      Message({
        showClose: true,
        message: message || '警告',
        type: 'warning'
      })
    },
    info: function(message) {
      Message({
        showClose: true,
        message: message || '提示'
      })
    },
    error: function(message) {
      Message({
        showClose: true,
        message: message || '异常',
        type: 'error'
      })
    }
}