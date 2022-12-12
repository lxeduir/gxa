<template>
<div id="Personal">
    <div class="picture">
        <img src="@/assets/personalimag.jpeg">
        <a>个人服务</a>
    </div>
    <a class="text1">全心全意为您服务</a>
    <manybox></manybox>
</div>
</template>

<script>
import manybox from "@/components/Manybox";

export default {
    name: 'Personal',
    components: {
        manybox
    }

}
</script>

<style scoped>
    #Personal{
        margin-left: auto;
        margin-right: auto;
        width: 1540px;
        height: auto;
    }
    .picture>img{
        height: 300px;
        width: 1540px;
    }
    .picture>a{
        position: absolute;
        top:200px;
        left: 100px;
        font-size: 30px;
        color: white;
    }
    .text1{
        position: absolute;
        top:250px;
        left: 100px;
        font-size: 20px;
        color: white;
    }
</style>
