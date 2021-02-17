<template>
<div>
    <div id="vditor" class="vditor"/>
</div>
</template>

<script>
import Vditor from 'vditor'
import "vditor/dist/index.css"
import {GetToken} from '@/utils/auth'
import {htmlToMarkdownFile} from '@/utils/common'

export default {
  name: 'MarkdownEditor',
  props: ["height"],
  data() {
    return {
      isLoading: true,
      isMobile: window.innerWidth <= 960,
      vditor: null,
    }
  },
  created() {
  },
  components: {
  },
  mounted() {
    this.initVditor()
    this.$nextTick(() => {
      this.isLoading = false
    })
  },
  methods: {
    initVditor() {
      const that = this
      const options = {
        cache: {
          enable: false,
        },
        mode: 'ir',
        height: this.height,  // 编辑器总高度
        theme: 'dark',  // 编辑器主题
        preview: {
          delay: 100,
          show: !this.isMobile,
          theme: {
            current: 'dark'  // 内容主题
          },
          hljs: {
            enable: true,  // 代码高亮
            style: 'vim',  
            lineNumber: true
          },
          markdown: {
            toc: true  // 展示目录
          }
        },
        toolbarConfig: {
            pin: false,  // 是否固定工具栏
        },
        upload: {
          max: 5 * 1024 * 1024,
          // linkToImgUrl: 'https://sm.ms/api/upload',
          handler(file) {
            let formData = new FormData()
            for (let i in file) {
              formData.append('smfile', file[i])
            }
            let request = new XMLHttpRequest()
            // 图片上传路径
            request.open('POST', process.env.PICTURE_API + '/ckeditor/imgUpload?token=' + GetToken())
            request.onload = that.onloadCallback
            request.send(formData)
          }
        }
      }
      this.vditor = new Vditor('vditor', options)
      return this.vditor
    },
    onloadCallback(oEvent) {
      const currentTarget = oEvent.currentTarget
      console.log("返回的结果", currentTarget)
      if (currentTarget.status !== 200) {
        return this.$message({
          type: 'error',
          message: currentTarget.status + ' ' + currentTarget.statusText
        })
      }
      let resp = JSON.parse(currentTarget.response)
      let imgMdStr = ''
      if (resp.uploaded !== 1) {
        return this.$message({
          type: 'error',
          message: resp.error.message
        })
      }
      if (resp.uploaded === 1) {
        imgMdStr = `![${resp.fileName}](${resp.url})`
      }
      this.vditor.insertValue(imgMdStr)
    },
    // 获取data
    getData: function() {
      return this.vditor.getHTML();
    },
    setData: function(data) {
      var that = this;
      this.$nextTick(() => {
        // DOM现在更新了
        let vditor = that.initVditor()
        let markdownText = commonUtil.htmlToMarkdown(data)
        localStorage.setItem('vditor', markdownText)
      });
    },
    initData: function () {
      var that = this
      this.$nextTick(() => {
        if(that.vditor.vditor.lute) {
          that.vditor.setValue("")
        }
      });
    }
  }
}
</script>

<style>
.vditor-panel {
  line-height: 0px;
}
.vditor {
  width: 100%;
  /*height: calc(100vh - 100px);*/
  top: 20px;
  /*margin: 20px auto;*/
  text-align: left;
}
.vditor-reset {
  font-size: 14px;
}
.vditor-textarea {
  font-size: 14px;
  height: 100% !important;
}
</style>
