const vm = new Vue({
    el: '#app',
    vuetify: new Vuetify(),
    data: {
        // login
        drawer: null,
        email: "",
        password: "",
        // alert
        type: 'success',
        show: false,
        dismissible: false,
        message: "",
        // rules
        rules: {
            required: v => !!v || 'Required.',
            counter: v => v.length <= 20 || 'Max 20 characters',
            passwordLen: v => (v && v.length >= 6) || 'Password must be more than 6 characters',
            email: v => {
                const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
                return pattern.test(v) || 'Invalid e-mail.'
            },
        },
    },
    methods: {
        signup() {
            axios.post('/signup', {
                email: this.email,
                password: this.password
            }).then(response => {
                this.show = true
                this.type = "success"
                this.message = "created account"
            }).catch(error => {
                console.log(error);
                this.show = true
                this.type = "error"
                this.message = error
            });
        }
    }
});

