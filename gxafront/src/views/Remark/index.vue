<template>
<div id="Remark">
    <div class="picture">
        <a class="text0">统一好差评</a>
        <a class="text1">全面及时准确了解群众对政务服务的感受和诉求，</a>
        <a class="text2">提升政府工作效能，优化营商环境，建设人民满意的服务型政府</a>
    </div>
    <div class="toremark" @click="toMyremark">评价我的办件</div>
    <div class="detail">
        <div>
            办件数
            <a>9999</a>
        </div>
        <div>
            好评数
            <a>9999</a>
        </div>
        <div>
            满意率
            <a>100.00%</a>
        </div>
    </div>
</div>
</template>

<script>
export default {
    data(){
        return {
        };
    },
    methods: {
        toMyremark() {
            this.$router.push({name:'myremark'})
        }

  }
}
</script>

<style scoped>
    .picture{
        height: 300px;
        width: 1540px;
        /* background-color: #6db6ff; */
        background-image: url('@/assets/zhuhai.jpg');
    }
    .text0{
        position: absolute;
        top:150px;
        left: 100px;
        font-size: 30px;
        color: white;
    }
    .text1{
        position: absolute;
        top:200px;
        left: 100px;
        font-size: 20px;
        color: white;
    }
    .text2{
        position: absolute;
        top:230px;
        left: 100px;
        font-size: 20px;
        color: white;
    }
    .toremark{
        position: relative;
        top: -130px;
        left: 100px;
        background-color: #1482F1;
        color: white;
        border-radius: 5px;
        width: 120px;
        height: 40px;
        line-height: 40px;
        cursor: pointer;
    }
    .detail{
        background-color: white;
        position: relative;
        top: -120px;
        left: 170px;
        width: 1200px;
        height: 250px;
        border-radius: 10px;
        border: 1px solid #b3d8fc;
        box-shadow: 1px 1px 1px 1px #b3d8fc;
    }
    .detail div{
        display: inline;
        margin: 100px;
        position: relative;
        top: 50px;
        font-size: 30px;
    }
    .detail a{
        font-size: 26px;
        position: relative;
        top: 50px;
        left: -90px;
    }
</style>
