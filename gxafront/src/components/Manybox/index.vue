<template>
<div>
    <div class="list">
        <div class="list1">
            <div>户籍办理</div>
            <div>就业创业</div>
            <div>职业资格</div>
            <div>入伍服役</div>
            <div>证件办理</div>
        </div>
        <div class="list2">
            <div>准营准办</div>
            <div>婚姻登记</div>
            <div>医疗卫生</div>
            <div>离职退休</div>
            <div>社会保障</div>       
        </div>
        <div class="list3">
            <div>其他</div>
        </div>
    </div>
    <div class="title">职业资格</div>
    <div class="bod"></div>
    <div class="check">
        <el-checkbox-group v-model="checkList">
        <el-checkbox label="全部"></el-checkbox>
        <el-checkbox label="支持网上申请"></el-checkbox>
        <el-checkbox label="零跑腿"></el-checkbox>
        <!-- <el-checkbox label="禁用" disabled></el-checkbox> -->
        <!-- <el-checkbox label="选中且禁用" disabled></el-checkbox> -->
        </el-checkbox-group>
    </div>
    <div class="search">
        <el-select
        v-model="value"
        multiple
        filterable
        remote
        reserve-keyword
        placeholder="请输入关键词">
        <!-- <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value">
        </el-option> -->
        </el-select>
    </div>
    <div class="detail">
        <div class="det1">
            <div class="name1">领取执业药师（药学，中药学）资格证书</div>
            <div class="select1">
                <div>办理指南</div>
                <div>订阅</div>
                <div>申请</div>
            </div>
        </div>
    </div>
</div>
</template>

<script>
export default {
    data(){
        return {
            checkList: ['选中且禁用','复选框 A']
        };
    },
methods: {},
};
</script>

<style scoped> 
    .list{
        position: relative;
        top: 20px;
    }
    .list1{
        position: relative;
        left: 160px;
        width: 1200px;
        height:60px;
    }
    .list1>div{
        position: relative;
        display: inline;
        margin: 20px;
        top: 20px;
        padding-left: 50px;
        padding-right: 50px;
        padding-top: 20px;
        padding-bottom: 20px;
        border: 1px solid #c7c7c7;
        background-color: white;
        box-shadow: 3px 3px 3px #c7c7c7;
    }
    .list2{
        position: relative;
        left: 160px;
        top: 10px;
        width: 1200px;
        height:60px;
    }
    .list2>div{
        position: relative;
        display: inline;
        margin: 20px;
        top: 20px;
        padding-left: 50px;
        padding-right: 50px;
        padding-top: 20px;
        padding-bottom: 20px;
        border: 1px solid #c7c7c7;
        background-color: white;
        box-shadow: 3px 3px 3px #c7c7c7;
    }
    .list3>div{
        position: relative;
        border: 1px solid #c7c7c7;
        left: 266px;
        top: 20px;
        padding-top: 20px;
        padding-bottom: 20px;
        width: 163px;
        background-color: white;
        box-shadow: 3px 3px 3px #c7c7c7;
    }
    .title{
        position: relative;
        width: 100px;
        height: 30px;
        top: 60px;
        left: 266px;
        border-bottom: 2px solid #1482F1;
        font-size: 20px;
        color: #1482F1;
    }
    .bod{
        border-left: 2px solid #1482F1;
        border-top: 2px solid #1482F1;
        width: 890px;
        position: relative;
        left: 366px;
        height: 15px;
        top: 43px;
    }
    .check{
        position: relative;
        top: 43px;
        left: 600px;
        width:500px;
    }
    .search{
        position: relative;
        top: 13px;
        left: 1050px;
        width: 200px;
    }
    .detail{
        
        position: relative;
        top: 20px;
    }
    .det1{
        border-left: 2px solid #1482F1;
        width: 989px;
        position: relative;
        left: 266px;
        display: flex;
    }
    .name1{
        height: 30px;
        width: 300px;
        line-height: 30px;
    }
    .select1{
        display: flex;
        float: right;
        transform: translateX(74%)
    }
    .select1 div{
        border: 1px solid blue;
        border-radius:5px; 
        margin-left: 30px;
        width: 100px;
        height: 25px;
        position: relative;
        top: 2px;
    }
</style>