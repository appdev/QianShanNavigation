import {$get, $post, $postUp} from "@/api/index";


export const login = function (params) {
    return $post('/login', params)
}
export const getJson = function (params) {
    return $get('/static/json/userweb.json', params)
}

export const getUserWebList = function (params) {
    return $get('/finds', params)
}
export const deleteItem = function (params) {
    return $get('/delete', params)
}
export const update = function (params) {
    return $postUp('/update', params)
}

export const updateCategory = function (params) {
    return $postUp('/update/category', params)
}