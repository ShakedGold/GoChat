<script setup>
import { ref } from 'vue';
import HubButton from './HubButton.vue';
const routes = ref([]);
fetch('/hubs')
  .then((res) => res.json())
  .then((data) => {
    routes.value = data;
  })
  .catch((err) => {
    console.error(err);
  });
</script>

<template>
    <div class="bg-zinc-800 flex flex-col gap-8 py-5">
      <h1 class="text-slate-400 text-xl text-center">Hubs</h1>
      <div class="flex flex-col justify-between flex-1">
          <div class="flex flex-col sm:grid-flow-row gap-2 px-5">
            <HubButton v-for="route in routes" :key="route" :hub-name="route" />
          </div>
          <div class="px-5">
            <HubButton hub-name="random" />
          </div>
      </div>
    </div>
</template>