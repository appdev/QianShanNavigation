<template>
  <div>
    <div class="login_bar">
      <img :src="loginImage" class="login" @click="addNew('login','','')">
      <img src="../../public/static/refresh.svg" class="login" @click="refreshBack">
    </div>
    <div id="menu" @mouseover="hover = true"><i></i></div>
    <div class="list" :class="{'closed':!hover&&!lock&&!editMode} " @mouseleave="hover = false">
      <div class="actionBar">
        <img class="modify" src="../../public/static/modify.svg" @click="modify"/>
        <img class="modify" :src="pinImage"
             @click="ding"/>
      </div>
      <ul v-for="(category,ind) in originalList" :key="ind">
        <!------>
        <li class="title" :class="{'editModeClass':editMode}" @click="addNew('modifyClassification',category[0])">
          {{ category[0] }}
        </li>
        <li class="edit-website" v-for="(item,index) in category[1]" :key="index" @mouseleave="leave()"
            @mouseenter="enter(index,ind)">
          <a v-show="!(showEditItem===index && showEditCategory === ind)" rel="nofollow" :href="item['url']"
             target="_blank">
            <img
                :src="'https://www.google.com/s2/favicons?domain='+item['url'].replace('https://','').replace('http://','')"/>
            {{ item['name'] }}
          </a>
          <a class="edit_text" @click="addNew('modifySite','',item)" href="#"
             v-show="showEditItem===index && showEditCategory === ind">编辑</a>
        </li>
        <li v-show="editMode" @click="addNew('addSite')"><a href="#">
          <img src="../../public/static/add.svg"/>
          添加站点</a>
        </li>
        <!------>
      </ul>
      <ul v-show="editMode">
        <!------>
        <li class="addNew" @click="addNew('addClassification')"><a href="#">
          <img class="addCategory" src="../../public/static/add.svg"/> 添加新分类</a>
        </li>
      </ul>
    </div>

    <div id="customize-mode-tips" v-show="editMode">
      <h4>自定义模式</h4>
      <a href="#" id="finish-customize" style="display: inherit;" @click="exitEdit">退出</a>
    </div>
    <LoginDialog :show="showLogin" @loginDialogClose="closeLoginDialog"/>
    <AddNew :show="addNewWeb" :destroy-on-close="true"
            :dialogType="addNewType" :name="categoryName"
            :webItem="webItem"
            @addWebDialogClose="closeAddNewDialog"/>
  </div>

</template>

<script>
import pin from "../../public/static/pin.svg"
import pinUp from "../../public/static/pin_up.svg"
import login from "../../public/static/login.svg"
import logout from "../../public/static/login_out.svg"
import LoginDialog from "@/components/LoginDialog";
import AddNew from "@/components/AddWebDialog"
import {getCookie, setCookie} from "@/utils";
import {getJson, getUserWebList} from "@/api/config";

export default {
  name: "MenuPage",
  components: {
    LoginDialog,
    AddNew
  }, watch: {
    getToken: {
      handler(newVal) {
        console.log(newVal)
        if (newVal) {
          this.loginImage = logout
        } else this.loginImage = login
      },
      immediate: true
    }
  }, computed: {
    getToken() {
      return getCookie("token")
    }
  },
  mounted: function () {
    getJson().then(res => {
      this.makeData(res)
    })
  },
  data() {
    return {
      hover: false,
      originalList: new Map(),
      editMode: false,
      lock: false,
      showLogin: false,
      addNewWeb: false,
      pinImage: pin,
      addNewType: "",
      showEditItem: -1,
      showEditCategory: -1,
      categoryName: "",
      webItem: '',
      loginImage: login,

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
    }, addNew(type, name, item) {
      console.log(type)
      if (type === "login") {
        if (this.getToken) {
          setCookie("token", "")
        } else
          this.showLogin = true
      } else if (type === "modifyClassification" ||
          type === "addSite" ||
          type === "modifySite" ||
          type === "addClassification") {
        // if (type === "addSite") {
        //   this.addNewWeb = true
        // } else {
        //   this.addNewWeb = false
        // }
        this.addNewWeb = true
        this.webItem = JSON.stringify(item)
        this.categoryName = name
        this.addNewType = type
      }
    }
    ,
    closeLoginDialog(data, loginSuccess) {
      this.showLogin = data
      if (loginSuccess) {
        this.getUserWebList()
      }
    }, getUserWebList() {
      getUserWebList().then(res => {
        this.makeData(res.data)
      })
    }, makeData(res) {
      res.forEach(item => {
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
    }
    ,
    closeAddNewDialog(data) {
      console.log(data)
      this.addNewWeb = data
    }
    ,
    enter(index, ind) {
      if (this.editMode) {
        this.showEditItem = index
        this.showEditCategory = ind
      }
    }
    ,
    leave() {
      this.showEditItem = -1
      this.showEditCategory = -1

    }
    ,
    refreshBack() {

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

.addNew {
  width: 100%;

  a {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    padding-left: 10px;
    padding-right: 10px;
    width: 100%;
    height: 40px;
    align-items: center;
    margin: 10px 0;
    text-indent: 10px;
    vertical-align: middle;
    text-align: center;
    color: #fff;
    font-size: 18px;
    font-weight: bold;
  }

  img {
    width: 20px;
    height: 20px;
    fill: currentColor;
    margin: 0 5px 0 0;
    line-height: 40px;
    overflow: hidden;
  }

}


.actionBar {
  display: flex;
  margin-top: 10px;
  flex-direction: row;

  .modify {
    height: 20px;
    width: 20px;
    cursor: pointer;
  }
}

#customize-mode-tips {
  position: fixed;
  top: 0;
  height: 70px;
  display: inherit;
  background-color: #0da081;
  color: #fff;
  text-align: center;
  width: 500px;
  opacity: .8;
  margin: auto;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
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

.edit_text {
  text-align: center;
}

.login_bar {
  transform: scale(0.8);
  position: absolute;
  left: 10px;
  top: 15px;
  cursor: pointer;
  transition: 0.5s;
  display: flex;
  z-index: 100;
  flex-direction: row;

  .login {
    width: 30px;
    height: 30px;
  }
}

.editModeClass:hover {
  cursor: pointer
}


</style>