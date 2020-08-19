<template>
    <div>
        <table class="table table-striped" style="width:100%">
            <thead>
            <tr>
                <th>Title</th>
                <th>Level</th>
                <th>Language</th>
                <th>People</th>
                <th>Topics</th>
                <th>Events</th>
            </tr>
            </thead>
            <tbody v-for="talk in talks" :key="talk.ID">
            <tr>
                <td>{{talk.TitleName}}</td>
                <td>{{talk.Level}}</td>
                <td>TODO LANGUAGE</td>
                <td>TODO PEOPLE</td>
                <td>TODO TOPICS</td>
                <td>TODO EVENTS</td>
            </tr>
            </tbody>
        </table>
    </div>
</template>


<script>
    export default {
        name: "Talks",
        data() {
            return {
                talks: []
            };
        },
        methods: {
            getTalks() {
                fetch(window.location.origin + "/api/talks")
                    .then(response => response.json())
                    .then(data => (this.talks = data));
            }
        },
        beforeMount(){
            this.getTalks()
        }
    };
</script>
