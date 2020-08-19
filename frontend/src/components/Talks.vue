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
                <td v-for="person in talk.People" :key="person">{{person.PersonName}}</td>
                <td v-for="topic in talk.Topics" :key="topic">{{topic.TopicName}}</td>
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
        beforeMount(){
            this.getTalks()
        }
    };
</script>
