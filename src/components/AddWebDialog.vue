<template>
  <a-modal
      v-model="showStatus"
      :title="currentTitle"
      centered
      :destroyOnClose="true"
      :afterClose="dialogClose"
  >
    <a-form
        id="loginFormDialog"
        :form="form"
        class="login-form"
        autoComplete="off"
        @submit="handleSubmit"
    >

      <a-form-item v-show="showCategory" label="分类名称">
        <a-input autocomplete="off"
                 v-decorator="['category', { rules: [{ required: showCategory, message: '请输入分类名称!' }],
              initialValue:this.name }]"
                 placeholder="分类名称"
        />
      </a-form-item>
      <!--        分类选择-->
      <a-form-item v-show="showWebInput && !showCategory" label="分类名称">
        <a-select :default-value="this.name" @change="handleChange">
          <a-select-option v-for="(category,ind) in categoryArray" :key="ind">{{ category }}</a-select-option>
        </a-select>
      </a-form-item>
      <!--      站点名字-->
      <a-form-item v-show="showWebInput" label="站点名称">
        <a-input
            v-decorator="['name', { rules: [{ required: showWebInput, message: '请输入站点名称'}],
            initialValue:this.webObj != null ? this.webObj.name : '' }]"
            placeholder="站点名称"
        />
      </a-form-item>
      <a-form-item v-show="showWebInput" label="站点地址">
        <a-input
            v-decorator="['address', { rules: [{ required: showWebInput, message: '请输入站点地址' }],
             initialValue:this.webObj != null ? this.webObj.url : ''}]"
            placeholder="站点地址"
        />
      </a-form-item>
      <a-form-item v-show="false">
        <a-checkbox
            v-decorator="[
          'displayIcon',
          {
            valuePropName: 'checked',
            initialValue: false,
          },
        ]"
        >
          不显示Favicon
        </a-checkbox>
      </a-form-item>

      <!--      <div class="icon" v-show="!customIcon">-->
      <!--        <label>Favicon:</label>-->
      <!--        <img :src="iconUrl" v-show="iconUrl!==''">-->
      <!--      </div>-->

      <!--      <a-form-item v-show="showIcon" label="Favicon">-->
      <!--        <div class="icon" v-show="!customIcon">-->
      <!--          <div>-->
      <!--            <a-button type="primary" @click="loadIcon" :disabled="false" icon="sync" :loading="loading">自动获取Favicon-->
      <!--            </a-button>-->
      <!--            <a-button style="margin-left: 20px" @click="customIcon = true" type="primary">自定义Favicon</a-button>-->
      <!--          </div>-->
      <!--          <img :src="iconUrl" v-show="iconUrl!==''" @load="loadFinish">-->
      <!--        </div>-->

      <!--        <a-input v-show="customIcon"-->
      <!--                 v-decorator="['favicon', { rules: [{ required: false, message: '请输入站点Favicon' }] }]"-->
      <!--                 :placeholder="name"-->
      <!--        />-->
      <!--      </a-form-item>-->
      <a-form-item class="btn">
        <a-button type="primary" html-type="submit" :disabled="hasErrors(form.getFieldsError())"
                  class="login-form-button">
          提 交
        </a-button>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script>
import {addweb, update, updateCategory} from "@/api/config";
import {showLoading, showSuccess, showWarning} from "@/utils";

function hasErrors(fieldsError) {
  return Object.keys(fieldsError).some(field => fieldsError[field]);
}

