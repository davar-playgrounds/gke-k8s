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

const airports = Vue.extend({
    template: `#listAirports-template`,
    props:    [],
    data() {
        return {
            airports: [],
            searchTimeout: null,
            airportSearch: ".*"
        };
    },
    mounted() {
        this.search();
    },
    methods: {
        search() {
            fetch(`/airports/search/${this.airportSearch || ".*"}`)
                .then(res => res.json()) // parse response as JSON (can be res.text() for plain response)
                .then(response => {
                    this.airports = response;
                })
                .catch(err => console.log(err));
        },
        onSearch() {
            window.clearTimeout(this.searchTimeout);

            this.searchTimeout = window.setTimeout(this.search, 300);
        },
        selectAirport(e) {
            console.log(e);
        }
    }
});

Vue.component(`listcountries-component`, countries);
Vue.component(`listairports-component`, airports);

const app = new Vue({
    el: '#vuejs',
    template: `#main-template`,
    data() {
        return {};
    }
});