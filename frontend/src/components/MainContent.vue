<script setup>
import LeftSideContent from './main_container/LeftSideContent.vue';
import RightSideContent from './main_container/RightSideContent.vue';
</script>

<template>
    <td style="vertical-align: top; padding: 2px;">
        <table style="width: 100%; border-spacing: 0;">
            <tr>
                <LeftSideContent/>
                <RightSideContent/>
            </tr>
        </table>
    </td>
</template>
