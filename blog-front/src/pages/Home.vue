<!-- 首页 -->
<template>
    <div>
        <wbc-nav></wbc-nav>
        <wbc-home></wbc-home>
        <wbc-footer></wbc-footer>   
    </div>
</template>

<script>
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'
    import Home from '@/components/Home.vue'

    export default {
        name: 'Home',
        data() { // 选项 / 数据
            return {

            }
        },
        methods: { // 事件处理器

        },
        components: { // 定义组件
            'wbc-nav': Header,
            'wbc-footer': Footer,
            'wbc-home': Home,
        },
        created() { // 生命周期函数

        }
    }
</script>

<style>

</style>
