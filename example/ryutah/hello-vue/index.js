import Vue from "vue";
import Vuex from "vuex";
import TodoItem from "./todo_item";
import Hello from "./Hello";

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    count: 0,
  },
  mutations: {
    increment(state) {
      state.count++;
    },
  },
});

store.commit("increment");
console.log(store.state.count);

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
