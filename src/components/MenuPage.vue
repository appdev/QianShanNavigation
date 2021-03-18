<template>
  <div>
    <div class="login_bar">
      <img :src="loginImage" class="login" alt="退出登录" @click="addNew('login','','')">
      <img v-show="showEditBar" src="https://static.apkdv.com/image/refresh.webp" alt="切换背景" class="login"
           @click="refreshBg">
      <img class="login" :src="imagePin" alt="禁止自动切换背景" @click="disAbleChanged"/>
      <span :class="{'hindColor':showDes}" @mouseleave="hind" @mouseenter="showText">{{ imageDes }}</span>
    </div>

    <div id="menu" :class="{'on':hover}" @click="openOrClose"><i></i></div>
    <div class="list" @scroll.passive="getScroll()" ref="container" :class="{'closed':!hover&&!lock&&!editMode} ">
      <div class="actionBar" v-show="showEditBar" :class="{affixBg:top>20}">
        <img class="modify editSize" :src="editImage" @click="modify" alt="自定义模式"/>
      </div>

      <ul v-for="(category,ind) in originalList" :key="ind">
        <!------>
        <li class="title" :class="{'editModeClass':editMode}" @click="addNew('modifyClassification',category[0])">
          {{ category[0] }}
        </li>
        <li class="edit-website" v-for="(item,index) in category[1]" @click.stop="handleTouchStart(index,ind,item)"
            :key="index" @mouseleave="leave()"
            @mouseenter="enter(index,ind)">
          <a v-show="!(showEditItem===index && showEditCategory === ind)" rel="nofollow" :href="item['url']"
             target="_blank">
            <img :src="item.favicon"
                 onerror="javascript:this.src='https://static.apkdv.com/image/web.png!/format/webp/lossless/true';"/>
            {{ item.name }}
            <!--                :src="'https://www.google.com/s2/favicons?domain='+item.url"/>-->
          </a>
          <div v-show="showEditItem===index && showEditCategory === ind" class="editBox">
            <a class="edit_text" @click="addNew('modifySite',category[0],item)" href="#">编辑</a>
            <a class="delete_text" @click="showConfirm(item)" href="#">删除</a>
          </div>

        </li>
        <li v-show="editMode" @click="addNew('addSite',category[0])"><a href="#" title="添加新站点">
          <img src="https://static.apkdv.com/image/add.svg"/>
          添加站点</a>
        </li>
        <!------>
      </ul>
      <ul v-show="editMode">
        <!------>
        <li class="addNew" @click="addNew('addClassification')"><a href="#" title="添加新分类">
          <img class="addCategory" src="https://static.apkdv.com/image/add.svg"/> 添加新分类</a>
        </li>
      </ul>
    </div>

    <div id="customize-mode-tips" v-show="editMode">
      <h4>自定义模式</h4>
      <a href="#" id="finish-customize" style="display: inherit;" @click="modify">退出</a>
    </div>
    <LoginDialog :show="showLogin" @loginDialogClose="closeLoginDialog"/>
    <AddNew :show="addNewWeb" :destroy-on-close="true"
            :dialogType="addNewType" :name="categoryName"
            :webItem="webItem"
            :categoryList="categoryList"
            @addWebDialogClose="closeAddNewDialog"/>
  </div>

</template>

<script>
import LoginDialog from "@/components/LoginDialog";
import AddNew from "@/components/AddWebDialog"
import {event, getCookie, setCookie, showSuccess, showWarning} from "@/utils";
import {deleteItem, getUserWebList} from "@/api/config";
import axios from "axios";
import {disEdit, edit, login, logout, pin, pinUp} from "@/utils/image";

