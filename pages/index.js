function Test() {
    axios.get('/ping', {
    })
        .then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            console.log(error);
        });
}


var app = new Vue({
    el: '#app',
    data: {
        message: 'Hello Vue!'
    },
    methods:{
        createTask: function(){
            Test()
            this.message = "创建成功"
        }
    }
})




