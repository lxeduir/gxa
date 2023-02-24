<template>
    <div class="sidebar">
      <ul>
        <!-- <li>
          <a>
            <span class="inconfont icon-shouqi"></span>
            <p>收起</p>
          </a>
        </li> -->
        <li>
          <a @click="todoc">
          <span class="iconfont icon-caozuozhinan"></span>
          <p>操作指南</p>
        </a>
        </li>
        <li>
          <a @click="toid">
            <span class="iconfont icon-ceshishenqing"></span>
            <p>身份认证</p>
          </a>
        </li>
        <li>
          <a @click="toremark">
            <span class="iconfont icon-pingjia"></span>
            <p>使用评价</p>
          </a>
        </li>
        <li>
          <a @click="tonotice">
            <span class="iconfont icon-gonggao"></span>
            <p>政务公告</p>
          </a>
        </li>
        <li>
          <a @click="toprogress">
            <span class="iconfont icon-jindu"></span>
            <p >办理进度</p>
          </a>
        </li>
        <li>
          <a href="#top">
            <span class="iconfont icon-jindu"></span>
            <p>回到顶部</p>
          </a>
        </li>
      </ul>
    </div>
  </template>

  <script>
  export default {
    name: "side",
    data(){
      return{
      }
    },
    methods:{
      toprogress(){
        this.$router.push({name:'myevent'})
      },
      toremark(){
        this.$router.push({name:'remark'})
      },
      todoc(){
        this.$router.push({name:'doc'})
      },
      // tonotice(){
      //   this.$router.push({name:'notice'})
      // },
      toid(){
        this.$router.push({name:'identify'})
      },
      async tonotice() {
        // try {
        //     const res = await this.$store.dispatch('user/reqPublicInfo')
        //   //   console.log("getapi-try")
        // }
        // catch (error) {
        //   console.log(error.message)
        // }
        this.$router.push({ name: 'notice' })
        // window.location.reload()
    },
    }
  }
  </script>

  <style scoped>
  body{
    background-color: #ddd;
    height: 2000px;
  }
  ul{
    margin: 0;
    padding: 0;
    list-style: none;
  }
  p{
    margin: 0;
  }
  a{
    text-decoration: none;
  }
  .sidebar{
    /* border: 1px solid red; */
    width: 80px;
    position: fixed;
    right: 20px;
    bottom: 200px;
  }
  .sidebar ul{
    /* border:1px solid red; */
  }
  .sidebar ul li{
    background-color: ghostwhite;
    width: 85px;
    height: 85px;
    position: relative;
    border-bottom: 1px solid #ddd;
  }
  .sidebar ul li:last-child{
    margin-top: 20px;
  }
  .sidebar ul li span{
    width: 85px;
    height: 50px;
    /* background-color: skyblue; */
    display: block;
    font-size: 28px;
    color: #666;
    text-align: center;
    line-height: 50px;
  }
  .sidebar ul li p{
    font-size: 14px;
    text-align: center;
    color: #666;
  }
  .sidebar ul li:hover span,.sidebar ul li:hover p{
    color: rgb(155, 204, 237);
  }
  .sidebar ul li .wxin{
    height: 100px;
    background-color: rgb(255, 255, 255);
    position: absolute;
    top:0px;
    left:-130px;
    padding: 15px;
    display:none;
  }
  .sidebar ul li:hover .wxin{
    display: block;
  }

  </style>
