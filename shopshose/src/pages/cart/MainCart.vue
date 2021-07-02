<template lang="">
    <div class="box">

    </div>
    <header class="container">
        <br />
        <router-link to="/">
            <span class="shopCart" style="font-size:15px">&ensp;Shop&ensp;</span>
        </router-link>
    </header>

    <section class="container">
        <!-- <div v-if="($store.state.cart).length > 0"> -->
            <ul class="products">
                <li class="row" v-for="product in $store?.state.cart" :key="product.ID">
                    <div class="col left">
                        <div class="thumbnail">
                            <a href="#"><img :src=" `http://127.0.0.1:3000/${product.Image}`" style="width:60px; height: 60px;"/></a>
                        </div>
                        <div class="detail">
                            <div class="name">
                                <a href="">{{ product.Name }}</a>
                            </div>
                            <!-- <div class="description">{{ product.description }}</div> -->
                            <div class="price">{{ formatter(product.PriceCurrent) }}</div>
                        </div>
                    </div>
                    <div class="col right">
                        <div class="quantity">
                            <input type="number" class="quantity" v-model="product.quantity" placeholder="1"/>
                        </div>
                        <div class="remove">
                            <i @click="removeItem(index)" class="close far fa-trash-alt"></i>
                        </div>
                    </div>
                </li>
            </ul>
        <!-- </div> -->
        <!-- <div v-else class="empty-product">
            <h3>There are no products in your cart.</h3>
            <button @click="backup">Shopping now</button>
        </div> -->
    </section>
</template>
<script>
    export default {
        name: "MainCart",
        // props: {
        //     products: Object,
        //     promoCode: String,
        // },

        methods: {
            formatter(value) {
                return value.toLocaleString('vi', {style : 'currency', currency : 'VND'});
            },
        
            removeItem(index) {
                this.$emit("deleteItem", index);
            },

            backup() {
                this.$emit("backup")
            }
        },
        computed: {
            
            totalItems() {
                return this.products.reduce((a, b) => a + parseInt(b.quantity), 0)
            },
        }

    };
</script>

