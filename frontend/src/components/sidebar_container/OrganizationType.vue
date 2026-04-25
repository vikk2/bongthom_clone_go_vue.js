<script setup>
import { ref, onMounted } from 'vue'

const organizationType = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:8080/api/organization-type')
    organizationType.value = await res.json()
})
</script>

<template>
    <tbody id="organization-type">
        <tr v-for="organization in organizationType" :key="organization.id">
            <td class="red-row3">
                » <a href="#">{{ organization.type }} &nbsp;({{ organization.type_num }})</a>
            </td>
        </tr>
    </tbody>
</template>
