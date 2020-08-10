<template>
    <table style="width:100%">
        <thead>
        <tr>
            <th>Name</th>
            <th>Start Date</th>
            <th>End Date</th>
            <th>Location</th>
            <th>Topics</th>
        </tr>
        </thead>
        <tbody v-for="location in locations" :key="location.ID">
        <tr v-for="event in location.Events" :key="event.ID">
            <td>{{event.EventName}}</td>
            <td>{{event.StartDate}}</td>
            <td>{{event.EndDate}}</td>
            <td>{{location.LocationName}}</td>
            <td>TODO TOPICS</td>
        </tr>
        </tbody>
    </table>
</template>

<script>
    export default {
        props: ['locations'],
        created: function () {
            console.log(this.locations)
        },
    }
</script>


