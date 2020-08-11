<template>
    <div>
        <h1>Conference APP</h1>
        <button v-on:click="getEvents">Get Events</button>
        <table style="width:100%">
            <thead>
            <tr>
                <th>Event Name</th>
                <th>Start Date</th>
                <th>End Date</th>
                <th>Location Name</th>
                <th>Room Name</th>
                <th>Talk Title</th>
                <th>Topic Name</th>
            </tr>
            </thead>
            <tbody v-for="event in events" :key="event.ID">
            <tr>
                <td>{{event.EventName}}</td>
                <td>{{event.StartDate}}</td>
                <td>{{event.EndDate}}</td>
                <td>{{event.LocationName}}</td>
                <td>{{event.RoomName}}</td>
                <td>{{event.TitleName}}</td>
                <td>{{event.TopicName}}</td>
            </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
    export default {
        name: "Events",
        data() {
            return {
                events: []
            };
        },
        methods: {
            getEvents() {
                fetch("http://conference.backend/api/events")
                    .then(response => response.json())
                    .then(data => (this.events = data));
            }
        }
    };
</script>
