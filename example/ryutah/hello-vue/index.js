import Vue from "vue";

Vue.component("todo-item", {
  props: ["todo"],
  template: "<li>{{ todo }}</li>",
});

const app = new Vue({
  el: "#app",
  data: {
    message: "Hello, Vue!",
    groceryList: [
      { id: 0, text: "Vegetables" },
      { id: 1, text: "Cheese" },
      { id: 2, text: "Whatever else humans are supposed to eat" },
    ],
  },
  methods: {
    reverseMessage() {
      this.message = this.message.split("").reverse().join("");
    },
  },
});

app.message = "I have changed the data";
