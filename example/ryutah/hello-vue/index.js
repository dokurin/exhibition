import Vue from "vue";

const app = new Vue({
  el: "#app",
  data: {
    message: "Hello, Vue!",
  },
});

app.message = "I have changed the data";
