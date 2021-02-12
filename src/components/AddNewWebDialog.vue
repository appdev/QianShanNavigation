<template>
  <div id="components-modal-demo-position">
    <a-modal
        type="primary"
        v-model="show"
        title="用户登录"
        cancelText="取消"
        okText="  登录  "
        @ok="handleSubmit"
        :ok-button-props="{ props: { disabled: hasErrors(form.getFieldsError()) } }"
    >
      <a-form layout="horizontal" :form="form" @submit="handleSubmit">
        <a-form-item :validate-status="userNameError() ? 'error' : ''" :help="userNameError() || ''">
          <a-input
              v-decorator="['userName',
          { rules: [{ required: true, message: '请输入用户名' }] },
        ]"
              placeholder="Username"
          >
            <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)"/>
          </a-input>
        </a-form-item>
        <a-form-item :validate-status="passwordError() ? 'error' : ''" :help="passwordError() || ''">
          <a-input
              v-decorator="[
          'password',
          { rules: [{ required: true, message: '请输入密码' }] },
        ]"
              type="password"
              placeholder="Password"
          >
            <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)"/>
          </a-input>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>
<script>
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
    }
    ,
    // Only show error after a field is touched.
    passwordError() {
      const {getFieldError, isFieldTouched} = this.form;
      return isFieldTouched('password') && getFieldError('password');
    }
    ,
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields((values) => {
        // if (!err) {
        // }

        console.log(values)
      });
    },
  },
}

</script>
<style scoped lang="scss">

input::-webkit-input-placeholder {
  /* placeholder颜色 */
  color: #aab2bd;
  /* placeholder字体大小 */
  font-size: 12px;
}

</style>