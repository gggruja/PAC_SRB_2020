import Vue from 'vue'
import App from './App.vue'
import router from "./router";
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import VueFilterDateFormat from 'vue-filter-date-format';
import VueFilterDateParse from '@vuejs-community/vue-filter-date-parse';
import * as Keycloak from 'keycloak-js'

Vue.use(VueFilterDateParse);

Vue.use(VueFilterDateFormat);

Vue.config.productionTip = false;

let initOptions = {
    url: 'https://conference.keycloak/auth', realm: 'PAC', clientId: 'conference.frontend', onLoad: 'login-required'
}

let keycloak = Keycloak(initOptions);

keycloak.init({onLoad: initOptions.onLoad, checkLoginIframe: false}).success((auth) => {

    if (!auth) {
        console.log("Not Authenticated");
    } else {
        console.log("Authenticated");
    }

    new Vue({
        router,
        render: h => h(App),
    }).$mount('#app')


    localStorage.setItem("vue-token", keycloak.token);
    localStorage.setItem("vue-refresh-token", keycloak.refreshToken);

    setInterval(() => {
        keycloak.updateToken(70).success((refreshed) => {
            if (refreshed) {
                console.log('Token refreshed' + refreshed);
            } else {
                console.log('Token not refreshed, valid for '
                    + Math.round(keycloak.tokenParsed.exp + keycloak.timeSkew - new Date().getTime() / 1000) + ' seconds');
            }
        }).error(() => {
            console.log('Failed to refresh token');
        });


    }, 60000)

}).error(() => {
    Vue.$log.error("Authenticated Failed");
});
