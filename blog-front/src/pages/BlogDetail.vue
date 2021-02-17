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
    import Header from '@/components/Header'
    import Footer from '@/components/Footer'
    import BlogDetail from '@/components/BlogDetail'

    export default {
        name: 'BlogDetail',
        data() { //选项 / 数据
            return {

            }
        },
        methods: { //事件处理器

        },
        components: { //定义组件
            'wbc-nav': Header,
            'wbc-detail': BlogDetail,
            'wbc-footer': Footer
        },
        created() { //生命周期函数

        },
        mounted(){
            var anchor = document.querySelector("#detail");
            // console.log(anchor,anchor.offsetTop);
            var top = anchor.offsetTop-60;
            document.body.scrollTop = top;
             // Firefox
             document.documentElement.scrollTop = top;
             // Safari
             window.pageYOffset = top;
        }
    }
</script>

<style>

</style>
