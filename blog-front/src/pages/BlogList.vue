<!-- 文章详情 -->
<template>
    <div>
        <wbc-nav></wbc-nav>
        <div class="container">
            <wbc-bloglist></wbc-bloglist>
        </div>
        <wbc-footer></wbc-footer>   
    </div>
</template>

<script>
    import Header from '../components/Header.vue'
    import Footer from '../components/Footer.vue'
    import BlogList from '..//components/BlogList'

    export default {
        name: 'BlogList',
        data() { //选项 / 数据
            return {

            }
        },
        methods: { //事件处理器

        },
        components: { //定义组件
            'wbc-nav': Header,
            'wbc-rightlist': BlogListl,
            'wbc-footer': Footer
        },
        created() { //生命周期函数

        }
    }
</script>

<style>

</style>
