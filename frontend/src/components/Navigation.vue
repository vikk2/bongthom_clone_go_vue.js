<script setup>
import Register from './navigation_bar/Register.vue';
import JobSearch from './navigation_bar/JobSearch.vue';
import JobList_tab from './navigation_bar/JobList_tab.vue';
import ClassifiedList_tab from './navigation_bar/ClassifiedList_tab.vue';
import ClassifiedSearch_tab from './navigation_bar/ClassifiedSearch_tab.vue';
</script>

<template>
    <tr>
        <td>
            <table style="width: 100%;">
                <tr>
                    <Register/>
                    <JobList_tab/>
                    <JobSearch/>
                    <ClassifiedList_tab/>
                    <ClassifiedSearch_tab/>
                </tr>
            </table>
        </td>
    </tr>
</template>

