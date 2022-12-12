import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'
const service = axios.create({
  baseURL: 'https://api.edulx.xyz/front',
  timeout: 5000 // request timeout
})

service.interceptors.request.use(
  config => {
    const jwt=store.getters.token
    if (jwt) {
      config.headers['Authorization'] = jwt
    }
    return config
  },
  error => {
    console.log(error)
    return Promise.reject(error)
  }
)

service.interceptors.response.use(
  response => {
    console.log(response)
    const res = response.data
    if (response.status!==200) {
      Message({
        message: res.message || res.err||res.error,
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(new Error(res.message || res.error))
    } else {
      if(res.err!=null){
        Message({
          message: res.err,
          type: 'error',
          duration: 5 * 1000
        })
      }
       else return response
    }
  },
  error => {
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)
export default service
