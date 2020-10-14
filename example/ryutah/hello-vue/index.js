import Vue from "vue";

const app = new Vue({
  el: "#app",
  data: {
    message: "Hello, Vue!",
    todos: [
      { text: "Learn JavaScript" },
      { text: "Learn Vue" },
      { text: "Build something awesome" },
    ],
  },
  methods: {
    reverseMessage() {
      this.message = this.message.split("").reverse().join("");
    },
  },
});

app.message = "I have changed the data";
app.todos.push({ text: "New Item" });
