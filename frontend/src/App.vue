<template>
    <div id="app">
        <h1>Conference APP</h1>
        <button v-on:click="getLocations">Get Events</button>
        <button v-on:click="getPeople">Get People</button>
        <button v-on:click="getTopics">Get Topics</button>
        <button v-on:click="getOverview">Get Overview</button>

        <Locations locations="locations">

        </Locations>

        <Talk>

        </Talk>
        <Person>

        </Person>

        <Overview>

        </Overview>
    </div>
</template>

<script>
    import Locations from "./components/Locations";
    import Talk from "./components/Talk";
    import Person from "./components/Person";
    import Overview from "./components/Overview";

    export default {
        name: 'App',
        components: {
            Locations,
            Talk,
            Person,
            Overview
        },
        data() {
            return {
                locations: [],
                people: [],
                topics: [],
                overview: []
            };
        },
        methods: {
            getLocations() {
                fetch("http://conference.backend/locations")
                    .then(response => response.json())
                    .then(data => (this.locations = data));
            },
            getPeople() {
                fetch("http://conference.backend/persons")
                    .then(response => response.json())
                    .then(data => (this.people = data));
            },
            getTopics() {
                fetch("http://conference.backend/topics")
                    .then(response => response.json())
                    .then(data => (this.topics = data));
            },
            getOverview() {
                fetch("http://conference.backend/topics")
                    .then(response => response.json())
                    .then(data => (this.overview = data));
            }
        }
    }
</script>

<style>
    #app {
        font-family: Avenir, Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        margin-top: 60px;
    }
</style>
