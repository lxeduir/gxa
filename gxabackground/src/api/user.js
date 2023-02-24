import request from '@/utils/request'
import request1 from '@/utils/mockrequest'
export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: `/admininfo`,
    method: 'get',

  })
}

export function logout() {
  return request({
    url: '/vue-admin-template/user/logout',
    method: 'post'
  })
}

export function login_mock(data) {
  return request1({
    url: '/vue-admin-template/user/login',
    method: 'post',
    data
  })
}
