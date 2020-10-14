import Vue from "vue";
import TodoItem from "./todo_item";
import Hello from "./Hello";

const app = new Vue({
  el: "#app",
  data: {
    message: "Hello, Vue!",
    greeting: "Hello!",
    name: "Anonymas",
    groceryList: [
      { id: 0, text: "Vegetables" },
      { id: 1, text: "Cheese" },
      { id: 2, text: "Whatever else humans are supposed to eat" },
    ],
  },
  methods: {
    reverseMessage() {
      this.message = this.message
        .split("")
        .reverse()
        .join("");
    },
  },
  computed: {
    sayHello() {
      return `${this.greeting} ${this.name}!!`;
    },
  },
  components: {
    TodoItem,
    Hello,
  },
});
