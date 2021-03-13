import Vue from 'vue';
import Cookies from "js-cookie";
import {message} from "ant-design-vue";

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

export const showWarning = (val) => {
    message.warning(val);
}
export const showSuccess = (val) => {
    message.success(val);
}

export const getTime = function() {
    let date = new Date();
    let month = date.getMonth() + 1;
    let strDate = date.getDate();
    if (month >= 1 && month <= 9) {
        month = "0" + month;
    }
    if (strDate >= 0 && strDate <= 9) {
        strDate = "0" + strDate;
    }
    return  date.getFullYear()  + month  + strDate;
};