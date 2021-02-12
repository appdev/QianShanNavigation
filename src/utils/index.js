import Vue from 'vue';
import Cookies from "js-cookie";

export var event = new Vue();


export const jsonFile = () => {
    require('../../public/static/userweb.json')
}

/**
 * @msg: 从cookie获取数据
 * @param {string} key
 */
export const getCookie = (key) => {
    const token = Cookies.get(key)
    if (token) {
        return token
    } else {
        return false
    }
}
/**
 * @msg: 从cookie删除数据
 * @param {string} key
 */
export const removeCookie = (key) => {
    Cookies.remove(key)
}

/**
 * @msg: cookie保存数据
 * @param key
 * @param value
 */

export const setCookie = (key, value) => {
    Cookies.set(key, value, {expires: 365})
}
