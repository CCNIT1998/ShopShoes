<template lang="">
    <div class="product__container product__container-detail">
        <div class="grid wide fix-wide-on-tablet product__container--padding">
            <!-- Nav list pagebreadcrumb> -->
            <pagebreadcrumb />
            <!-- Product product__Content-->
            <productContent :products="products" />
        </div>
    </div>
</template>
<script>
import pagebreadcrumb from "@/pages/product/page-breadcrumb.vue";
import productContent from "@/pages/product/product-content.vue";

export default {
  name: "ProductDetail",
  components: {
    pagebreadcrumb,
    productContent,

  },
  props: ["id"],
  data() {
    return {
        products: [],
    };
  },

  async created() {
      // Dùng fetch
      const response = await fetch(
      `http://127.0.0.1:3000/api/product/${this.id}`
      );
      this.products = await response.json();
      // console.log(a);
  },
};
</script>
<style lang="">
</style>