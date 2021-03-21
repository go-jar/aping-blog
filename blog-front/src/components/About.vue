<!-- 关于我 -->
<template>
    <div>
        <div class="container">
            <div class="about">
                <section>
                    <p class="">
                        <img :src="this.$store.state.aboutmeObj.image?this.$store.state.aboutmeObj.image:'static/img/logo.png'" alt="" onerror="this.onerror=null;this.src= 'static/img/maoto.png'" class="aboutmeImg">
                    </p>
                    <div class='content'>
                        <h2>
                            欢迎来到阿萍的博客
                        </h2>
                        <br/>
                        <h4>
                            本站建于 2020 年 12 月 07 日
                        </h4>
                        <h3>
                            要有最朴素的生活，与最遥远的梦想，即使明日天寒地冻，路远马亡 --- 七堇年
                        </h3>
                        <h3>
                            每个人心里都有下不完的雨，如果是我的话，我会选择用伞 ---《秋之回忆》
                        </h3>
                        <h3>
                            我们满怀希望，在有限的岁月里做着和梦想有关的事，岁月推移，总有一天会成为自己想要的样子
                        </h3>
                    </div>
                </section>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        data() { // 选项 / 数据
            return {

            }
        },
        methods: { // 事件处理器

        },
        components: { // 定义组件

        },
        created() { // 生命周期函数

        }
    }
</script>

<style>
.aboutmeImg{
    max-width:10%;
    /*height:300px;*/
    margin-top: 50px;
}
.about{
    white-space: normal;
    word-wrap: break-word;
    word-break: break-all;
    position: relative;
    background: rgba(230, 244, 249, 0.85);
    border-radius: 5px;
    font-size: 15px;
    height: 93.3%;
    margin-top: 38px;
    opacity: 0.98;
    text-align: center;
}
.content{
    background-position: center 50%;
    /* margin-left: 45%; */
}
.about h2{
    font-size: 50px;
	font-weight: 700;
    margin-top: 50px;
    font-family: Arial, sans-serif;
    display: inline-block;
}
.about h3{
    font-size: 20px;
	font-weight: 600;
    margin-top: 50px;
}
.about h4{
    font-size: 18px;
	font-weight: 500;
    margin-top: 50px;
}
.about p a{
    color:#01AAED!important;

}
.about p a:hover{
    color:#48456d!important;
}
.about p i{
    font-style: italic;
}
.about p b{
    font-weight: bold;
}
</style>
