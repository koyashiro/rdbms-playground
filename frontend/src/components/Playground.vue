<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <p>Query</p>
        <form>
          <v-container>
            <textarea v-model="query" />
          </v-container>
          <v-btn color="primary">Execute</v-btn>
        </form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "Playground",
  setup() {
    const query = ref("");
    const execute = () => {};

    return { query, execute };
  },
});
</script>

<style scoped>
textarea {
  border: 1px solid black;
  border-radius: 4px;
}
</style>