<style scoped>
    .shopCart{
        border: 1px solid #ee4d2d;
        border-radius: 6px;
        width: 90px !important;
    }
    .shopCart:hover{
        cursor: pointer;
        color: #f58551;
    }
    img {
        max-width: 100%;
        vertical-align: middle;
        border-radius: 4px;
    }
    .box{
        width: 50px;
        height:120px;
        display: none;
    }
    @media (max-width:1023px){
        .box{

        display: block;
    }
    }
    a {
        text-decoration: none;
        color: #333333;
    }

    a:hover {
        color: #f58551;
    }

    button {
        background-color: #16cc9b;
        border: 2px solid #16cc9b;
        color: #ffffff;
        transition: all 0.25s linear;
        cursor: pointer;
    }

    button::after {
        position: relative;
        right: 0;
        content: " \276f";
        transition: all 0.15s linear;
    }

    button:hover {
        background-color: #f58551;
        border-color: #f58551;
    }

    button:hover::after {
        right: -5px;
    }

    button:focus {
        outline: none;
    }

    ul {
        padding: 0;
        margin: 0;
        list-style-type: none;
    }

    input {
        transition: all 0.25s linear;
    }

    input[type="number"]::-webkit-inner-spin-button,
    input[type="number"]::-webkit-outer-spin-button {
        -webkit-appearance: none;
        -moz-appearance: none;
        appearance: none;
        margin: 0;
    }

    input {
        outline: none;
    }

    .container {
        width: 90%;
        margin: 0 auto;
        overflow: auto;
    }

    /* --- HEADER --- */
    header.container {
        margin-bottom: 1.5rem;
    }

    header .breadcrumb {
        color: #7d7d7d;
    }

    header .breadcrumb li {
        float: left;
        padding: 0 6px;
        height: 20px;
        line-height: 20px;
    }

    header .breadcrumb li:first-child {
        padding-left: 2px;
    }

    header .breadcrumb li:not(:last-child)::after {
        content: " \276f";
        padding-left: 8px;
    }

    header .count {
        font-size: 15px;
        float: right;
        color: #333333;
        height: 20px;
        line-height: 20px;
    }

    /* --- PRODUCT LIST --- */
    .products {
        border-top: 1px solid #ddd;
    }

    .products>li {
        padding: 1rem 0;
        border-bottom: 1px solid #ddd;
    }

    .row {
        position: relative;
        overflow: auto;
        width: 100%;
    }

    .remove {
        height: 100%;
        display: inline-block;
    }

    .col,
    .quantity,
    .remove {
        float: left;
    }

    .col.left {
        width: 30%;
        display: flex;
    }

    .col.right {
        width: 30%;
        position: absolute;
        right: 0;
        top: calc(50% - 30px);
    }

    .detail {
        padding: 0 0.5rem;
        line-height: 2.2rem;
    }

    .detail .name {
        font-size: 1rem;
    }

    .detail .description {
        color: #7d7d7d;
        font-size: 1rem;
    }

    .detail .price {
        font-size: 1.5rem;
    }

    .quantity,
    .remove {
        width: 50%;
        /* text-align: center; */
        right: 0px;
    }

    .remove svg {
        width: 60px;
        height: 60px;
    }

    i.close.far.fa-trash-alt {
        padding-top: 28px;
    }

    .quantity>input {
        display: inline-block;
        width: 60px;
        height: 60px;
        position: relative;
        left: calc(50% - 30px);
        background: #fff;
        border: 2px solid #ddd;
        color: #7f7f7f;
        text-align: center;
        font: 600 1.5rem Helvetica, Arial, sans-serif;
    }

    .quantity>input:hover,
    .quantity>input:focus {
        border-color: #f58551;
    }

    .close {
        fill: #7d7d7d;
        transition: color 150ms linear, background-color 150ms linear,
            fill 150ms linear, 150ms opacity linear;
        cursor: pointer;
    }

    .close:hover {
        fill: #f58551;
    }

    /* --- SUMMARY --- */
    .promotion,
    .summary,
    .checkout {
        float: left;
        width: 100%;
        margin-top: 1.5rem;
    }

    .promotion>label {
        float: left;
        width: 100%;
        margin-bottom: 1rem;
    }

    .promotion>input {
        float: left;
        width: 65%;
        font-size: 1rem;
        padding: 0.5rem 0 0.5rem 1.8rem;
        border: 2px solid #16cc9b;
        border-radius: 2rem 0 0 2rem;
    }

    .promotion:hover>input {
        border-color: #f58551;
    }

    .promotion>button {
        float: left;
        width: 20%;
        height: 2.4rem;
        border-radius: 0 2rem 2rem 0;
    }

    .promotion:hover>button {
        border-color: #f58551;
        background-color: #f58551;
    }

    .promotion>button::after {
        content: "\276f";
        font-size: 1rem;
    }

    .summary {
        font-size: 1.2rem;
        text-align: right;
    }

    .summary ul li {
        padding: 0.5rem 0;
    }

    .summary ul li span {
        display: inline-block;
        width: 30%;
    }

    .summary ul li.total {
        font-weight: bold;
    }

    .checkout {
        text-align: right;
    }

    .checkout>button {
        font-size: 1.2rem;
        padding: 0.8rem 2.8rem;
        border-radius: 1.5rem;
    }

    .empty-product {
        text-align: center;
    }

    .empty-product>button {
        font-size: 1.3rem;
        padding: 10px 30px;
        border-radius: 5px;
    }

    /* --- SMALL SCREEN --- */
    @media all and (max-width: 599px) {
        .thumbnail img {
            display: none;
        }

        .quantity>input {
            width: 40px;
            height: 40px;
            left: calc(50% - 20px);
        }

        .remove svg {
            width: 40px;
            height: 40px;
        }
    }

    /* --- MEDIUM & LARGE SCREEN --- */
    @media all and (min-width: 600px) {
        html {
            font-size: 14px;
        }

        .container {
            width: 75%;
            max-width: 960px;
        }

        .thumbnail,
        .detail {
            float: left;
        }

        .thumbnail {
            width: 20%;
        }

        .detail {
            width: 65%;
        }

        .promotion,
        .summary {
            width: 50%;
        }

        .checkout {
            width: 100%;
        }

        .checkout,
        .summary {
            text-align: right;
        }
    }

    /* --- LARGE SCREEN --- */
    @media all and (min-width: 992px) {
        html {
            font-size: 16px;
        }
    }
</style>