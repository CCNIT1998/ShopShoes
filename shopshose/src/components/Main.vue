<template lang="">
    <Slide />
    <div class="app__container">
        <div class="grid wide fix-wide-on-tablet">
            <div class="row app__content sm-gutter">
                <!-- List -->
                <div class="col p-2 t-3 m-0">
                    <nav class="category">
                        <h3 class="category__heading">Danh mục</h3>

                        <ul class="category-list">
                            <router-link to="/" @click=categoryShoseSportBoy>
                                <li class="category-item category-item--active">
                                    <a class="category-item__link">Giày thể thao nam</a>
                                </li>
                            </router-link>
                            <router-link to="/" @click=categoryShoseSportGirl>
                                <li class="category-item">
                                    <a class="category-item__link">Giày thể thao nữ</a>
                                </li>
                            </router-link>
                            <router-link to="/" @click=categoryShoseLeatherBoy>
                                <li class="category-item">
                                    <a class="category-item__link">Giày da nam</a>
                                </li>
                            </router-link>
                            <router-link to="/" @click=categoryShoseLeatherGirl>
                                <li class="category-item">
                                    <a class="category-item__link">Giày da nữ</a>
                                </li>
                            </router-link>
                        </ul>

                        <ul class="category-list">
                            <li><img src="https://via.placeholder.com/190x515" alt=""></li>
                        </ul>
                    </nav>
                </div>

                <div class="col p-10 t-9 m-12">
                    <!-- Filter -->
                    <div class="home-filter hide-on-mobile-tablet">
                        <span class="home-filter__label">Sắp xếp theo</span>
                        <button @click="shoesPopular" class="home-filter__btn btn">Phổ biến</button>
                        <button @click="shoesTopTraning" class="home-filter__btn btn btn--primary">Yêu Thích</button>
                        <button @click="shoesTopSell" class="home-filter__btn btn">Bán chạy</button>

                        <!-- sort Price -->
                        <div class="select-input">
                            <span class="select-input__label">Giá</span>
                            <i class="select-input__icon fas fa-angle-down"></i>

                            <ul class="select-input__list">
                                <li class="select-input__item">
                                    <router-link to="/">
                                        <a href="" @click="sortPriceDESC" class="select-input__link">Giá: Thấp đến cao
                                        </a>
                                    </router-link>
                                </li>
                                <li class="select-input__item">
                                    <router-link to="/">
                                        <a href="" @click="sortPriceASC" class="select-input__link">Giá: Cao đến thấp
                                        </a>
                                    </router-link>


                                </li>
                            </ul>
                        </div>

                        <div class="home-filter__page">
                            <span class="home-filter__page-num">
                                <span class="home-filter__page-current">1</span>/4
                            </span>

                            <div class="home-filter__page-control">
                                <a href="#" class="home-filter__page-btn home-filter__page-btn--disabled">
                                    <i class="home-filter__page-icon fas fa-angle-left"></i>
                                </a>
                                <a href="#" class="home-filter__page-btn">
                                    <i class="home-filter__page-icon fas fa-angle-right"></i>
                                </a>
                            </div>
                        </div>
                    </div>

                    <!-- Nav list mobile -->
                    <nav class="mobile-category">
                        <u class="mobile-category__list">
                            <li class="mobile-category__item">
                                <a href="#" class="mobile-category__link">Giày thể thao nam</a>
                            </li>
                            <li class="mobile-category__item">
                                <a href="#" class="mobile-category__link">Giày thể thao nữ</a>
                            </li>

                        </u>
                    </nav>

                    <!-- Product -->
                    <div class="home-product">

                        <!-- grid->row->column -->
                        <div class="row sm-gutter">
                            <!-- Product Item -->
                            <div class="col p-2-4 t-4 m-6" v-for="product in this.$store.state.products" :key="product.ID">
                                <router-link :to="{ name: 'Product', params: { id: product.ID } }">
                                    <a class="home-product-item">
                                        <div class="home-product-item__img hov-img-zoom"
                                            :style="{backgroundImage: `url(http://127.0.0.1:3000/${product.Image})` }">
                                        </div>

                                        <h4 class="home-product-item__name">
                                            {{ product.Name}}
                                        </h4>


                                        <div class="home-product-item__price">
                                            <span class="home-product-item__price-old">{{ formatter(product.PriceOld) }}</span>
                                            <span class="home-product-item__price-current">{{ formatter(product.PriceCurrent)
                                                }}</span>
                                        </div>

                                        <div class="home-product-item__action">
                                            <!-- Liked: home-product-item__like--liked -->
                                            <span class="home-product-item__like home-product-item__like--liked">
                                                <i class="home-product-item__like-icon-empty far fa-heart"></i>
                                                <i class="home-product-item__like-icon-fill fas fa-heart"></i>
                                            </span>

                                            <div v-if="product.RatingStar > 4" class="home-product-item__rating">
                                                <i class="home-product-item__star--gold fas fa-star"
                                                    v-for="start in RatingStar(product.RatingStar)" :key="start.ID"></i>
                                            </div>
                                            <div v-else class="home-product-item__rating">
                                                <i class="home-product-item__star--gold fas fa-star"
                                                    v-for="start in RatingStar(product.RatingStar)" :key="start.ID"></i>
                                                <i class="fas fa-star"></i>
                                            </div>
                                            <span class="home-product-item__sold">{{product.HistoricalSold}} đã
                                                bán</span>
                                        </div>

                                        <div class="home-product-item__origin">
                                            <span class="home-product-item__brand">Whoo</span>
                                            <span class="home-product-item__origin-name">{{product.Madein}}</span>
                                        </div>

                                        <div class="home-product-item__favorite">
                                            <i class="fas fa-check"></i>
                                            <span>Yêu thích</span>
                                        </div>

                                        <div class="home-product-item__sale-off">
                                            <span class="home-product-item__sale-off-percent">{{product.Sale}}%</span>
                                            <span class="home-product-item__sale-off-label">GIẢM</span>
                                        </div>
                                    </a>
                                </router-link>
                            </div>


                        </div>

                    </div>

                    <ul class="pagination pagination__home-product">
                        <router-link to="/">
                        <li class="pagination-item" @click="previousPage">
                            <a href="#" class="pagination-item__link">
                                <i class="pagination-item__icon fas fa-angle-left"></i>
                            </a>
                        </li>
                        </router-link>

                        <li class="pagination-item" v-for="(page, index) in totalPage" :key="index">
                            <a href="#" @click="paging(index)" class="pagination-item__link">{{"index + 1"}}</a>
                        </li>
                        <li class="pagination-item pagination-item--active">
                            <a href="#" class="pagination-item__link">
                                1
                            </a>
                        </li>
                        <li class="pagination-item">
                            <a href="#" class="pagination-item__link">
                                2
                            </a>
                        </li>
                        <router-link to="/">
                        <li class="pagination-item" @click="nextPage">
                            <a href="#" class="pagination-item__link">
                                <i class="pagination-item__icon fas fa-angle-right"></i>
                            </a>
                        </li>
                        </router-link>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
    import Slide from "@/components/Slide.vue";

    export default {
        name: "Main",
        components: {
            Slide,
        },

        data() {
            return {
                listProducts: [],
            };
        },
        computed: {

            products() {
                return this.$store.state.products
            },

            shoesPopular() {
                let temp = this.$store.state.productsCopy
                this.$store.commit("filterProduct", temp);

                return this.$store.state.products == temp
            },

            shoesTopSell() {
                let temp = this.$store.state.productsCopy
                this.$store.commit("filterProduct", temp);
                let tempProduct = this.$store.state.products.filter(product => product.HistoricalSold >= 1000)
                this.$store.commit("filterProduct", tempProduct);
                return this.$store.state.products
            },

            shoesTopTraning() {
                let temp = this.$store.state.productsCopy
                this.$store.commit("filterProduct", temp);
                let tempProduct = this.$store.state.products.filter(product => product.RatingStar == 5)
                this.$store.commit("filterProduct", tempProduct);
                return this.$store.state.products
            },

        },

        methods: {
            formatter(value) {
                return value.toLocaleString('vi', {style : 'currency', currency : 'VND'});
            },

            nextPage() {
                this.$store.commit("nextPage");
                this.$store.dispatch("getProducts")
            },
            previousPage() {
                this.$store.commit("previousPage");
                this.$store.dispatch("getProducts")
            },

            search() {
                let temp = this.$store.state.productsCopy
                this.$store.commit("filterProduct", temp);
                return this.$store.state.products.filter((product) =>
                    product.Name.toLowerCase().includes(this.$store.state.nameSearch.toLowerCase())
                );
            },

            categoryShoseSportBoy() {
                let temp = this.$store.state.productsCopy
                this.$store.commit("filterProduct", temp);
                let tempProduct = this.$store.state.products.filter(product => product.CategoryID == 2)
                this.$store.commit("filterProduct", tempProduct);
            },
            categoryShoseSportGirl() {
                let temp = this.$store.state.productsCopy
                this.$store.commit("filterProduct", temp);
                let tempProduct = this.$store.state.products.filter(product => product.CategoryID == 3)
                this.$store.commit("filterProduct", tempProduct);
            },
            categoryShoseLeatherBoy() {
                let temp = this.$store.state.productsCopy
                this.$store.commit("filterProduct", temp);
                let tempProduct = this.$store.state.products.filter(product => product.CategoryID == 4)
                this.$store.commit("filterProduct", tempProduct);
            },
            categoryShoseLeatherGirl() {
                let temp = this.$store.state.productsCopy
                this.$store.commit("filterProduct", temp);
                let tempProduct = this.$store.state.products.filter(product => product.CategoryID == 1)
                this.$store.commit("filterProduct", tempProduct);
            },


            RatingStar(product) {
                let temp = []
                for (let i = 0; i < product; i++) {
                    temp.push(i)
                }
                return temp
            },

            sortPriceASC() {
                this.$store.state.products.sort(function (a, b) {
                    return a.PriceCurrent - b.PriceCurrent;
                })
            },
            sortPriceDESC() {
                this.$store.state.products.sort(function (a, b) {
                    return b.PriceCurrent - a.PriceCurrent;
                })
            },

        },

        mounted() {
            this.$store.dispatch("getProducts")
        },
        
    };
</script>
<style lang="">
</style>