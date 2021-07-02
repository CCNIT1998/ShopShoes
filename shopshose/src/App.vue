<template>
<div class="wrapper">
  <Header />
  <router-view v-slot="{ Component }">
    <transition name="fade">
      <component :is="Component" />
    </transition>
  </router-view>
  <Footer />
  
</div>
</template>

<script>
import Header from "@/components/Header.vue";
// import Main from "@/components/Main.vue";
import Footer from "@/components/Footer.vue";
export default {
  name: "App",
  components: {
    Header,
    // Main,
    Footer,
    // ProductDetail,
  },
};
</script>

<style>

</style>
