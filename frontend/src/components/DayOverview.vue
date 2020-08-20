<template>
    <div>
        Choose Event:
        <select class="selectpicker" @change="onChange($event)" data-size="5" required id="dropDown">
            <option>Select here</option>
            <option v-for="event in events" :key="event.ID" v-bind:value="event.ID">{{ event.EventName }}
            </option>
        </select>
        <table class="table table-striped" style="width:100%">
            <thead>
            <tr v-for="room in rooms" :key="room.ID">
                <th>Time \ Room Name</th>
                <th>{{room.RoomName}}</th>
            </tr>
            </thead>
            <tbody v-for="room in rooms" :key="room.ID" v-bind:value="room.ID" @load="onChangeRoom($event)">
            <tr>
                <td>{{room.RoomName}}</td>
                <td>{{room.RoomName}}</td>
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
                events: [],
                rooms: []
            };
        },
        methods: {
            getSelectBox() {
                fetch(window.location.origin + "/api/events/select-box")
                    .then(response => response.json())
                    .then(data => (this.events = data));
            },
            onChange(event) {
                fetch(window.location.origin + "/api/locations/" + event.target.value + "/rooms")
                    .then(response => response.json())
                    .then(data => (this.rooms = data));
            },
            onChangeRoom(event) {
                console.log(event.target.value);
            }
        },
        beforeMount() {
            this.getSelectBox()
        }
    };
</script>
