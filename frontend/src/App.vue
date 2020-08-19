<template>
    <div id="app">
        <h1>Conference APP</h1>
        <nav >
            <router-link to="/">Events</router-link>
            <router-link to="/day-overview">Day overview</router-link>
            <router-link to="/people">People</router-link>
            <router-link to="/talks">Talks</router-link>
        </nav>

        <router-view/>
    </div>
</template>

<script>

    export default {
        name: 'App',
        components: {
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

    nav {
        padding: 20px 20px 20px 0;
    }
    nav a {
        padding: 10px;
        text-decoration: none;
        background: #fff;
        border-radius: 3px;
        color: rgb(0, 110, 255);
        font-weight: bold;
        margin-right: 15px;
    }

    .footer-copyright{
        position:absolute;
        bottom:0;
        width:100%;
        height:60px;
    }
</style>
