<template>
  <div id="content" @click="hideSuggestion">
    <div class="con">
      <div class="shlogo"></div>
      <div class="sou">
        <form @submit.prevent="submit(q)">
          <!--          <div class="lg" :class="isGoogle? 'googleSearch':'baiduSearch'" @click="changed"></div>-->
          <div class="lg" @click="changed" :style="{'background': `url(${searchLogo})no-repeat center/cover`}"></div>
          <!--input class="t" type="text" value="" name="t" hidden-->
          <input class="wd" :class="{'hasText':hasInputText}" type="text" @input="inputFun" placeholder="←点击图标切换搜索引擎"
                 v-model="q"
                 x-webkit-speech lang="zh-CN" @keydown="selectDown($event)"
                 autocomplete="off">
          <img src="../../public/static/search.svg"/>
        </form>
        <div id="word" v-show="showSuggestion">
          <li v-for="(item, index)  in suggestion" :key="item.id" @click="submit(item)"
              :class="{'selectItem':index === selectSuggestion}">
            <img src="../../public/static/search.svg"/>
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
import googleLogo from "@/assets/images/g.svg"
import baiduLogo from "@/assets/images/baidu.svg"
import {getJsonp} from "@/api";
import {event} from "@/utils";

export default {

  name: 'searchPage',
  data() {
    return {
      isGoogle: true,
      q: "",
      originalContent: "",
      suggestion: [],
      showSuggestion: false,
      hasInputText: false,
      selectSuggestion: -1
    }
  },
  props: {
    msg: String
  }, methods: {
    changed() {
      this.isGoogle = !this.isGoogle
    }, submit(content) {
      if (content) {
        window.open("https://www.google.com/search?q=" + content, '_blank');
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
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "../assets/css/style.css";

.googleSearch {
  background: url("~@/assets/images/g.svg") no-repeat center/cover;
}

.baiduSearch {
  background: url("~@/assets/images/baidu.svg") no-repeat center/cover;
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