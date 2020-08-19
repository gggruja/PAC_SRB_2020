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
                <td>{{talk.LanguageId}}</td>
                <td>
                    <span v-for="(person, index) in talk.People" v-bind:key="person">
                        <span>{{person.PersonName}}</span>
                        <span v-if="index+1 < talk.People.length">, </span>
                    </span>
                </td>
                <td>
                    <span v-for="(topic, index) in talk.Topics" v-bind:key="topic">
                        <span>{{topic.TopicName}}</span>
                        <span v-if="index+1 < talk.Topics.length">, </span>
                    </span>
                </td>
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
                talks: [],
                peopleName: null,
                language: null
            };
        },
        methods: {
            getTalks() {
                fetch(window.location.origin + "/api/talks")
                    .then(response => response.json())
                    .then(data => (this.talks = data));
            },
            getLanguage(id) {
                fetch(window.location.origin + "/api/language/" + id)
                    .then(response => response.json())
                    .then(data => (this.language = data));
            }
        },
        beforeMount() {
            this.getTalks()
        }
    };
</script>
