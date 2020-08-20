import Vue from 'vue'
import App from './App.vue'
import router from "./router";
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import VueFilterDateFormat from 'vue-filter-date-format';
import VueFilterDateParse from '@vuejs-community/vue-filter-date-parse';

Vue.use(VueFilterDateParse);

Vue.use(VueFilterDateFormat);

Vue.config.productionTip = false

new Vue({
    router,
    render: h => h(App),
}).$mount('#app')
