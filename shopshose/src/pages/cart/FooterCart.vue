<template lang="">
  <section class="container">
    <div class="promotion">
      <label for="promo-code">Nhập mã giảm giá</label>
      <!-- v-model="promoCode" -->
      <input v-model="promoCode.code" />
      <button type="button" @click="submitPromoCode()"></button>
    </div>
    <div class="summary">
      <ul>
        <li>Tổng tiền sản phẩm <span>{{ formatter(subTotal) }}</span></li>
        <li v-if="promoCode[0].discount > 0">Phiếu giảm giá <span>{{ formatter(subTotal *promoCode[0].discount*0.01) }}</span>
        </li>
        <li>Thuế <span>{{ formatter(Tax) }}</span></li>
        <li class="total">Tổng thanh toán <span>{{ formatter(Total) }}</span></li>

      </ul>
    </div>
    <div class="checkout">
      <button type="button">Mua ngay</button>
    </div>
    <br /><br />
    
  </section>
</template>
<script>
  export default {

    name: "FooterCart",
    data() {
      return {
        promotions: [
          {
            code: "PROMO1",
            discount: "30%"
          },
          {
            code: "PROMO2",
            discount: "20%"
          },
          {
            code: "PROMO3",
            discount: "10%"
          }
        ],

        promoCode: [
          {
            code: "abc",
            discount: 0
          }
        ],
      }
    },

    props: {
      products: Object,
    },

    computed: {
      subTotal() {
        return this.products.reduce((a, b) => (a + b.PriceCurrent), 0)
      },

      Tax() {
        return parseFloat(this.subTotal * 0.1)
      },

      Total() {
        let discount = this.promoCode[0].discount * 0.01 * this.subTotal
        return this.subTotal + this.Tax - discount
      },

    },

    methods: {
      formatter(value) {
        return value.toLocaleString('vi', {style : 'currency', currency : 'VND'});
      },

      submitPromoCode() {
        var checkpromoCode = 0
        var checkDiscount = 0
        for (var i = 0; i < this.promotions.length; i++) {
          if (this.promoCode.code === this.promotions[i].code) {
            checkpromoCode++;
            checkDiscount = parseFloat(this.promotions[i].discount.replace("%", ""))
          }
        }
        if (checkpromoCode > 0) {
          this.promoCode[0].discount = checkDiscount
        } else {
          this.promoCode[0].discount = 0
        }

      },

    }
  };
</script>

<style scoped>
  img {
    max-width: 100%;
    vertical-align: middle;
    border-radius: 4px;
  }

  .checkout{
    margin-bottom: 10px !important; 
  }

  a {
    text-decoration: none;
    color: #333333;
  }

  a:hover {
    color: #f58551;
  }

  button {
    background-color: #ee4d2d;
    border: 2px solid #ee4d2d;
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
    width: 70%;
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
    text-align: center;
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
    border: 2px solid #ee4d2d;
    border-radius: 2rem 0 0 2rem;
  }

  .promotion:hover>input {
    border-color: #f58551;
  }

  .promotion>button {
    float: left;
    width: 20%;
    height: 2.45rem;
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
      width: 35%;
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