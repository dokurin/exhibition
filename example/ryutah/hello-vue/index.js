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
});

app.message = "I have changed the data";
app.todos.push({ text: "New Item" });
