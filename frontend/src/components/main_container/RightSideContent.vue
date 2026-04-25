<script setup>
import { ref, onMounted } from 'vue'

const classifieds = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:8080/api/classified-list')
    classifieds.value = await res.json()
})
</script>

<template>
    <td style="width: 50%; padding: 0; vertical-align: top; border-right-width: 3px; color:#800000;">
        <table style="width: 100%; border-spacing: 0;">
            <tr>
                <td style="background-color: #E0E8F8; text-align: center; font-size: medium; color:#800000; font-weight: bolder; font-size: 16px; padding: 0;">
                    Quick View Latest Classifieds
                </td>
            </tr>
            <tr>
                <td style="background-color: #FFE8E8; padding: 0; font-size: 8pt; text-align: center; font-weight: bolder; color: #000000;">
                    Search Classifieds Title
                    <input type="text" size="14" style="font-size: 10pt; color: #000080; border: 1px solid #C0C0C0; height:25px; padding: 0;">
                    <input type="submit" value=">" class="submit_btn">
                </td>
            </tr>
            <tr>
                <td style="background-color: #F3E8FF; padding: 0; font-size: 12px; text-align: center; font-weight: bolder; color: #000000;">
                    <ul id="quick-view-classified">
                        <li v-for="classified in classifieds" :key="classified.id" style="text-align: left; padding-top: 5px;">
                            <a href="#" style="color:#0000FF;">{{ classified.class_title }}</a>
                            <br>
                            <span style="font-style: italic; font-size: 12px; font-weight: normal;">{{ classified.catagory_field }}</span> 
                        </li>    
                    </ul>
                </td>  
            </tr>
            <tr>
                <td style="text-align: center; font-size: 12px; font-weight: bolder; background-color: #F3E8FF; padding: 0 0 12px 0;">
                    <a href="#" style="color:#490106; text-decoration: none;">Quick view all classified ads</a>
                </td>
            </tr>
        </table>
    </td>
</template>