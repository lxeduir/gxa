<template>
  <div class="box">
        <ul class="daohang">
            <span class="geren"></span>
            <li>出生收养</li>
            <li>学习教育</li>
            <li>就业创业</li>
            <li class="last">职业资格</li>
        </ul>
        <ul class="content">
            <li>666666
              <img src="@/assets/chusheng.png" alt="">
            </li>
            <li>666666
              <img src="@/assets/edu.png" alt="">
            </li>
            <li>666666
              <img src="@/assets/chuangye.png" alt="" height="150">
            </li>
            <li class="last">666666
              <img src="@/assets/zizhi.png" alt="">
            </li>
        </ul>
        <ul class="daohang1">
            <span class="qiye"></span>
            <li>工作求职</li>
            <li>社保就医</li>
            <li>退休养老</li>
            <li class="last">婚姻服务</li>
        </ul>
        <ul class="content">
            <li>666666
              <img src="@/assets/job.png" alt="" height="150">
            </li>
            <li>666666
              <img src="@/assets/shebao.png" alt="">
            </li>
            <li>666666
              <img src="@/assets/yaolao.png" alt="">
            </li>
            <li class="last">666666
              <img src="@/assets/hunyin.png" alt="">
            </li>
        </ul>
        <div class="more">
          查看更多<img src="@/assets/查看更多.png" alt="" width="15px" height="15px">
        </div>
    </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {},
};
</script>

<style scoped>
    li{
        list-style: none;
    }
    a{
        text-decoration: none;
    }
 .box{
        width:1230px;
        margin: 0 auto;
        height: 600px; 
        position: relative;
        /* background-color: gray; */
    }
    .box .mainDaohang ul li{
        width: 300px;
        float: left;
        height: 50px;
        margin-right: 10px;
        margin-bottom: 10px;
    }
    .box .daohang li{
        background: radial-gradient(#ebebeb,white);
        width: 300px;
        float: left;
        height: 50px;
        margin-right: 10px;
        margin-bottom: 10px;
        font-size: 20px;
    }
    .box .daohang1 li{
      background: radial-gradient(#ebebeb,white);
        width: 300px;
        float: left;
        height: 50px;
        margin-right: 10px;
        margin-bottom: 10px;
        margin-top: 10px;
        font-size: 20px;
        
    }
    .box .daohang .last{
        margin-right: 0;
    }
    .box .daohang1 .last{
        margin-right: 0;
    }
    .box .content li{
        float: left;
        width: 300px;
        height: 250px; 
        background:linear-gradient(to top,#d4e8f4,white);
        margin-right: 10px;
    }
    .box .content .last{
        margin-right: 0;
    }
    .box .more{
      width: 100px;
      position: absolute;
      height: 20px;
      border: 1px solid rgb(169, 165, 165);
      top:650px;
      left: 1100px;
    }
    .box .more:hover{
      background-color: rgb(212, 220, 220);
    }

</style>