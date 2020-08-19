import Vue from 'vue'
import Router from 'vue-router'
import Events from './components/Events'
import People from "./components/People";
import Talks from "./components/Talks";
import DayOverview from "./components/DayOverview";

Vue.use(Router)

export default new Router({
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
