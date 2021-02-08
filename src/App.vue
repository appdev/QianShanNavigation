<template>
  <div id="app" @click="pageClicked">
    <MenuPage/>
    <SearchPage/>
  </div>
</template>

<script>
import SearchPage from './components/SearchPage.vue'
import MenuPage from './components/MenuPage.vue'
import {event} from '@/utils'

export default {
  name: 'App',
  components: {
    SearchPage,
    MenuPage
  }, methods: {
    pageClicked() {
      event.$emit("mainPageClick", true)
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

* {
  padding: 0;
  margin: 0;
  outline: none;
  -webkit-tap-highlight-color: transparent;
}

html,
body {
  width: 100%;
  height: 100%;
  position: relative;
  overflow: hidden;
  background: #fff;
}

html {
  font-family: 'Noto Sans SC Sliced', PingFangSC-Light, Microsoft YaHei UI, Microsoft YaHei, helvetica, sans-serif;
  font-weight: 500;
  color: #000;
}

form,
input,
button {
  padding: 0;
  margin: 0;
  border: none;
  outline: none;
  background: none;
}

input::-webkit-input-placeholder {
  color: #ccc;
  letter-spacing: 2px;
  font-size: 16px;
}

ul,
li {
  display: block;
  list-style: none;
}

a {
  text-decoration: none;
  color: #000;
}

::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: #222d46;
  border-radius: 0px;
}

::-webkit-scrollbar-thumb {
  background: #293550;
  border-radius: 0px;
}

::-webkit-scrollbar-thumb:hover {
  background: #293550;
}

::-webkit-scrollbar-corner {
  background: #222d46;
}

body {
  background: url("https://cn.bing.com//th?id=OHR.HuntsMesa_JA-JP3140979616_1920x1080.jpg&rf=LaDigue_1920x1080.jpg&pid=hp") no-repeat center/cover;
}
</style>
