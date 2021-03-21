<template>
<div>
    <div id="vditor" class="vditor"/>
</div>
</template>

<script>
import Vditor from 'vditor'
import "vditor/dist/index.css"
import "vditor/dist/index.min.js"

export default {
  name: 'ViewMarkdown',
  props: ["height", "content"],
  data() {
    return {
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
          enable: false,
        },
        mode: 'wysiwyg',
        height: this.height,  // 编辑器总高度
        typewriterMode: true,
        toolbarConfig: {
          hide: true,  // 隐藏工具栏
        },
        theme: 'dark',  // 编辑器主题
        preview: {
          delay: 100,
          theme: {
            current: 'dark'  // 内容主题
          },
          hljs: {
            enable: true,  // 代码高亮
            style: 'vim',  
            lineNumber: true
          },
          actions: ["desktop", "tablet", "mobile"],
        },
        outline: {  // 显示大纲
          enable: true
        },
        toolbar: [
          'fullscreen',
          'both',
          'preview',
          'info',
          'help',
        ],
        value: this.content == null? "": this.content,
        after: () => {
          var evt = document.createEvent('Event');
          evt.initEvent('click', true, true);
          this.vditor.vditor.toolbar.elements.preview.firstElementChild.dispatchEvent(evt);
        },
      }
      this.vditor = new Vditor('vditor', options)
      return this.vditor
    },
  }
}
</script>

<style scoped>

</style>
