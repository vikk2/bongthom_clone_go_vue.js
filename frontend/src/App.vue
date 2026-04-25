<script setup>
import Header from '@/components/Header.vue'
import Language_row from '@/components/Language_row.vue'
import Navigation from '@/components/Navigation.vue'
import Sidebar from '@/components/Sidebar.vue'
import MainContent from '@/components/MainContent.vue'
import Footer from '@/components/Footer.vue'
</script>

<template>
  <table class="big_table">
    <Header/>
    <Language_row/>
    <tr>
        <td style="height: 18px;">

        </td>
    </tr>
    <Navigation/>
    <tr>
      <td style="padding: 0;">
          <table style="width: 100%; border-spacing: 0;">
              <tr>
                <Sidebar/>
                <MainContent/>
              </tr>
            </table>
        </td>
    </tr>
  </table>
  <Footer/>

</template>
