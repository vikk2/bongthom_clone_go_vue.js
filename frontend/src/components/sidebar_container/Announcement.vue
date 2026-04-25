<script setup>

import { ref, onMounted } from 'vue'

const announcement = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:8080/api/announcement-list')
    announcement.value = await res.json()
})

</script>

<template>
    <tbody id="announcement">
        <tr v-for="announce in announcement" :key="announce.id">
            <td class="red-row3">
                » <a href="#">{{ announce.type }} &nbsp;({{ announce.announcement_num }})</a>
            </td>
        </tr>
    </tbody>
</template>