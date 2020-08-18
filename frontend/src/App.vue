<template>
    <div id="app">
        <h1>Conference APP</h1>
        <Events>

        </Events>

    </div>
</template>

<script>
    import Events from "./components/Events";

    export default {
        name: 'App',
        components: {
            Events
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
