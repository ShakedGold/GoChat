<script setup>
import { ref } from 'vue';
import store from './store';
import Sidebar from './components/Sidebar.vue';

const username = ref(`anon-${Math.floor(Math.random() * 1000)}`);
store.commit('setUsername', username.value);
</script>

<template>
  <div class="flex flex-col flex-1 h-svh">
    <div class="flex sm:flex-row flex-col flex-1">
      <Sidebar class="w-[350px]" />
      <router-view class="flex-1"></router-view>
    </div>
  </div>
</template>