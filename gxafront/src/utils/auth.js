import Cookies from "js-cookie";
const TokenKey = 'yibing_gxa'

export function getToken() {
    return Cookies.get(TokenKey)
}

export function setToken(token) {
    return Cookies.set(TokenKey, token)
}

export function removeToken() {
    console.log("已下线")
    return Cookies.remove(TokenKey)
}
export function setUser(username){
    return Cookies.set('USER',username)
}
export function getUser(){
    return Cookies.get('USER')
}
export function removeUser() {
    return Cookies.remove('USER')
}
export function setUid(uid){
    return Cookies.set('UID',uid)
}
export function getUid(){
    return Cookies.get('UID')
}

export function removeUid() {
    return Cookies.remove('UID')
}
