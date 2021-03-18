import axios from 'axios'
import {getCookie} from "@/utils";
import * as Qs from "qs";

export const baseURL = 'https://tools.io.sb';
export const EmptyUrl = 'https://static.apkdv.com/image/web.png!/format/webp/lossless/true';


axios.defaults.headers['Content-Type'] = 'application/json';
axios.defaults.headers.common['token'] = getCookie('token') || '';
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
axios.defaults.baseURL = baseURL;

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
// 添加请求拦截器，在发送请求之前做些什么(**具体查看axios文档**)--------------------------------------------
axios.interceptors.request.use(function (config) {
    // 显示loading
    return config
}, function (error) {
    // 请求错误时弹框提示，或做些其他事
    return Promise.reject(error)
})

// 添加响应拦截器(**具体查看axios文档**)----------------------------------------------------------------
axios.interceptors.response.use(function (res) {
    return res.data
}, function (error) {
    // 对响应错误做点什么
    return Promise.reject(error)
})
// 封装axios--------------------------------------------------------------------------------------
export const $get = (url, params) => {
    return axios.get(url, {params: params})
}

export const $post = (url, params) => {
    const options = {
        headers: {'Content-Type': 'application/x-www-form-urlencoded'}
    }
    return axios.post(url, params, options)
}

export const $postQs = (url, params) => {
    const options = {
        headers: {'Content-Type': 'application/x-www-form-urlencoded'}
    }
    return axios.post(url, Qs.stringify(params), options)
}

export const $postJS = (url, params) => {
    const options = {headers: {'Content-Type': 'application/json'}}
    return axios.post(url, params, options)
};
export const $postUp = (url, params) => {
    const options = {headers: {'Content-Type': 'multipart/form-data'}};
    return axios.post(url, params, options)
};
export default axios
