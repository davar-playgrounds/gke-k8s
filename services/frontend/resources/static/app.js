const bus = new Vue();

const fetchWithTimeout = function (url, timeout = 10000, options = {}) {
    return Promise.race([
        fetch(url, options).then(resp => {
            if (resp.status < 200 || resp.status >= 300) {
                return resp.text().then(text => Promise.reject(`${resp.status}: ${text}`));
            }
            return resp;
        }),
        new Promise((_, reject) =>
            setTimeout(() => reject(new Error('timeout')), timeout)
        )
    ]);
};

const countries = Vue.extend({
    template: `#listCountries-template`,
    props: [],
    data() {
        return {
            countries: [],
            searchTimeout: null,
            countrySearch: "",
            message: "",
            selected: "",
            currentPromise: ""
        };
    },
    mounted() {
        this.search();
        bus.$on('countrySelected', (countryCode) => this.selected = countryCode);
    },
    methods: {
        search() {
            this.message = "loading...";
            this.countries = [];
            const thisTime = Date.now();
            this.currentPromise = thisTime;
            fetchWithTimeout(`/countries/search/${this.countrySearch.trim() || ".*"}`, 8000)
                .then(res => res.json())
                .then(response => {
                    if (thisTime === this.currentPromise) {
                        this.message = "";
                        this.countries = response || [];
                    }
                })
                .catch(err => {
                    if (thisTime === this.currentPromise) {
                        this.message = err.toString();
                    }
                });
        },
        onSearch() {
            window.clearTimeout(this.searchTimeout);

            this.searchTimeout = window.setTimeout(this.search, 300);
        },
        selectCountry(e) {
            bus.$emit('countrySelected', e);
        }
    }
});

const airports = Vue.extend({
    template: `#listAirports-template`,
    props: [],
    data() {
        return {
            airports: [],
            searchTimeout: null,
            message: "",
            selected: "",
            airportSearch: "",
            countryCode: "",
            currentPromise: ""
        };
    },
    mounted() {
        bus.$on('countrySelected', (countryCode) => {
            this.selected = "";
            this.airportSearch = "";
            this.countryCode = countryCode;
            this.search();
        });
        bus.$on('airportSelected', (airportIdent) => this.selected = airportIdent);
    },
    methods: {
        search() {
            this.message = "loading...";
            this.airports = [];
            const thisTime = Date.now();
            this.currentPromise = thisTime;
            fetchWithTimeout(`/airports/country_code/${this.countryCode}/search/${this.airportSearch.trim() || ".*"}`, 8000)
                .then(res => res.json())
                .then(response => {
                    if (thisTime === this.currentPromise) {
                        this.message = "";
                        this.airports = response || [];
                    }
                })
                .catch(err => {
                    if (thisTime === this.currentPromise) {
                        this.message = err.toString();
                    }
                });
        },
        onSearch() {
            window.clearTimeout(this.searchTimeout);

            this.searchTimeout = window.setTimeout(() => {
                this.selected = "";
                this.search();
                bus.$emit('airportSearched', {search: this.airportSearch, countryCode: this.countryCode});
            }, 300);
        },
        selectAirport(e) {
            bus.$emit('airportSelected', e);
        }
    }
});

const runways = Vue.extend({
    template: `#listRunways-template`,
    props: [],
    data() {
        return {
            runways: [],
            message: "",
            currentPromise: ""
        };
    },
    mounted() {
        bus.$on('countrySelected', (countryCode) => this.searchByCountryCode(countryCode));
        bus.$on('airportSelected', (airportRef) => this.searchByAirportRef(airportRef));
        bus.$on('airportSearched', ({search, countryCode}) => {
            if (search.trim().length)
                this.searchByCountryCodeAndAirportName(search, countryCode);
            else
                this.searchByCountryCode(countryCode);
        });
    },
    methods: {
        searchByCountryCodeAndAirportName(search, countryCode) {
            this.runways = [];
            this.message = "loading...";
            const thisTime = Date.now();
            this.currentPromise = thisTime;
            fetchWithTimeout(`/runways-country/country_code/${countryCode}/search/${search.trim() || ".*"}`, 8000)
                .then(res => res.json())
                .then(response => {
                    if (thisTime === this.currentPromise) {
                        this.message = "";
                        this.runways = response || [];
                    }
                })
                .catch(err => {
                    if (thisTime === this.currentPromise) {
                        this.message = err.toString();
                    }
                });
        },
        searchByCountryCode(countryCode) {
            this.runways = [];
            this.message = "loading...";
            const thisTime = Date.now();
            this.currentPromise = thisTime;
            fetchWithTimeout(`/runways-country/country_code/${countryCode}`, 8000)
                .then(res => res.json())
                .then(response => {
                    if (thisTime === this.currentPromise) {
                        this.message = "";
                        this.runways = response || [];
                    }
                })
                .catch(err => {
                    if (thisTime === this.currentPromise) {
                        this.message = err.toString();
                    }
                });
        },
        searchByAirportRef(airportRef) {
            this.runways = [];
            this.message = "loading...";
            const thisTime = Date.now();
            this.currentPromise = thisTime;
            fetchWithTimeout(`/runways/airport_ident/${airportRef}`, 8000)
                .then(res => res.json())
                .then(response => {
                    if (thisTime === this.currentPromise) {

                        this.message = "";
                        this.runways = response || [];
                    }
                })
                .catch(err => {
                    if (thisTime === this.currentPromise) {
                        this.message = err.toString();
                    }
                });
        }
    }
});

Vue.component(`listcountries-component`, countries);
Vue.component(`listairports-component`, airports);
Vue.component(`listrunways-component`, runways);

const app = new Vue({
    el: '#vuejs',
    template: `#main-template`,
    data() {
        return {};
    }
});