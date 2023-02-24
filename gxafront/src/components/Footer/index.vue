<template>
  <div class="footer">
        <p>主办：lx</p>
        <p>承办：lx大数据中心</p>
        <p>支持：hzq有限公司</p>
        <p>网址标识码：87654321</p>
        <p>联系我们：123123</p>
</div>
</template>
<script>
export default {
    data(){
        return {
        };
    },
methods: {},
};
</script>

<style scoped lang='less'>
    .footer{
        position: relative;
        top: 50px;
        width: 100%;
        height: 75px;
        line-height: 50px;
        display: flex;
        background-color: #eaebeb;
        justify-content: center;
    }
    .footer p{
        font-size: 15px;
        margin-right: 25px;
    }
</style>