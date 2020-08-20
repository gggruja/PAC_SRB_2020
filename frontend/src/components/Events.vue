<template>
    <div>
        <table class="table table-striped" style="width:100%">
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
                <td>{{ event.StartDate | dateParse('YYYY.MM.DD HH:mm:ss') | dateFormat('DD.MM.YYYY') }}</td>
                <td>{{ event.EndDate | dateParse('YYYY.MM.DD HH:mm:ss') | dateFormat('DD.MM.YYYY') }}</td>
                <td>{{event.LocationName}}</td>
                <td>{{event.RoomName}}</td>
                <td>{{event.TitleName}}</td>
                <td style="font-weight:bold">{{event.TopicName}}</td>
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
                fetch(window.location.origin + "/api/events")
                    .then(response => response.json())
                    .then(data => (this.events = data));
            }
        },
        beforeMount(){
            this.getEvents()
        }
    };
</script>
