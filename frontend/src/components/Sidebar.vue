<script setup>
import QuickLink from './sidebar_container/QuickLink.vue'
import SpecialSched from './sidebar_container/SpecialSched.vue'
import OrganizationType from './sidebar_container/OrganizationType.vue'
import Announcement from './sidebar_container/Announcement.vue'
import OtherService from './sidebar_container/OtherService.vue'
import RssFile from './sidebar_container/RssFile.vue'
</script>

<template>
    <td style="width: 20%; vertical-align: top;">
        <table style="width: 100%;">
            <tr>
                <td style="padding: 0; line-height: 0;">
                    <img src="../assets/images/secure.gif" alt="secure" width="13" style="vertical-align: top;">
                </td>
            </tr>
            <tr>
                <td style="background-color: #E0E8F8; text-align: center; font-size: medium; color:#800000; font-weight: bolder; font-size: 16px;">
                    Quick Links
                </td>
            </tr>
            <tr>
                <td style="background-color: #FFE8E8; padding: 0; font-size: 8pt; text-align: center; font-weight: bolder;">
                    BongThom-ID: 
                    <input type="text" size="3" style="font-size: 10pt; color: #000080; border: 1px solid #C0C0C0; height:25px; padding: 0;">
                    <input type="submit" value=">" class="submit_btn">
                </td>
            </tr>
            <QuickLink/>
            <tr>
                <td class="red-row2">
                    <img src="../assets/images/facebook.jpg" alt="facebook_icon">
                    <a href="#">See the Bong Thom dot Com Jobs on Facebook</a>
                </td>
            </tr>
            <tr>
                <td style="height: 1px;"></td>
            </tr>
            <tr>
                <td style="background-color: #E0E8F8; text-align: center; font-size: medium; color:#800000; font-weight: bolder; font-size: 16px; padding: 0;">
                    Job Quick Lists
                </td>
            </tr>
            <tr>
                <td style="height: 1px;"></td>
            </tr>
            <tr>
                <td style="background-color: #E0E8F8; text-align: center; font-size: medium; color:#000080; font-weight: bolder; font-size: 8pt; padding: 0;">
                    Special Schedule
                </td>
            </tr>
            <SpecialSched/>
            <tr>
                <td style="height: 1px;"></td>
            </tr>
            <tr>
                <td style="background-color: #E0E8F8; text-align: center; font-size: medium; color:#000080; font-weight: bolder; font-size: 8pt; padding: 0;">
                    Organization Type
                </td>
            </tr>
            <OrganizationType/>
            <tr>
                <td style="height: 1px;"></td>
            </tr>
            <tr>
                <td style="background-color: #E0E8F8; text-align: center; font-size: medium; color:#000080; font-weight: bolder; font-size: 8pt; padding: 0;">
                    Other Announcements
                </td>
            </tr>
            <Announcement/>
            <tr>
                <td style="height: 1px;"></td>
            </tr>
            <tr>
                <td class="red-row3-link">
                    » <a href="#">All job categories...</a>
                </td>
            </tr>
            <tr>
                <td style="height: 14px; padding: 0;"></td>
            </tr>
            <tr>
                <td style="background-color: #E0E8F8; text-align: center; font-size: medium; color:#800000; font-weight: bolder; font-size: 16px; padding: 0;">
                    Other Services
                </td>
            </tr>
            <OtherService/>
            <tr>
                <td class="red-row4" style="text-align: center;">
                    <a href="#">
                        <img src="../assets/images/bongthom.com" alt="bongthom_com">
                    </a>
                </td>
            </tr>
            <tr>
                <td style="height: 14px; padding: 0;"></td>
            </tr>
            <tr>
                <td style="background-color: #E0E8F8; text-align: center; font-size: medium; color:#800000; font-weight: bolder; font-size: 9pt; padding: 0;">
                    RSS Files
                </td>
            </tr>
            <RssFile/>
            <tr>
                <td style="height: 1px;"></td>
            </tr>
            <tr>
                <td class="red-row6">
                    <a href="#">Click here</a> to download an RSS browser
                </td>
            </tr>
            <tr>
                <td style="height: 14px; padding: 0;"></td>
            </tr>
            <tr>
                <td class="red-row3">
                    * NOTE: The numbers in brackets above () are the number of announcements containing job positions in that category
                </td>
            </tr>
        </table>
    </td>
</template>
