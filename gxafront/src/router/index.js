import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '@/store'
import { getToken, getUid } from '@/utils/auth'
// 组件通过export暴露接口，路由中导入组件
import Register from '../views/Register'
import Home from '../views/Home'
import Personal from '../views/Personal'
import Identify from '../views/Identify'
import Remark from '../views/Remark'
import Progress from '../views/Progress'
import Company from '../views/Company'
import Doc from '../views/Doc'
import Login from '../views/Login'
import Notice from '../views/Notice'
import Account from '../views/Account'
import Center from '../views/Center'
import Myremark from '../components/Myremark'
import Myevent from '../components/Myevent'
import Notdetail from '../views/Notdetail'

Vue.use(VueRouter);

let router = new VueRouter({
    model: 'history',
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home,
            meta: { Show: true }
        },
        {
            path: '/register',
            name: 'register',
            component: Register,
            meta: { Show: false }
        },
        {
            path: '/personal',
            name: 'personal',
            component: Personal,
            meta: { Show: true }
        },
        {
            path: '/company',
            name: 'company',
            component: Company,
            meta: { Show: true }
        },
        {
            path: '/remark',
            name: 'remark',
            component: Remark,
            meta: { Show: true }
        },
        {
            path: '/progress',
            name: 'progress',
            component: Progress,
            meta: { Show: true }
        },
        {
            path: '/notice',
            name: 'notice',
            component: Notice,
            meta: { Show: true },
        },
        {
                path: '/notice/detail/:index',
                name: 'notdetail',
                component: Notdetail
        },
        {
            path: '/login',
            name: 'login',
            component: Login,
            meta: {
                Show: false ,
                cheackIsLogin: true
            }
        },
        {
            path: '/doc',
            name: 'doc',
            component: Doc,
            meta: { Show: true }
        },
        {
            path: '/identify',
            name: 'identify',
            component: Identify,
            meta: { Show: true }
        },
        {
            path: '/center',
            name: 'center',
            component: Center,
            redirect: "/center/account",
            meta: { Show: true },
            children:[
                {
                    path: '/center/account',
                    name: 'account',
                    component: Account,
                    meta: { Show: true }
                },
                {
                    path: '/center/myremark',
                    name: 'myremark',
                    component: Myremark,
                    meta: {Show: true}
                },
                {
                    path: '/center/myevent',
                    name: 'myevent',
                    component: Myevent,
                    meta: {Show: true}
                }
            ]
        }
    ],
    scrollBehavior(to, from, savedPosition) {
        return { y: 0 }
    }
});
//前置守卫  在路由之前进行判断
router.beforeEach(async (to, from, next) => {

    let name = store.getters.username
    let token = getToken()
    //如果用户已经登录
    if (token) {
        //去登录或者注册页面
        if (to.path == '/login' || to.path == '/register') next({ path: '/' })
        else {
            //已有用户名 代表已获取到用户信息
            if (name) { next() }
            else {
                //如果没有获取到用户信息，就派发action获取用户信息之后再跳转
                try {
                    await store.dispatch('user/getInfo')
                    next()
                } catch (error) {
                    //token失效了 获取不到用户信息
                    //清除token
                    await store.dispatch('user/logout')
                    next({ path: '/login' })
                }
            }
        }
    } else if (to.path == '/center/account' || to.path == '/center/myremark' || to.path == '/center/myevent' || to.path == '/identify') {
                next({ path: '/' })
                window.alert("请先登录")
            } 
    else next()
    
    if(to.path == '/notice/detail/:index') next({path: '/notice'})

})

export default router
