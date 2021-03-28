<!-- 底部公用 -->
<template>
	<div class='footBack'>
		<div v-if="this.$store.state.themeObj.user_start!=0" class="footBack footBackQian">
			<div class="fcontainer">
				<p>
					博客已运行：<span v-html='longTime'>{{longTime}}</span><span class="timeJump"></span>
				</p>
				<p>
					© 2020 阿萍. 由 <a href="https://cn.vuejs.org/" target="_blank">Vue</a> 强力驱动. 京ICP备20000257号.
				</p>
			</div>
		</div>
		<div v-else class="footBack footBackHui">
			<p>
				博客已运行：<span v-html='longTime'>{{longTime}}</span><span class="timeJump"></span>
			</p>
			<p>
				© 2021 阿萍. 由 <a href="https://cn.vuejs.org/" target="_blank">Vue</a> 强力驱动. 京ICP备20000257号.
			</p>
		</div>
	</div>
</template>

<script>
    export default {
        name: 'Footer',
        data() { // 选项 / 数据
            return {
                longTime: '',
            }
        },
        methods: { // 事件处理器
            runTime: function(){ // 运行倒计时
                var that = this;
                var oldTime =new Date('2021/03/26 00:00:00');  // 博客上线时间
                var timer = setInterval(function(){
                    var nowTime = new Date();
                    var longTime = nowTime - oldTime;
                    var days = parseInt(longTime / 1000 / 60 / 60 / 24 , 10);
                    var hours = parseInt(longTime / 1000 / 60 / 60 % 24 , 10);
                    var minutes = parseInt(longTime / 1000 / 60 % 60, 10);
                    var seconds = parseInt(longTime / 1000 % 60, 10);
                    that.longTime = days + "天" + hours + "小时" + minutes + "分" + seconds + "秒";
                }, 1000)
            }
        },
        components: { // 定义组件

        },
        created() { // 生命周期函数
            // 替换底部图片
            var that = this;
            that.runTime();
        }
    }
</script>

<style>
.footBack{
    color: #888;
    /* margin-top: -10px; */
    font-size: 12px;
    line-height: 1.5;
    text-align: center;
    width: 100%;
    min-height: 50px;
    margin-top: -25px;
    /*position: relative;*/
    position: absolute;
    /*min-height: 368px;*/
}
.footBack .fcontainer{
    width: 100%;
    background: #232323;
    /*top:294px;*/
    /*left:0;*/
    /*position: absolute;*/
    padding: 10px 10px 10px 10px;
    box-sizing: border-box;
    /*z-index: 1;*/
    width: 100%!important;
}
.footBack .footer-img{
    height: 368px;
    z-index: 1;
    position: relative;
    width: 100%;
    bottom: -74px;
    pointer-events: none;
}
.footBackHui{
    /*min-height: 50px;*/
    margin-top: -50px;
    z-index: -1;
}
.footBackHui .footer-img img{
    width:100%;
    height:auto;
    margin:0;

}
.footBackHui .fcontainer{
    /*z-index: 1;*/
    /*position: absolute;*/
    /*bottom:0;*/

}
.footBack p{
    margin:5px 0;
    z-index: 3;
}
.footBack a{
    color:#795548;
    z-index: 3;
}
.footBack a:hover{
    color:#000;
}

.timeJump{
    display:inline-block;
	margin:0 5px;
}

@media (max-width: 500px) {     
	.footBack .fcontainer{
        width: 100%;
        background: #232323;
        padding: 15px 10px 10px 10px;
        box-sizing: border-box;
        width: 100%!important;
    }
}
</style>
