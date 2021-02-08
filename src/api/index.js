import axios from 'axios'
import Qs from 'qs'
import baseURL from './config.js'
import Cookies from 'js-cookie'

axios.defaults.baseURL=baseURL;
axios.defaults.headers['Content-Type']='application/json';
axios.defaults.headers.common['token'] = getCookie('token') || '';
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';

/**
 * jsonp跨域
 * @param {String} url [请求地址]
 * @param {Object} params [请求携带参数]
 */
axios.jsonp = (url, params) => {
    if (!url) {
        console.error('至少需要一个url参数!')
        return
    }
    return new Promise((resolve) => {
        window["jQuery33103307611929552594_1612617062297"] = (result) => {
            resolve(result)
        }
        const JSONP = document.createElement('script')
        JSONP.type = 'text/javascript'
        JSONP.src = `${url}${Qs.stringify(params)}`
        document.getElementsByTagName('head')[0].appendChild(JSONP)
        setTimeout(() => {
            document.getElementsByTagName('head')[0].removeChild(JSONP)
        }, 500)
    })
}
/**
 * 接口方法
 * @param {String} url [请求地址]
 * @param {Object} params [请求携带参数]
 */
export const getJsonp = function (url, params) {
    return new Promise((resolve, reject) => {
        axios
            .jsonp(url, params)
            .then(res => {
                resolve(res)
            })
            .catch(error => {
                reject(error)
            })
    })
}

export const $get = (url, params) => {
    return axios.get(url, {params: params})
}
export const $post = (url, params) => {
    const options = {
        headers: {'Content-Type': 'application/x-www-form-urlencoded'}
    }
    return axios.post(url, params, options)
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
    Cookies.set(key, value, {expires: 1 || 1})
}
