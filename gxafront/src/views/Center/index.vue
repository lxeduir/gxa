<template>
    <div id="Center" v-if="isReloadData">
        <div class="picture">
            <a>个人中心</a>
        </div>
        <div class="divider">

        </div>
        <div class="backgroud">
            <div class="nav">
                <div class="login-info">
                    <img v-if="this.sex==2" src="../../assets/girl.png">
                    <img v-else-if="this.sex==1" src="../../assets/boy.png">
                    <!-- <img v-else src="../../assets/xxx.png"> -->
                    <!-- <img src="../../assets/boy.png"> -->
                    <a>您好！{{username}}</a>
                </div>
                <el-menu
                default-active="account"
                class="el-menu-vertical-demo"
                @select="changeSidebar">
                <el-menu-item index="account">
                    <i class="el-icon-document"></i>
                    <span slot="title">我的资料</span>
                </el-menu-item>
                <el-menu-item index="myremark">
                    <i class="el-icon-edit-outline"></i>
                    <span slot="title">我的评价</span>
                </el-menu-item>
                <el-menu-item index="myevent">
                    <i class="el-icon-s-order"></i>
                    <span slot="title">我的办理</span>
                </el-menu-item>
                </el-menu>
            </div>
            <div class="message">
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>

<script>
import { validUsername } from '@/utils/validate';

export default {
    data(){
        return {
            sex: '',
            username: '',
            isReloadData: false
        };
    },
    created() {
        this.getuserinfo()
    },
    methods:{ 
        changeSidebar(path) {
            this.$router.push(path);
        },
        async getuserinfo() {
            try {
                const res = await this.$store.dispatch('user/reqIdentityInfo')
                this.sex = res[0].Sex
                this.username = this.$store.getters.username
                this.isReloadData = true
                // res[0].Sex = String(res[0].Sex)
                // this.identityInfo = res[0]
                // this.form = res[0]
            }
            catch (error) {
                console.log(error.message)
            }

        },
      
    }

}
</script>

<style scoped lang='less'>
.picture {
    height: 300px;
    width: 1540px;
    background-color: #6db6ff;

}

.picture>a {
    position: absolute;
    top: 200px;
    left: 100px;
    font-size: 30px;
    color: white;
}

.backgroud {
    border: 1px solid rgb(190, 189, 189);
    position: relative;
    top: -120px;
    left: 170px;
    width: 1250px;
    height: 1000px;
    background-color: white;
    box-shadow: 3px 3px 3px 3px #c7c7c7;
}

.nav {
    border-right: 1px solid grey;
    position: relative;
    width: 250px;
    height: 1000px;

}

.nav p {
    font-size: 20px;
    margin-top: 50px;
}

.login-info {
    /*border: 1px solid red;*/
    width: 200px;
    height: 200px;
    padding: 8px 24px 0;
    text-align: center;
    position: relative;
    top: -50px;
}

.login-info img{
    /*border: 1px solid red;*/
    position: relative;
    top: -50px;
    left: -30px;
    width: 260px;
}

.login-info a{
    font-size: 25px;
    position: relative;
    top: -80px;
}

.message {
    /*border: 1px solid red;*/
    width: 950px;
    height: 1000px;
    position: relative;
    top: -1000px;
    left: 250px;
    margin: 10px;
}
</style>