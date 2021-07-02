import { createRouter, createWebHistory } from "vue-router";

import Product from "@/pages/product/index.vue";
import Cart from "@/pages/cart/index.vue";
// import About from "@/views/About.vue";

const routes = [
  {path: "/", component: () => import("@/components/Main.vue") },
  {path: "/:pathMatch(.*)*", component: () => import("@/components/404.vue") },
  {path: "/Product/:id",name: "Product",component: Product, props: true},
  {path: "/Cart", component: Cart},
//   {
//     path: "/about",
//     name: "About",
//     component: About,
//   },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

