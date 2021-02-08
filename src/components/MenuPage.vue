<template>
  <div>
    <div id="menu" @mouseover="hover = true"><i></i></div>
    <div class="list" :class="{'closed':!hover&&!lock&&!editMode} " @mouseleave="hover = false">
      <div class="actionBar">
        <img class="modify" src="../../public/static/modify.svg" @click="modify"/>
        <img class="modify" :src="pinImage"
             @click="ding"/>
      </div>
      <ul v-for="category in originalList" :key="category.id">
        <!------>
        <li class="title">
          {{ category[0] }}
        </li>
        <li class="edit-website" v-for="item in category[1]" :key="item.id">
          <a rel="nofollow" :href="item['url']"
             target="_blank">
            <img
                :src="'https://www.google.com/s2/favicons?domain='+item['url'].replace('https://','').replace('http://','')"/>
            {{ item['name'] }}

          </a>

        </li>
        <li v-show="editMode"><a rel="nofollow" href="#">
          <img src="../../public/static/add.svg"/>
          添加站点</a>
        </li>
        <!------>
      </ul>
      <ul v-show="editMode">
        <!------>
        <li class="title" @click="addNewCategory">
          添加新分类
        </li>
      </ul>
    </div>

    <div id="customize-mode-tips" v-show="editMode" style="display: inherit; left: 483.5px;">
      <h4>自定义模式</h4>
      <a href="#" id="finish-customize" style="display: inherit;" @click="exitEdit">退出</a>
    </div>

  </div>

</template>

<script>
import {$get, $post} from "@/api";
import pin from "../../public/static/pin.svg"
import pinUp from "../../public/static/pin_up.svg"

export default {
  name: "MenuPage",

  mounted: function () {
    $get("./../static/userweb.json").then(res => {

      res.data.forEach(item => {
        let category = item["category"]
        let webItem = {
          "url": item['url'],
          "name": item["name"],
          "weight": item["weight"]
        }
        if (this.originalList.get(category)) {
          this.originalList.get(category).push(webItem)
        } else {
          this.originalList.set(category, [webItem])
        }
      })
    })

  },
  data() {
    return {
      hover: false,
      originalList: new Map(),
      editMode: false,
      lock: false,
      pinImage: pin
    };
  }, methods: {
    modify() {
      this.editMode = !this.editMode
    }, ding() {
      this.lock = !this.lock
      this.pinImage = this.lock ? pinUp : pin
    }, exitEdit() {
      this.editMode = !this.editMode
      // post 登陆

      $post("/login", {
        username: "violet",
        password: "123456"
      }).then(res => {
        console.log(res)
      }).catch(e => {
        console.log(e)
      })
    }, addNewCategory() {

    }
  }
}
</script>

<style scoped lang="scss">
@import "../assets/css/style.css";

img {
  width: 14px;
  height: 14px;
  margin: 0 5px 0 8px;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}

.actionBar {
  display: flex;
  margin-top: 10px;
  flex-direction: row;
  justify-content: end;

  .modify {
    height: 20px;
    width: 20px;
    padding: 10px;
    cursor: pointer;
  }
}

#customize-mode-tips {
  position: fixed;
  top: 0;
  height: 70px;
  display: none;
  background-color: #0da081;
  color: #fff;
  text-align: center;
  width: 500px;
  opacity: .8;
  z-index: 10;
  border-radius: 0 0 5px 5px;

  h4 {
    margin-top: 10px;
    margin-bottom: 10px;
  }

  #finish-customize {
    line-height: 30px;
    display: none;
    text-decoration: none;
    color: #fff;
    border-radius: 0 0 5px 5px
  }

  #finish-customize:hover {
    color: #fff;
    background-color: #ed5851
  }
}

span.edit-website {
  display: none;
  background-color: #fff;
  color: #fff;
  text-align: center;
  vertical-align: middle;
  position: absolute;
  border-radius: 2px;
  line-height: 45px;
  height: 45px;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  z-index: 1;
  opacity: .7;
  cursor: pointer
}
</style>