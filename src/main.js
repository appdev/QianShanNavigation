import Vue from 'vue';
import Antd, {message} from 'ant-design-vue';
import App from './App';
import 'ant-design-vue/dist/antd.css';
Vue.config.productionTip = false;

Vue.use(Antd);

Vue.config.devtools = process.env.NODE_ENV === 'development'

message.config({
    top: `200px`,
    duration: 2,
    maxCount: 3,
});

new Vue({
    render: h => h(App),
}).$mount('#app')
