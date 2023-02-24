import router, { asyncRoutes, constantRoutes } from '@/router'
import path_search from '@/api/permisson'
import companyRouter from '@/router/faren'

export function filterAsyncRoutes(routes) {
  routes.filter((element, index, self) => self.indexOf(element) === index)

  return res
}

const state = {
  routes: [],
  addRoutes: []
}
// mutations 是唯一可以更改state的函数，使用SET_ROUTES定义更改方法，SET_ROUTES(state, routes)的入参routes赋值给addRoutes数组，将constantRoutes静态路由数组增加routes；

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = filterAsyncRoutes(router)
  }

}

const actions = {
  async generateRoutes({ commit }) {
    const result = await path_search()
    const res = result.data
    if (result.status == 200) {
      router.addRoutes(res)
      commit('SET_ROUTES', res)
    } else {
      return Promise.reject(new Error('falie'))
    }
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
