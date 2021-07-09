<template>
  <v-app-bar>
    <v-app-bar-title>
      <v-icon>mdi-elephant</v-icon>
      <v-icon>mdi-database</v-icon>
      Postgres Playground
    </v-app-bar-title>
  </v-app-bar>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Header",
});
</script>
