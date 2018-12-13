const countries = Vue.extend({
    template: `#listCountries-template`,
    props:    [],
    data() {
        return {
            countries: [],
            searchTimeout: null,
            countrySearch: ".*"
        };
    },
    mounted() {
        this.search();
    },
    methods: {
        search() {
            fetch(`/countries/search/${this.countrySearch || ".*"}`)
                .then(res => res.json()) // parse response as JSON (can be res.text() for plain response)
                .then(response => {
                    this.countries = response;
                })
                .catch(err => console.log(err));
        },
        onSearch() {
            window.clearTimeout(this.searchTimeout);

            this.searchTimeout = window.setTimeout(this.search, 300);
        },
        selectCountry(e) {
            console.log(e);
        }
    }
});

Vue.component(`listcountries-component`, countries);

const app = new Vue({
    el: '#vuejs',
    template: `#main-template`,
    data() {
        return {};
    }
});