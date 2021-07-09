<template>
  <v-app>
    <Header />

    <v-main>
      <Playground />
    </v-main>
  </v-app>
</template>

<script>
import Header from "./components/Header.vue";
import Playground from "./components/Playground.vue";

export default {
  name: "App",

  components: {
    Header,
    Playground,
  },

  data: () => ({
    //
  }),
};
</script>
