<template>
  <div id="content" @click="hideSuggestion">
    <div class="con">
      <lottie :options="defaultOptions" :height="250" :width="250" v-on:animCreated="handleAnimation"/>
      <div class="sou">
        <form @submit.prevent="submit(q)">
          <div class="lg" @click="changed" :style="{'background': `url(${searchLogo})no-repeat center/cover`}"></div>
          <input class="wd" :class="{'hasText':hasInputText}" type="text" @input="inputFun" placeholder="←点击图标切换搜索引擎"
                 v-model="q"
                 x-webkit-speech lang="zh-CN" @keydown="selectDown($event)"
                 autocomplete="off">
          <img src="https://static.apkdv.com/image/search.svg"/>
        </form>
        <div id="word" v-show="showSuggestion">
          <li v-for="(item, index)  in suggestion" :key="item.id" @click="submit(item)"
              :class="{'selectItem':index === selectSuggestion}">
            <img src="https://static.apkdv.com/image/search.svg"/>
            {{ item }}
          </li>
        </div>
      </div>
    </div>
    <div class="foot" style="height: 40px;">
      <a href="https://github.com/appdev/QianShanNavigation" style="color: #777;">Github</a><br>
      © 2016-2021 by <a href="https://www.apkdv.com/">LengYue</a> . All rights reserved.
    </div>
  </div>
</template>

<script>
import lottie from 'vue-lottie'
import {getJsonp} from "@/api";
import {event, getTime} from "@/utils";
import {getImage, getNewImage} from "@/api/config";
import {baiduLogo, googleLogo} from "@/utils/image";
import animationData from '@/assets/sun.json';

export default {

  name: 'searchPage',
  components: {
    'lottie': lottie
  },
  data() {
    return {
      isGoogle: true,
      q: "",
      originalContent: "",
      suggestion: [],
      showSuggestion: false,
      hasInputText: false,
      selectSuggestion: -1,
      defaultOptions: {animationData: animationData,loop:false},
      animationSpeed: 1
    }
  },
  props: {
    msg: String
  }, methods: {
    handleAnimation: function (anim) {
      this.anim = anim;
    },
    stop: function () {
      this.anim.stop();
    },

    play: function () {
      this.anim.play();
    },

    pause: function () {
      this.anim.pause();
    },
    onSpeedChange: function () {
      this.anim.setSpeed(this.animationSpeed);
    },
    setBg(urlPah, data) {
      document.querySelector('body')
          .setAttribute('style', 'background:url("' + urlPah + '") no-repeat center/cover; ')
      if (data) {
        document.querySelector('span').innerText = data;
      }
    }, loadEmptyImage() {
      let url = localStorage.getItem("image")
      let des = localStorage.getItem("des")
      if (url === '') {
        url = "https://images.unsplash.com/photo-1615594793621-bdd4c64811a4?crop=entropy&cs=tinysrgb&fit=crop&fm=jpg&h=1080&ixlib=rb-1.2.1&q=80&w=1920"
      }
      this.setBg(url, des)
    },
    loadImage() {
      let time = getTime()
      let saveTime = localStorage.getItem("time")
      let saveImage = localStorage.getItem("image")
      if (saveTime === '' || saveImage === '' || time !== saveTime) {
        getNewImage().then(res => {
          if (res.code === 200) {
            localStorage.setItem("time", time)
            localStorage.setItem("image", res.data.url)
            localStorage.setItem("des", res.data.copyright)
            this.setBg(res.data.url, res.data.copyright)
          } else {
            this.loadEmptyImage()
          }
        })
      } else this.loadEmptyImage()
    },
    changed() {
      this.isGoogle = !this.isGoogle
    }, submit(content) {
      if (content) {
        let url = this.isGoogle ? "https://www.google.com/search?q=" : "https://www.baidu.com/s?wd="
        window.open(url + content, '_blank');
        this.q = ""
        this.suggestion = []
        this.hasInputText = false
        this.showSuggestion = false
      }

    }, inputFun(e) {
      //e.target 指向了dom元素
      let inputValue = e.target.value;
      this.originalContent = inputValue
      if (inputValue) {
        this.hasInputText = true
        // 使用
        getJsonp("https://suggestion.baidu.com/su?", {
          wd: inputValue,
          cb: "jQuery33103307611929552594_1612617062297"
        }).then(json => {
          let sug = json["s"]
          if (sug.length === 0) {
            this.showSuggestion = false
          } else {
            this.suggestion = sug
            this.selectSuggestion = -1
            this.showSuggestion = true
          }
        })
      } else {
        this.showSuggestion = false
        this.hasInputText = false
      }

    },
    selectDown(e) {
      if (this.showSuggestion) {
        if (e.keyCode === 40) {
          if (this.suggestion.length - 1 > this.selectSuggestion) {
            this.selectSuggestion += 1
            this.q = this.suggestion[this.selectSuggestion]
          } else {
            this.selectSuggestion = -1
            this.q = this.originalContent
          }
        }

        if (e.keyCode === 38) {
          e.preventDefault();
          if (this.selectSuggestion === 0) {
            this.selectSuggestion = -1
            this.q = this.originalContent
          } else if (this.selectSuggestion === -1) {
            this.selectSuggestion = this.suggestion.length - 1
            this.q = this.suggestion[this.selectSuggestion]
          } else {
            this.selectSuggestion -= 1
            this.q = this.suggestion[this.selectSuggestion]
          }
        }
      }
    }
    ,
    hideSuggestion() {
      event.$emit("searchPageClick", true)
      this.showSuggestion = false
      this.hasInputText = false
    }
  },
  computed: {
    searchLogo() {
      return this.isGoogle ? googleLogo : baiduLogo
    }
  }, mounted() {
    let status = localStorage.getItem("lockImage")
    if (status)
      this.loadEmptyImage()
    else {
      this.loadImage()
    }
    event.$on("changeImage", (val) => {//监听aevent事件
      if (val) {
        getImage().then(res => {
          if (res.code === 200) {
            localStorage.setItem("image", res.data.url)
            localStorage.setItem("des", res.data.copyright)
            localStorage.setItem("time", "")
            this.setBg(res.data.url, res.data.copyright)
          } else {
            this.loadEmptyImage()
          }
        })
      }
    })
  },
  // 销毁前清除
  // beforeDestroy() {
  //   document.querySelector('body').removeAttribute('style')
  // },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "../assets/css/style.css";

.googleSearch {
  background: url("https://static.apkdv.com/image/g.svg") no-repeat center/cover;
}

.baiduSearch {
  background: url("https://static.apkdv.com/image/baidu.svg") no-repeat center/cover;
}

.selectItem {
  background: #EEEEEE;
}

li {
  line-height: 30px;

  img {
    width: 18px;
    vertical-align: middle;
    height: 18px;
  }
}

.hasText {
  background: #fff;
  box-shadow: 0 0px 15px 0 rgba(32, 33, 36, 0.2);
  border-color: #fff
}

</style>