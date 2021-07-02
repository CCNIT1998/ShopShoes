<template lang="">
    <header class="header">
        <div class="grid wide fix-wide-on-tablet">
            <!-- Header navbar -->
            <nav class="header__navbar hide-on-mobile-tablet">
                <ul class="header__navbar-list">
                    <li class="header__navbar-item header__navber-item--has-link header__navbar-item--separate">
                        Vào cửa hàng trên ứng dụng CCN Shop
                        <div class="header__qr">
                            <img src="@/assets/img/qr_code.png" alt="QR CODE" class="header__qr-img">
                            <div class="header__qr-apps">
                                <a href="#" class="header__qr-link">
                                    <img src="@/assets/img/google_play.png" alt="Google Play"
                                        class="header__qr-download-img">
                                </a>
                                <a href="#" class="header__qr-link">
                                    <img src="@/assets/img/app_store.png" alt="App Store"
                                        class="header__qr-download-img">
                                </a>
                            </div>
                        </div>
                    </li>
                    <li class="header__navbar-item">
                        <span class="header__navbar-title--no-pointer">
                            Kết nối
                        </span>
                        <a href="#" class="header__navbar-icon-link">
                            <i class="header__navbar-icon fab fa-facebook"></i>
                        </a>
                        <a href="#" class="header__navbar-icon-link">
                            <i class="header__navbar-icon fab fa-instagram"></i>
                        </a>
                    </li>
                </ul>
                <ul class="header__navbar-list">
                    <li class="header__navbar-item header__navber-item--has-notify">
                        <a href="#" class="header__navbar-item-link">
                            <i class="header__navbar-icon far fa-bell"></i>
                            Thông báo
                        </a>


                        <div class="header__notify">
                            <header class="header__notify-header">
                                <h3>Thông báo mới nhận</h3>
                            </header>
                            <ul class="header__notify-list">
                                <li class="header__notify-item" v-for="product in this.$store.state.cart"
                                    :key="product.ID">
                                    <a href="#" class="header__notify-link">
                                        <img :src="`http://127.0.0.1:3000/${product.Image}`" alt=""
                                            class="header__notify-img">
                                        <div class="header__notify-info">
                                            <span class="header__notify-name">{{product.Name}}</span>
                                            <span class="header__notify-descriotion">{{product.Category.Name}}</span>
                                        </div>
                                    </a>
                                </li>

                            </ul>
                            <footer class="header__notify-footer">
                                <a href="#" class="header__notify-footer-btn">
                                    Xem tất cả
                                </a>
                            </footer>
                        </div>
                    </li>
                    <li class="header__navbar-item">
                        <a href="#" class="header__navbar-item-link">
                            <i class="header__navbar-icon far fa-question-circle"></i>
                            Trợ giúp
                        </a>
                    </li>
                    <!-- <li class="header__navbar-item header__navbar-item--bold header__navbar-item--separate">Đăng kí</li>
                <li class="header__navbar-item header__navbar-item--bold">Đăng nhập</li> -->
                    <li class="header__navbar-item header__navbar-user">
                        <img src="https://yt3.ggpht.com/yti/ANoDKi5pv5fU3lXkrswX8iF2gEL5yetOMwkB3vUMdHEHPg=s88-c-k-c0x00ffffff-no-rj-mo"
                            alt="Username" class="header__navbar-user-img">
                        <span class="header__navbar-user-name">UserName</span>

                        <ul class="header__navbar-user-menu">
                            <li class="header__navbar-user-item"><a href="#">Tài khoản của tôi</a></li>
                            <li class="header__navbar-user-item"><a href="#">Địa chỉ của tôi</a></li>
                            <li class="header__navbar-user-item"><a href="#">Đơn mua</a></li>
                            <li class="header__navbar-user-item header__navbar-user-item--separate"><a href="#">Đăng
                                    xuất</a></li>
                        </ul>
                    </li>
                </ul>
            </nav>

            <!-- Header with search -->
            <div class="header-with-search">
                <!-- List Mobile icon -->
                <label for="mobile-list-input" class="header__mobile-list">
                    <i class="header__mobile-list-icon fas fa-bars"></i>
                </label>

                <!-- Search mobile icon -->
                <label for="mobile-search-input" class="header__mobile-search">
                    <i class="header__mobile-search-icon fas fa-search"></i>
                </label>

                <!-- Logo -->
                <div class="header__logo">
                    <router-link to="/" class="header__logo-link">
                        <img src="https://via.placeholder.com/60x60" alt="" srcset="" />
                    </router-link>
                </div>

                <input type="checkbox" name="" id="mobile-search-input" class="header__search-checkbox" hidden>

                <div class="header__search">
                    <div class="header__search-input-wrap">
                        <input v-model="amount" type="text" class="header__search-input"
                            placeholder="Tìm kiếm sản phẩm">

                        <!-- Search history -->
                        <div class="header__search-history">
                            <h3 class="header__search-history-heading">Lịch sử tìm kiếm</h3>
                            <ul class="header__search-history-list">
                                <li class="header__search-history-item">
                                    <a href="#">Giày thể thao nam</a>
                                </li>
                                <li class="header__search-history-item">
                                    <a href="#">Giày thể thao nữ</a>
                                </li>
                                <li class="header__search-history-item">
                                    <a href="#">Giày da</a>
                                </li>
                            </ul>
                        </div>
                    </div>
                    <div class="header__search-select">
                        <i class="fas fa-camera"></i>
                    </div>
                    <button @click="btnNameSearch" class="header__search-btn">
                        <i class="header__search-btn-icon fas fa-search"></i>
                    </button>
                </div>

                <label for="header__mobile-notify" class="header__notify-mobile">
                    <i class="header__notify-mobile--icon far fa-bell"></i>
                    <span class="header__notify-mobile--notice">5+</span>
                </label>

                <input type="checkbox" hidden name="" id="header__mobile-notify" class="header__mobile-noti">
                <label for="header__mobile-notify" class="header__mobile-notify--overlay"></label>
                <div class="header__notify">
                    <header class="header__notify-header">
                        <h3>Thông báo mới nhận</h3>
                    </header>
                    <ul class="header__notify-list">
                        <li class="header__notify-item" v-for="product in this.$store.state.cart" :key="product.ID">
                            <a href="#" class="header__notify-link">
                                <img :src="`http://127.0.0.1:3000/${product.Image}`" alt="" class="header__notify-img">
                                <div class="header__notify-info">
                                    <span class="header__notify-name">{{product.Name}}</span>
                                    <span class="header__notify-descriotion">{{product.Category.Name}}</span>
                                </div>
                            </a>
                        </li>
                    </ul>
                    <footer class="header__notify-footer">
                        <a href="#" class="header__notify-footer-btn">
                            Xem tất cả
                        </a>
                    </footer>
                </div>

                <!-- Cart layout -->
                <div class="header__cart">
                    <div class="header__cart-wrap">
                        <i class="header__cart-icon fas fa-shopping-cart"></i>
                        <span class="header__cart-notice">{{ cartNotice }}</span>

                        <!-- No cart: header__cart-list--no-cart -->
                        <div class="header__cart-list">
                            <img src="@/assets/img/no-cart.png" alt="" class="header__cart-no-cart-img">
                            <span class="header__cart-list--no-cart-msg">Chưa có sản phẩm</span>

                            <h4 class="header__cart-heading">Sản phẩm đã thêm</h4>
                            <ul class="header__cart-list-item">
                                <li class="header__cart-item" v-for="product in this.$store.state.cart"
                                    :key="product.ID">
                                    <img :src="`http://127.0.0.1:3000/${product.Image}`" alt="Bộ kem sáng da mềm mịn"
                                        class="header__cart-img">
                                    <div class="header__cart-item-info">
                                        <div class="header__cart-item-head">
                                            <h5 class="header__cart-item-name">{{product.Name}}</h5>
                                            <div class="header__cart-item-price-wrap">
                                                <span class="header__cart-item-price">{{ formatter(product.PriceCurrent) }}</span>
                                                <span class="header__cart-item-multiply">x</span>
                                                <span class="header__cart-item-qnt">1</span>
                                            </div>
                                        </div>

                                        <div class="header__cart-item-body">
                                            <!-- <span class="header__cart-item-description">
                                            Phân chia loại sản phẩm Phân chia loại sản phẩm Phân chia loại sản phẩm
                                        </span> -->
                                            <span class="header__cart-item-del">
                                                Xóa
                                            </span>
                                        </div>
                                    </div>
                                </li>
                            </ul>
                            <div class="header__cart-view-cart">
                                <router-link to="/Cart" class="header__cart-btn btn btn--primary">Xem giỏ hàng
                                </router-link>
                            </div>

                        </div>
                    </div>
                </div>
            </div>
            <input type="checkbox" hidden class="nav__input" id="mobile-list-input">

            <label for="mobile-list-input" class="nav__overlay"></label>

            <nav class="header__mobile-nav">
                <div class="header__mobile-nav--heading">
                    <img src="https://yt3.ggpht.com/yti/ANoDKi5pv5fU3lXkrswX8iF2gEL5yetOMwkB3vUMdHEHPg=s88-c-k-c0x00ffffff-no-rj-mo"
                        alt="UserName" class="header__mobile-nav--img">
                    <div class="header__mobile-nav--text">
                        <h2 class="header__mobile-name">UserName</h2>
                        <span class="header__mobile-rule">
                            <i class="header__mobile-rule--icon fas fa-user-shield"></i>
                            User
                        </span>
                    </div>
                    <label for="mobile-list-input" class="header__mobile-close">
                        <i class="header__mobile-close-icon fas fa-times"></i>
                    </label>
                </div>
                <ul class="header__mobile-nav--list">
                    <li>
                        <a href="#" class="header__mobile-link">
                            <i class="fas fa-home"></i>
                            Trang chủ
                        </a>
                    </li>
                    <li>
                        <a href="#" class="header__mobile-link">
                            <i class="fas fa-id-card"></i>
                            Thông tin tài khoản
                        </a>
                    </li>
                    <li>
                        <a href="#" class="header__mobile-link">
                            <i class="fas fa-box"></i>
                            Quản lí đơn hàng
                        </a>
                    </li>
                    <li>
                        <a href="#" class="header__mobile-link">
                            <i class="fas fa-map-marker-alt"></i>
                            Địa chỉ của tôi
                        </a>
                    </li>
                    <li>
                        <a href="#" class="header__mobile-link">
                            <i class="fas fa-sign-out-alt"></i>
                            Đăng xuất
                        </a>
                    </li>
                </ul>
            </nav>
        </div>

        <ul class="header__sort-bar">
            <li class="header__sort-item">
                <a href="#" class="header__sort-link">Liên quan</a>
            </li>
            <li class="header__sort-item header__sort-item--active">
                <a href="#" class="header__sort-link">Mới nhất</a>
            </li>
            <li class="header__sort-item">
                <a href="#" class="header__sort-link">Bán chạy</a>
            </li>
            <li class="header__sort-item">
                <a href="#" class="header__sort-link">Giá</a>
            </li>
        </ul>
    </header>
</template>
<script>
    export default {
        name: "Header",
        data() {
            return {
                amount: "AD",
            };
        },

        computed: {
            cartNotice() {
                return this.$store.state.cart.length
            },
            
        },

        methods: {
            formatter(value) {
                return value.toLocaleString('vi', {style : 'currency', currency : 'VND'});
            },
            btnNameSearch() {
                let temp = this.$store.state.productsCopy
                this.$store.commit("filterProduct", temp);
                this.$store.commit("incrementAmount", this.amount);
                return this.$store.state.products
            },
        }
    }
</script>
<style lang="">

</style>