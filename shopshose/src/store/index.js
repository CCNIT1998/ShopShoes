import {createStore} from "vuex";
import axios from "axios"


const store = createStore({
    state: {
        products: [],
        productsCopy: [],
        cart: [],
        nameSearch: '',
        offset: 0,
    },
    mutations: {
        nextPage(state) {
            state.offset+=10;
        },
        previousPage(state) {
            state.offset-=10;
        },

        setProduct(state, products){
            state.products =  products
            state.productsCopy = products
        },
        setProductCopy(state, products){
            state.productsCopy = products
        },
        filterProduct(state, products){
            state.products =  products
        },

        incrementAmount(state, amount) {
            state.nameSearch == amount;
        },

    },
    actions: {
        async getProducts(context) {
            try {
                const res = await axios.get("http://127.0.0.1:3000/api/product/search/search_items?limit=10&newest="+context.state.offset)
                context.commit("setProduct", res.data['items'])
                // console.log(this.products)
            } catch (e) {
                // console.log(e)
            }
        },
    },

});

export default store;