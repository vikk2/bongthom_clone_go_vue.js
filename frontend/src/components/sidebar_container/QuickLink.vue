<script setup>
// import links from '../../data/link_list.json'
import { ref, onMounted } from 'vue'

const links = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:8080/api/link-list')
    links.value = await res.json()
})
</script>

<template>
    <tbody id="quick-link">
        <tr v-for="link in links" :key="link.id">
            <td class="red-row">
                » <a href="#">{{ link.link_item }}</a> 
                <img v-if="link.green_icon" src="../../assets/images/LPIcon.png" alt="live_music" width="16" style="vertical-align: bottom; margin-left: 3px;"> 
                <span v-if="link.new_label" class="label-new">NEW</span>
            </td>
        </tr>
    </tbody>
</template>
