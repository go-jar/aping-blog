<!-- 首页 -->
<template>
<div>
    <wbc-nav></wbc-nav>
    <div class="container">
        <wbc-content></wbc-content>
    </div>
    <wbc-footer></wbc-footer>   
</div>
</template>

<script>
    import Header from '@/components/Header'
    import Footer from '@/components/Footer'
    import ManageTag from '@/components/ManageTag'

    export default {
        data() { // 选项 / 数据
            return {

            }
        },
        methods: { // 事件处理器

        },
        components: { // 定义组件
            'wbc-nav': Header,
            'wbc-content': ManageTag,
            'wbc-footer': Footer
        },
        created() { // 生命周期函数

        }
    }
</script>

<style>

</style>
