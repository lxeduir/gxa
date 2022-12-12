import Cookies from "js-cookie";
const TokenKey = 'yibing_gxa'

export function getToken() {
    return Cookies.get(TokenKey)
}

export function setToken(token) {
    return Cookies.set(TokenKey, token)
}

export function removeToken() {
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