export default {

  props: {
    show: {
      type: Boolean,
      default: false
    }, dialogType: {
      type: String,
      default: "addWeb" // 添加、修改站点、添加、修改 分类
    }, name: {
      type: String,
      default: ""
    }, webItem: {
      type: String,
      default: ""
    },
    categoryList: {
      type: String,
      default: ""
    }
  }, data() {
    return {
      hasErrors,
      form: this.$form.createForm(this),
      showStatus: false,
      showIcon: false,
      currentTitle: "",
      showCategory: false,
      showWebInput: false,
      iconUrl: "",
      loading: false,
      customIcon: false,
      webObj: {},
      categoryArray: []
    }
  },

  mounted() {
    this.$nextTick(() => {
      // To disabled submit button at the beginning.
      this.form.validateFields();
    });
  },
  methods: {
    dialogClose(success) {
      this.showStatus = false
      this.$emit('addWebDialogClose', false, success)//子组件对openStatus修改后向父组件发送事件通知
    }, handleChange(value) {
      this.webObj.category = this.categoryArray[value]
    },
    handleSubmit(e) {
      e.preventDefault();
      let inputValues
      this.form.validateFields((err, values) => {
        if (!err) {
          inputValues = values
        } else showWarning(err)
      });
      switch (this.dialogType) {
        case "addSite":
          this.addSite(inputValues)
          break
        case "modifySite":
          if (inputValues)
            this.modifySite(inputValues)
          break
        case "addClassification":
          this.addClassification(inputValues)
          break
        case "modifyClassification": // 修改分类
          this.modifyClassification(inputValues)
          break
      }

    }, loadIcon() {
      this.form.validateFields((err, values) => {
        if (!err) {
          this.iconUrl = "https://www.google.com/s2/favicons?domain=" + values.address
          this.loading = true
        }
      });
    }, loadFinish() {
      this.loading = false
    },
    // 修改站点
    modifySite(inputValues) {
      showLoading()
      let params = {
        "id": this.webObj.id,
        "name": inputValues.name,
        "url": inputValues.address,
        "category": this.webObj.category
      }
      update(params).then(res => {
        if (res.code === 200) {
          showSuccess("修改成功")
          this.dialogClose(true)
        } else {
          showWarning(res.msg)
        }
      }).catch(() => {
        showWarning("服务器访问异常，请检查网络")
      })
    },
    // 修改分类名字
    modifyClassification(inputValues) {
      showLoading()
      let params = {
        "oldCategory": this.name,
        "newCategory": inputValues.category,
      }
      updateCategory(params).then(res => {
        if (res.code === 200) {
          showSuccess("修改成功")
          this.dialogClose(true)
        } else {
          showWarning(res.msg)
        }
      }).catch(() => {
        showWarning("服务器访问异常，请检查网络")
      })
    }, addClassification(inputValues) {
      showLoading()
      let params = {
        "name": inputValues.name,
        "url": inputValues.address,
        "category": inputValues.category,
        "displayIcon": inputValues.displayIcon
      }
      this.addNewWeb(params)
    }, addSite(inputValues) {
      showLoading()
      let params = {
        "name": inputValues.name,
        "url": inputValues.address,
        "displayIcon": inputValues.displayIcon,
        "category": this.name
      }
      this.addNewWeb(params)
    }, addNewWeb(params) {
      showLoading()
      addweb(params).then(res => {
        if (res.code === 200) {
          showSuccess("添加成功")
          this.dialogClose(true)
        } else {
          if (res.msg)
            showWarning(res.msg)
          else showWarning("添加失败")
        }
      }).catch(() => {
        showWarning("服务器访问异常，请检查网络")
      })
    }
  },
  watch: {
    show(val) {
      this.showStatus = val
    }, webItem(val) {
      if (val !== '') {
        this.webObj = JSON.parse(val)
      } else this.webObj = {}
      this.customIcon = false
      this.iconUrl = "https://www.google.com/s2/favicons?domain=" + this.webObj.url
    }, categoryList(val) {
      if (val !== '') {
        this.categoryArray = JSON.parse(val)
      } else this.webObj = []
    },
    dialogType(val) {
      this.loading = false
      switch (val) {
        case "addSite":
          this.currentTitle = "添加新站点"
          this.showIcon = true
          this.showCategory = false
          this.showWebInput = true
          break
        case "modifySite":
          this.currentTitle = "修改站点"
          this.showIcon = true
          this.showCategory = false
          this.showWebInput = true
          break
        case "addClassification":
          this.currentTitle = "添加新分类"
          this.showIcon = true
          this.showWebInput = true
          this.showCategory = true
          break
        case "modifyClassification": // 修改分类
          this.currentTitle = "修改分类"
          this.showIcon = false
          this.showWebInput = false
          this.showCategory = true
          break
      }
    }
  },
}

</script>
<style scoped lang="scss">
#loginFormDialog .login-form {
  max-width: 300px;
}

#loginFormDialog .login-form-forgot {
  float: right;
}

#loginFormDialog .login-form-button {
  width: 100%;
}

input::-webkit-input-placeholder {
  /* placeholder颜色 */
  color: #aab2bd;
  /* placeholder字体大小 */
  font-size: 12px;
}

::v-deep .ant-modal-footer {
  display: none !important;
}

::v-deep .ant-modal-body {
  padding: 10px 24px 24px 24px;
}

::v-deep .ant-form-item {
  margin-bottom: 10px;
}

.icon {
  display: flex;
  flex-direction: row;
  margin-top: 20px;
  align-items: center;

  img {
    width: 30px;
    margin-left: 20px;
    height: 30px;
  }
}

.btn {
  margin-top: 20px;
}
</style>