export default {

  name: "MenuPage",
  components: {
    LoginDialog,
    AddNew
  },
  mounted: function () {
    this.token = getCookie("token")
    this.showEditBar = this.token
    this.getJson()
    if (this.token) {
      this.loginImage = logout
    } else this.loginImage = login
    // event.$on("searchPageClick", (val) => {//监听aevent事件
    //   if (val) {
    //     this.hover = false
    //   }
    // })
    this.lockImage = localStorage.getItem("lockImage")
    this.imageDes = localStorage.getItem("des")
    this.hover = localStorage.getItem("ding") === "1"
  },
  data() {
    return {
      hover: false,
      originalList: new Map(),
      editMode: false,
      lock: false,
      lockImage: false,
      showLogin: false,
      addNewWeb: false,
      showEditBar: false,
      showDes: false,
      pinImage: pin,
      imagePin: pin,
      editImage: disEdit,
      addNewType: "",
      showEditItem: -1,
      showEditCategory: -1,
      categoryName: "",
      webItem: '',
      loginImage: login,
      categoryList: "",
      token: "",
      imageDes: "",
      top: 13,
    };
  }, methods: {
    openOrClose() {
      this.hover = !this.hover
      localStorage.setItem("ding", this.hover ? "1" : "0")
    },
    getScroll() {
      this.top = this.$refs.container.scrollTop
    }, disAbleChanged() {
      this.lockImage = !this.lockImage
      this.imagePin = this.lockImage ? pinUp : pin
      if (this.lockImage) {
        localStorage.setItem("lockImage", this.lockImage)
        showSuccess("固定使用当前壁纸")
      } else {
        showSuccess("已开启每日壁纸自动切换,当前壁纸会在下次刷新后更新")
        localStorage.setItem("lockImage", "")
        localStorage.setItem("time", "")
      }
    },
    getJson() {
      if (this.token) {
        this.getUserWebList()
      } else {
        this.accessToLocalData()
      }
    }, accessToLocalData() {
      axios.create({
        baseURL: ""
      }).get("/start/json/userweb.json").then(res => {
        this.makeData(res.data)
      })
    }, handleTouchStart(index, ind) {
      if (this.editMode) {
        this.showEditItem = index
        this.showEditCategory = ind
        // } else {
        //   // e.preventDefault()//添加阻止click事件触发
        //   window.open(item.url, '_blank')
      }
    }, errorImg() {
      let img = event.srcElement;
      img.src = "images/defaultImg.png";
      img.onerror = null; //解绑onerror事件
    },
    modify() {
      this.editMode = !this.editMode
      this.editImage = this.editMode ? edit : disEdit
    }, ding() {
      this.lock = !this.lock
      this.pinImage = this.lock ? pinUp : pin
    }, addNew(type, name, item) {
      if (type === "login") {
        if (this.token) {
          this.token = ""
          setCookie("token", "")
          this.loginImage = login
          this.showEditBar = false
          this.getJson()
          localStorage.setItem("lockImage", "")
          localStorage.setItem("time", "")
          showSuccess("账号已退出")
        } else
          this.showLogin = true
      } else if ((type === "modifyClassification" ||
          type === "addSite" ||
          type === "modifySite" ||
          type === "addClassification") && this.editMode) {
        this.addNewWeb = true
        this.webItem = JSON.stringify(item)
        this.categoryName = name
        this.addNewType = type
      }
    }, closeLoginDialog(data, loginSuccess) {
      this.showLogin = data
      if (loginSuccess) {
        this.showEditBar = true
        this.loginImage = logout
        this.token = getCookie("token")
        this.getUserWebList()
      }
    }, getUserWebList() {
      getUserWebList().then(res => {
        if (res.code === 200)
          this.makeData(res.data)
      }).catch((err) => {
        if (!err.status) {
          this.showEditBar = false
          showWarning("服务器访问异常，正在切换到原始导航数据，请检查网络")
          this.accessToLocalData()
        } else {
          showWarning("服务器访问异常")
        }
      })
    }, makeData(res) {
      this.originalList = new Map()
      res.forEach(item => {
        let category = item.category
        if (this.originalList.get(category)) {
          this.originalList.get(category).push(item)
        } else {
          this.originalList.set(category, [item])
        }
      })
      this.categoryList = JSON.stringify(Array.from(this.originalList.keys()))
    }, closeAddNewDialog(data, success) {
      this.addNewWeb = data
      if (success) {
        this.getUserWebList()
      }
    }, deleteItem(item) {
      let params = {
        "id": item.id
      }
      deleteItem(params).then(res => {
        if (res.code === 200) {
          showSuccess("删除成功")
          this.getUserWebList()
        } else {
          showSuccess(res.msg)
        }
      }).catch(() => {
            showWarning("服务器访问异常，请检查网络")
          }
      )
    }, showConfirm(item) {
      this.$confirm({
        title: `确定删除${item.name}?`,
        content: '删除后将无法回复',
        okText: '确定',
        okType: 'danger',
        cancelText: '取消',
        destroyOnClose: true,
        onOk: () => {
          this.deleteItem(item)
        }
      });
    }, enter(index, ind) {
      if (this.editMode) {
        this.showEditItem = index
        this.showEditCategory = ind
      }
    }, leave() {
      this.showEditItem = -1
      this.showEditCategory = -1
    }, refreshBg() {
      if (!this.token) {
        showWarning('这个功能需要登陆才能使用(￣▽￣)"')
      } else {
        event.$emit("changeImage", true)
      }
    }, showText() {
      this.showDes = true
    }, hind() {
      this.showDes = false
    }
  }, watch: {
    lockImage(val) {
      this.imagePin = val ? pinUp : pin
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
  position: fixed;
  width: 100%;
  padding-left: 20px;
  margin-bottom: 20px;
  height: 40px;
  flex-direction: row;

  .modify {
    margin-top: 8px;
    cursor: pointer;
    height: 25px;
    width: 25px;
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

.editBox {
  width: 100%;
  display: flex;
  flex-direction: row;

  .edit_text {
    flex: 1 1 auto;
    text-align: center;
  }

  .delete_text {
    flex: 1 1 auto;
    text-align: center;
    background: red;
  }
}


.login_bar {
  position: absolute;
  left: 20px;
  top: 15px;
  align-content: baseline;
  transition: 0.5s;
  display: flex;
  z-index: 100;
  flex-direction: row;

  .login {
    width: 20px;
    height: 20px;
    cursor: pointer;
  }

  span {
    margin-left: 20px;
    font-size: 17px;
    line-height: 30px;
    font-weight: 500;
    color: transparent;
  }
}

.hindColor {
  color: white !important;
}

.editModeClass:hover {
  cursor: pointer
}

.affixBg {
  background: #1D3B55;
  z-index: 500;
}
@media (max-width: 640px) {
  .login_bar{
    span{
      display: none;
    }
  }
}
</style>