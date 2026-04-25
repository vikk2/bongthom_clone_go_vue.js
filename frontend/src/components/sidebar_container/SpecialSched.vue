<script setup>
import { ref, onMounted } from 'vue'

const specialSched = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:8080/api/specialSched-link')
    specialSched.value = await res.json()
})
</script>

<template>
    <tbody id="special-schedule">
        <tr v-for="schedule in specialSched" :key="schedule.id">
            <td class="red-row3">
                » <a href="#">{{ schedule.sched_title }} &nbsp;({{ schedule.sched_num }})</a>
            </td>
        </tr>
    </tbody>
</template>