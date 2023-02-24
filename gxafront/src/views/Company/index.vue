<template>
    <div id="company">
        <div class="picture">
            <img src="@/assets/biulding.jpg">
            <a>法人服务</a>
        </div>
        <a class="text1">全心全意为您服务</a>
        <companybox></companybox>
    </div>
</template>

<script>
import companybox from "@/components/Companybox";

export default {
    components: {
        companybox
    }

}
</script>

<style scoped>
#company {
    margin-left: auto;
    margin-right: auto;
    width: 1540px;
    height: auto;
}

.picture>img {
    height: 300px;
    width: 1540px;
}

.picture>a {
    position: absolute;
    top: 200px;
    left: 100px;
    font-size: 30px;
    color: white;
}

.text1 {
    position: absolute;
    top: 250px;
    left: 100px;
    font-size: 20px;
    color: white;
}
</style>
