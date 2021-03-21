<template>
<div>
    <div id="vditor" class="vditor"/>
</div>
</template>

<script>
import {Code} from '@/const/code.js'
import Vditor from 'vditor'
import "vditor/dist/index.css"
import "vditor/dist/index.min.js"
import {htmlToMarkdownFile} from '@/utils/common'

export default {
  name: 'MarkdownEditor',
  props: ["height", "content"],
  data() {
    return {
      isMobile: window.innerWidth <= 960,
      vditor: null,
    }
  },
  created() {
  },
  components: {
  },
  mounted() {
    this.initVditor();
  },
  methods: {
    initVditor() {
      const that = this
      const options = {
        cache: {
          enable: true,
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
          outline: {  // 显示大纲
            enable: true
          },
          // markdown: {
          //   toc: true  // 展示目录
          // }
        },
        toolbarConfig: {
            pin: false,  // 是否固定工具栏
        },
        after: () => {
          if (this.content) {
            this.vditor.setValue(this.content)
          }
        },
        upload: {
          url: process.env.WEB_API + '/file/upload',
          max: 5 * 1024 * 1024,
          linkToImgUrl: process.env.WEB_API + '/file/upload',
          handler(files) {
            let formData = new FormData()
            for (let i = 0; i < files.length; i++) {
              formData.append('file', files[i])
            }
            let request = new XMLHttpRequest()
            // 图片上传路径
            request.open('POST', process.env.WEB_API + '/file/upload')
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
      
      if (resp.Code === Code.SUCCESS) {
        imgMdStr = `![${'img'}](${resp.Data.ImgUrl})`
      } else {
        return this.$message({
          type: 'error',
          message: resp.Msg
        })
      }
      this.vditor.insertValue(imgMdStr)
    },
    // 获取data
    getData: function() {
      // return this.vditor.getHTML();
      return this.vditor.getValue();
    },
    setData: function(data) {
      var that = this;
      this.$nextTick(() => {
        // DOM 现在更新了
        let markdownText = htmlToMarkdownFile(data)
        that.vditor.setValue(markdownText)
      });
    },
    initData: function () {
      var that = this
      this.$nextTick(() => {
        if(that.vditor.lute) {
          that.vditor.setValue("")
        }
      });
    },
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
