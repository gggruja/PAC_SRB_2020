<template>
    <div>
        Choose Event:
        <select class="selectpicker" @change="onChange($event)" data-size="5" required id="dropDown">
            <option>Select here</option>
            <option v-for="event in events" :key="event.ID" v-bind:value="event.ID">{{ event.EventName }}
            </option>
        </select>
        <table v-show="rooms.length > 0" class="table table-striped" style="width:100%">
            <thead>
            <tr>
                <th>Time \ Room Name</th>
                <th v-for="room in rooms" :key="room.ID">{{room.Room.RoomName}}</th>
            </tr>
            </thead>
            <tbody v-for="room in rooms" :key="room.ID" v-bind:value="room.ID" @load="onChangeRoom($event)">
            <tr>
                <td>
                    <p>{{ room.StartDate | dateParse('YYYY.MM.DD HH:mm:ss') | dateFormat('DD.MM.YYYY HH:mm:ss') }}</p>
                    <p>-</p>
                    <p>{{ room.EndDate | dateParse('YYYY.MM.DD HH:mm:ss') | dateFormat('DD.MM.YYYY HH:mm:ss') }}</p>
                </td>
                <td>
                    <p><b>Title: </b>{{room.TitleName}}</p>
                    <p><b>Level: </b>{{room.Level}}</p>
                    <p>
                        <b>Topics: </b>
                        <span v-for="(topic, index) in room.Topics" v-bind:key="topic">
                        <span>{{topic.TopicName}}</span>
                        <span v-if="index+1 < room.Topics.length">, </span>
                    </span>
                    </p>
                    <p>
                        <b>People: </b>
                        <span v-for="(person, index) in room.People" v-bind:key="person">
                        <span>{{person.PersonName}}</span>
                        <span v-if="index+1 < room.People.length">, </span>
                    </span>
                    </p>
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
