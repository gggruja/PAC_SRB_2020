<template>
    <div>
        Choose Event:
        <select class="selectpicker" data-size="5" required id="dropDown">
            <option>Select here</option>
            <option v-for="event in events" :key="event.ID">{{ event.EventName }}
            </option>
        </select>
        <table class="table table-striped" style="width:100%">
            <thead>
            <tr>
                <th>RoomName</th>
            </tr>
            </thead>
            <tbody v-for="room in rooms" :key="room.ID">
            <tr>
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
            }
        },
        beforeMount(){
            this.getSelectBox()
        }
    };
</script>
