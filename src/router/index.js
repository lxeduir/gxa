import Vue from 'vue'
import VueRouter from 'vue-router'

// 组件通过export暴露接口，路由中导入组件
import Register from '../views/Register'
import Home from '../views/Home'
import Personal from '../views/Personal'
import Identity from '../views/Identity'
import Remark from '../views/Remark'
import Progress from '../views/Progress'
import Company from '../views/Company'
import Doc from '../views/Doc'
import Login from '../views/Login'
import Notice from '../views/Notice'
// 导入 vue-router 依赖
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
            path: '/identity',
            name: 'identity',
            component: Identity,
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
            meta: { Show: true }
        },
        {
            path: '/login',
            name: 'login',
            component: Login,
            meta: { Show: false }
        },
        {
            path:'/doc',
            name: 'doc',
            component: Doc,
            meta: { Show: true }
        }
    ],
    scrollBehavior(to, from, savedPosition) {
        return { y: 0 }
    }
});
// router.beforeEach((to,from,next)=>{
//     if(from.path=='/login') router.go(0)
// })
export default router
