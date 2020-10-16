import Vue from "vue";
import Vuex, { mapState, mapGetters, mapMutations } from "vuex";
import TodoItem from "./todo_item";
import Hello from "./Hello";

Vue.use(Vuex);

const INCREMENT = "INVREMENT";
const DECREMENT = "DECREMENT";

const moduleA = {
  namespaced: true,
  state: () => ({
    value: "this is default value",
  }),
  getters: {
    reverseValue(state) {
      return state.value
        .split("")
        .reverse()
        .join("");
    },
  },
  mutations: {
    setValue: (state, val) => (state.value = val),
  },
  actions: {
    valueChange({ commit }, val) {
      commit("setValue", val);
    },
  },
};

const store = new Vuex.Store({
  modules: {
    moduleA,
  },
  state: {
    count: 0,
    todos: [
      { id: 1, text: "...", done: true },
      { id: 2, text: "...", done: false },
    ],
  },
  getters: {
    doneTodos: (state) => state.todos.filter((todo) => todo.done),
    doneTodoCount: (_, getters) => getters.doneTodos.length,
    getTodoById: (state) => (id) => state.todos.find((todo) => todo.id === id),
  },
  mutations: {
    [INCREMENT](state) {
      state.count++;
    },
    [DECREMENT](state) {
      state.count--;
    },
  },
  actions: {
    incrementAsync({ commit }) {
      setTimeout(() => commit(INCREMENT), 1000);
    },
  },
});

const app = new Vue({
  el: "#app",
  store,
  data: {
    message: "Hello, Vue!",
    greeting: "Hello!",
    name: "Anonymas",
    groceryList: [
      { id: 0, text: "Vegetables" },
      { id: 1, text: "Cheese" },
      { id: 2, text: "Whatever else humans are supposed to eat" },
    ],
    modA: {
      value: "",
    },
  },
  created() {
    this.$data.modA = {
      value: "default",
    };
  },
  methods: {
    reverseMessage() {
      this.message = this.message
        .split("")
        .reverse()
        .join("");
    },
    increment() {
      this.$store.dispatch("incrementAsync");
    },
    ...mapMutations({
      decrement: DECREMENT,
    }),
    valueChange() {
      this.$store.dispatch("moduleA/valueChange", this.modA.value);
    },
  },
  computed: {
    sayHello() {
      return `${this.greeting} ${this.name}!!`;
    },
    ...mapState(["count"]),
    ...mapGetters(["doneTodos", "doneTodoCount", "getTodoById"]),
    ...mapState({
      moduleAValue: (state) => state.moduleA.value,
    }),
    ...mapGetters("moduleA", ["reverseValue"]),
  },
  components: {
    TodoItem,
    Hello,
  },
});
