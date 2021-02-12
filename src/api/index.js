import axios from 'axios'
import Qs from 'qs'
import {getCookie} from "@/utils";

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

// response interceptor
axios.interceptors.response.use(
    response => {
        // 如果返回的状态码为200，说明接口请求成功，可以正常拿到数据
        // 否则的话抛出错误
        if (response.status === 200) {
            return Promise.resolve(response.data);
        } else {
            return Promise.reject(response);
        }
        // if (response.config.url === "./../static/userweb.json"){
        //     return Promise.resolve(response);
        // }else {
        //     if (response.status === 200) {
        //         return Promise.resolve(response.data);
        //     } else {
        //         return Promise.reject(response);
        //     }
        // }
    },
    // 服务器状态码不是2开头的的情况
    // 这里可以跟你们的后台开发人员协商好统一的错误状态码
    // 然后根据返回的状态码进行一些操作，例如登录过期提示，错误提示等等
    // 下面列举几个常见的操作，其他需求可自行扩展
    // error => {
        // if (error.response.status) {
        //     switch (error.response.status) {
        //         // 401: 未登录
        //         // 未登录则跳转登录页面，并携带当前页面的路径
        //         // 在登录成功后返回当前页面，这一步需要在登录页操作。
        //         case 401:
        //             router.replace({
        //                 path: '/login',
        //                 query: {
        //                     redirect: router.currentRoute.fullPath
        //                 }
        //             });
        //             break;
        //
        //         // 403 token过期
        //         // 登录过期对用户进行提示
        //         // 清除本地token和清空vuex中token对象
        //         // 跳转登录页面
        //         case 403:
        //             Message({
        //                 message: '登录过期，请重新登录',
        //                 duration: 1000,
        //                 forbidClick: true
        //             });
        //             // 清除token
        //             localStorage.removeItem('token');
        //             store.commit('loginSuccess', null);
        //             // 跳转登录页面，并将要浏览的页面fullPath传过去，登录成功后跳转需要访问的页面
        //             setTimeout(() => {
        //                 router.replace({
        //                     path: '/login',
        //                     query: {
        //                         redirect: router.currentRoute.fullPath
        //                     }
        //                 });
        //             }, 1000);
        //             break;
        //
        //         // 404请求不存在
        //         case 404:
        //             Message({
        //                 message: '网络请求不存在',
        //                 duration: 1500,
        //                 forbidClick: true
        //             });
        //             break;
        //         // 其他错误，直接抛出错误提示
        //         default:
        //             Message({
        //                 message: error.response.data.message,
        //                 duration: 1500,
        //                 forbidClick: true
        //             });
        //     }
        //     return Promise.reject(error.response);
        // }

    // }
    );


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

export const $postJS = (url, params) => {
    const options = {headers: {'Content-Type': 'application/json'}}
    return axios.post(url, params, options)
};

