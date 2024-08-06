<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();
fetch('http://localhost:8080/hubs')
  .then((res) => res.json())
  .then((data) => {
    // pick random
    const randomRoute = data[Math.floor(Math.random() * data.length)];
    router.push(randomRoute);
  })
  .catch((err) => {
    console.error(err);
  });
</script>

<template>
    <!-- Loading -->
    <div class="bg-gray-700"></div>
</template>