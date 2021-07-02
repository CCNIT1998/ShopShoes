<template>
  <!-- <HeaderCart :products="products" /> -->
  <MainCart :products="products" @deleteItem="deleteItem" @backup="backup" />

  <FooterCart :products="products" :promoCode="promoCode" />
</template>

<script>
// import HeaderCart from "@/pages/cart/HeaderCart.vue";
import MainCart from "@/pages/cart/MainCart.vue";
import FooterCart from "@/pages/cart/FooterCart.vue";
export default {
  name: "Cart",
  components: {
    // HeaderCart,
    MainCart,
    FooterCart,
  },

  data() {
    return {
      products: [],

      backupProducts: [],
    };
  },
  computed: {
    
  },
  mounted(){
    this.products = this.$store.state.cart
  },

  created() {
    this.backupProducts = [...this.products];
  },

  methods: {

    deleteItem(index) {
      if (confirm("Bạn có chắc chắn xoá sản phẩm này!")) {
        this.products.splice(index, 1);
      }
    },

    backup() {
      this.products = [...this.backupProducts];
    },

    checkPromoCode() {
      for (let i = 0; i < this.promotions.length; i++) {
        if (this.promoCode === this.promotions[i].code) {
          this.discount = parseFloat(
            this.promotions[i].discount.replace("%", "")
          );
          console.log(this.discount);
          return;
        }
      }
    },
  },
};
</script>

