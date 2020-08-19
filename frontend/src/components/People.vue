<template>
    <div>
        <table class="table table-striped" style="width:100%">
            <thead>
            <tr>
                <th>Person Name</th>
            </tr>
            </thead>
            <tbody v-for="person in persons" :key="person.ID">
            <tr>
                <td class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent pointer"
                    @click="getAllTalksForOnePerson(person.ID, person.PersonName)">{{person.PersonName}}
                </td>
            </tr>
            </tbody>
        </table>

        <div v-show="talks.length > 0">
            <h1>Talks of: {{name}}</h1>
            <table class="table table-striped" style="width:100%">
                <thead>
                <tr>
                    <th>Talk Title</th>
                    <th>Start Date</th>
                    <th>End Date</th>
                </tr>
                </thead>
                <tbody v-for="talk in talks" :key="talk.ID">
                <tr>
                    <td>{{talk.TitleName}}</td>
                    <td>{{talk.StartDate}}</td>
                    <td>{{talk.EndDate}}</td>
                </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>


<script>
    export default {
        name: "People",
        data() {
            return {
                persons: [],
                talks: [],
                name: null
            };
        },
        methods: {
            getPeople() {
                fetch(window.location.origin + "/api/persons")
                    .then(response => response.json())
                    .then(data => (this.persons = data));
            },
            getAllTalksForOnePerson(id, name) {
                fetch(window.location.origin + "/api/persons/" + id + "/talks")
                    .then(response => response.json())
                    .then(data => (this.talks = data));
                this.name = name;
            }
        },
        beforeMount(){
            this.getPeople()
        }
    };
</script>

<style>
    .pointer {
        cursor: pointer;
    }
</style>
