import Vue from 'vue'
import VueRouter from 'vue-router'
import Events from './components/Events'
import People from "./components/People";
import Talks from "./components/Talks";
import DayOverview from "./components/DayOverview";

Vue.use(VueRouter)

export default new VueRouter({
    mode: 'history',
    routes: [
        {
            path: "/",
            name: "events",
            component: Events
        },
        {
            path: "/people",
            name: "people",
            component: People
        },
        {
            path: "/talks",
            name: "talks",
            component: Talks
        },
        {
            path: "/day-overview",
            name: "day overview",
            component: DayOverview
        }
    ]
})
