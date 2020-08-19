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
                <td class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent"
                    @click="getAllTalksForOnePerson(person.ID)">{{person.PersonName}}
                </td>
            </tr>
            </tbody>
        </table>
    </div>
</template>


<script>
    export default {
        name: "People",
        data() {
            return {
                persons: [],
                talks: []
            };
        },
        methods: {
            getPeople() {
                fetch(window.location.origin + "/api/persons")
                    .then(response => response.json())
                    .then(data => (this.persons = data));
            },
            getAllTalksForOnePerson(id) {
                fetch(window.location.origin + "/api/persons/" + id + "/talks")
                    .then(response => response.json())
                    .then(data => (this.talks = data));
            }
        },
        beforeMount(){
            this.getPeople()
        }
    };
</script>
