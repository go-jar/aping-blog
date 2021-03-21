<!-- 文章详情 -->
<template>
    <div>
       <wbc-nav></wbc-nav>
        <div class="container">
            <wbc-detail></wbc-detail>
        </div>
        <wbc-footer></wbc-footer>   
    </div>
</template>

<script>
    import Header from '@/components/Header'
    import Footer from '@/components/Footer'
    import ArticleDetail from '@/components/ArticleDetail'

    export default {
        name: 'ArticleDetail',
        data() {  // 选项 / 数据
            return {

            }
        },
        methods: {  // 事件处理器

        },
        components: {  // 定义组件
            'wbc-nav': Header,
            'wbc-detail': ArticleDetail,
            'wbc-footer': Footer
        },
        created() {  // 生命周期函数

        },
        mounted(){
        }
    }
</script>

<style>

</style>
