<template>
  <header class="container">
    <h1>Shopping Cart</h1>
    <ul class="breadcrumb">
      <li>Home</li>
      <li>Shopping Cart</li>
    </ul>
    <span class="count">{{totalItems}} items in the bag</span>
  </header>
</template>

<script>
export default {
  name: "HeaderCart",
  props: {
    products: Object,
  },
  computed:{
    totalItems(){
      return this.products.reduce((a,b)=>a + parseInt(b.quantity),0)
    },
  }
};
</script>

<style >
</style>
