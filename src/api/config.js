import {$get, $post, $postUp} from "@/api/index";


export const login = function (params) {
    return $post('/login', params)
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
export const addweb = function (params) {
    return $postUp('/addweb', params)
}
export const addCategory = function (params) {
    return $postUp('/add/category', params)
}
export const getNewImage = function (params) {
    return $get('/background', params)
}