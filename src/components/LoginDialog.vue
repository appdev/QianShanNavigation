<template>
  <a-modal
      v-model="showStatus"
      title="登录&同步"
      centered
      :destroyOnClose="true"
      :afterClose="dialogClose"
  >
    <a-form
        id="loginFormDialog"
        :form="form"
        class="login-form"
        @submit="handleSubmit"
    >
      <a-form-item>
        <a-input
            v-decorator="[
          'userName',
          { rules: [{ required: true, message: '请输入用户名!' }] },
        ]"
            placeholder="Username"
        >
          <a-icon slot="prefix" type="user" style="color: rgba(0,0,0,.25)"/>
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-input
            v-decorator="[
          'password',
          { rules: [{ required: true, message: '请输入密码!' }] },
        ]"
            type="password"
            placeholder="Password"
        >
          <a-icon slot="prefix" type="lock" style="color: rgba(0,0,0,.25)"/>
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-checkbox
            v-decorator="[
          'remember',
          {
            valuePropName: 'checked',
            initialValue: true,
          },
        ]"
        >
          一年内免登陆
        </a-checkbox>
        <a-button type="primary" html-type="submit" :disabled="hasErrors(form.getFieldsError())"
                  class="login-form-button">
          登 录
        </a-button>
        没有账号则自动注册并登陆
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script>
import {login} from "@/api/config";
import {setCookie, showSuccess, showWarning} from "@/utils";

function hasErrors(fieldsError) {
  return Object.keys(fieldsError).some(field => fieldsError[field]);
}

export default {

  props: {
    show: {
      type: Boolean,
      default: false
    }
  }, data() {
    return {
      hasErrors,
      form: this.$form.createForm(this, {name: 'horizontal_login'}),
      showStatus: false
    }
  },

  mounted() {
    this.$nextTick(() => {
      // To disabled submit button at the beginning.
      this.form.validateFields();
    });
  },

  methods: {
    // Only show error after a field is touched.
    userNameError() {
      const {getFieldError, isFieldTouched} = this.form;
      return isFieldTouched('userName') && getFieldError('userName');
    }, dialogClose(success) {
      this.showStatus = false
      this.$emit('loginDialogClose', false,success)//子组件对openStatus修改后向父组件发送事件通知
    },
    // Only show error after a field is touched.
    passwordError() {
      const {getFieldError, isFieldTouched} = this.form;
      return isFieldTouched('password') && getFieldError('password');
    }
    ,
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields((err, values) => {
        if (!err) {
          let rember = values.remember
          let params = {
            "username": values.userName,
            "password": values.password,
            "remember": rember
          }
          login(params).then(res => {
            if (res.code === 200) {
              showSuccess(res.msg)
              setCookie("token", res.data.token)
              this.dialogClose(true)
            } else {
              showWarning(res.msg)
            }
          })

        }
      });
    },
  }, watch: {
    show(val) {
      this.showStatus = val
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
</style>