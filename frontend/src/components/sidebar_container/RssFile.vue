<script setup>
import { ref, onMounted } from 'vue'

const rssfiles = ref([])
onMounted(async () => {
    const res = await fetch('http://127.0.0.1:8080/api/rssfiles-link')
    rssfiles.value = await res.json()
}) 
</script>

<template>
    <tbody id="rss-files">
        <tr v-for="rss in rssfiles" :key="rss.id">
            <td class="red-row5">
                <img src="../../assets/images/feed-icon.gif" alt="feed-icon">
                <a href="#">{{ rss.rss }}</a>
            </td>
        </tr>
    </tbody>
</template>