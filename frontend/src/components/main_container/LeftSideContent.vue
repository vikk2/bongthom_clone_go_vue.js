<script setup>
import { ref, onMounted } from 'vue'

const jobs = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:8080/api/job-link')
    jobs.value = await res.json()
})
</script>

<template>
    <td style="width: 50%; padding: 0; vertical-align: top;">
        <table style="width: 100%; border-spacing: 0; border-collapse: collapse;">
            <tr>
                <td style="background-color: #E0E8F8; text-align: center; font-size: medium; color:#800000; font-weight: bolder; font-size: 16px; padding: 0; border-right-style: dotted; border-right-width: 3px; color:#800000;">
                    Quick View Latest Jobs <span style="font-size: 8pt; font-weight: normal;">( <img src="../../assets/images/new.gif" alt="new_icon"> = New )</span>
                </td>
            </tr>
            <tr>
                <td style="background-color: #FFE8E8; padding: 0; font-size: 8pt; text-align: center; font-weight: bolder; color: #000000; border-right-style: dotted; border-right-width: 3px; color:#000000;">
                    Search Job Title
                    <input type="text" size="14" style="font-size: 10pt; color: #000080; border: 1px solid #C0C0C0; height:25px; padding: 0;">
                    <input type="submit" value=">" class="submit_btn">
                </td>
            </tr>
            <tr>
                <td style="background-color: #EFF7F8; padding: 0; font-size: 12px; text-align: center; font-weight: bolder; color: #000000; border-right-style: dotted; border-right-width: 3px; color:#000000;">
                    <ul style="margin-top: 0;" id="quick-job-list">
                        <li v-for="job in jobs" :key="job.id" style="text-align: left;">
                            <a href="#" style="color:#0000FF">
                                <img v-if="job.new" src="../../assets/images/new.gif" alt="blue-icon">
                                <img v-else="job.new" src="../../assets/images/static.png" alt="blue-icon">
                                &nbsp;&nbsp;{{ job.title }}
                            </a><br>
                            <span style="font-style: italic; font-size: 12px; font-weight: normal;">{{ job.company }}</span>&nbsp;
                            <img src="../../assets/images/companydetails.png" alt="arrow_icon" style="vertical-align: bottom;">
                        </li>
                    </ul>
                </td>
            </tr>
            <tr>
                <td style="text-align: center; font-size: 12px; font-weight: bolder; background-color: #EFF7F8; padding: 0 0 12px 0; border-right-style: dotted; border-width: 3px;">
                    <a href="#" style="color:#490106; text-decoration: none;">Quick view all active jobs</a>
                </td>
            </tr>
        </table>
    </td>
</template>