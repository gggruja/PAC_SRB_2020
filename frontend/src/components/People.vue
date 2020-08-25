<template>
    <div>
        <table class="table table-striped" style="width:100%">
            <thead>
            <tr>
                <th>Person Name</th>
                <th>See Talks</th>
                <th>Update Person</th>
            </tr>
            </thead>
            <tbody v-for="person in persons" :key="person.ID">
            <tr>
                <td v-if="editPerson === person.ID">
                    <div class="input-group mb-3">
                        <input type="text" class="form-control" v-on:keyup.13="updatePerson(person)" v-model="person.PersonName"/>
                    </div>
                </td>
                <td v-else>
                    {{person.PersonName}}
                </td>
                <td>
                    <button v-on:click="getAllTalksForOnePerson(person.ID, person.PersonName)" type="button"
                            class="btn btn-info">See Talks
                    </button>
                </td>
                <td v-if="editPerson === person.ID">
                    <button type="button" class="btn btn-success" @click="updatePerson(person)">Save</button>
                </td>
                <td v-else>
                    <button type="button" class="btn btn-primary" @click="editPerson = person.ID">Update Person</button>
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
                    <td style="font-weight:bold">{{talk.TitleName}}</td>
                    <td>{{ talk.StartDate | dateParse('YYYY.MM.DD HH:mm:ss') | dateFormat('DD.MM.YYYY HH:mm') }}</td>
                    <td>{{ talk.EndDate | dateParse('YYYY.MM.DD HH:mm:ss') | dateFormat('DD.MM.YYYY HH:mm') }}</td>
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
                name: null,
                editPerson: null
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
            },
            updatePerson(person) {
                fetch(window.location.origin + "/api/persons/" + person.ID, {
                    body: JSON.stringify(person),
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                    },
                })
                    .then(() => {
                        this.editPerson = null;
                    })
            }
        },
        beforeMount() {
            this.getPeople()
        }
    };
</script>

<style>
    .pointer {
        cursor: pointer;
    }
</style>
