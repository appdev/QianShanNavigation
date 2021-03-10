import {$get, $post} from "@/api/index";


export const login = function (params) {
    return $post('/login', params)
}
export const getJson = function (params) {
    return $get('/static/json/userweb.json', params)
}

export const getUserWebList = function (params){
    return $get('/finds', params)
}