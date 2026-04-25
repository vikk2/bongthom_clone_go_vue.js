<script setup>
import { ref, onMounted } from 'vue'

const services = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:8080/api/other-service-link')
    services.value = await res.json()
})
</script>

<template>
    <tbody id="other-service">
        <tr v-for="service in services" :key="service.id">
            <td class="red-row4">
                » <a href="#">{{ service.service }}</a>
            </td>
        </tr>
    </tbody>
</template>