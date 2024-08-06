<script setup>
import { computed, nextTick, onBeforeUnmount, onMounted, onUpdated, ref, watch } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from 'vuex';
import Message from './Message.vue';

const store = useStore();
const route = useRoute();

const props = defineProps({
  hubName: String,
});

const username = ref(store.state.username);
const messages = ref([]);
const message = ref('');

// connect to localhost:8888 via sockets
const socket = computed(() => {
  messages.value = [];
  const s = new WebSocket(`ws://${location.host}/ws/${props.hubName}`);
  s.onmessage = (event) => {
    const message = JSON.parse(event.data);
    messages.value.push(message);
    return false;
  };
  return s;
});

const sendMessage = () => {
  // check if socket is open
  if (socket.value.readyState !== WebSocket.OPEN) {
    console.error('Socket is not open');
    return;
  }
  if (message.value === '') {
    return;
  }
  socket.value.send(JSON.stringify({
    from: store.state.username,
    content: message.value,
    created: '',
  }));
  message.value = '';
};

const capitalize = (str) => {
  return str.charAt(0).toUpperCase() + str.slice(1);
};

watch(socket, (newSocket, prevSocket) => {
  prevSocket?.close();
}, { immediate: true });
</script>

<template>
  <div class="flex flex-col gap-[50px] px-10 bg-zinc-700 flex-1 h-full sm:max-h-svh max-h-[calc(100vh-40px)]">
    <div class="text-white flex items-center justify-between flex-wrap">
      <h1 class="text-2xl py-5 text-slate-400">{{ capitalize(hubName) }}</h1>
      <div class="flex items-center gap-4 flex-wrap">
        <input class="w-fit border border-gray-900 bg-zinc-800 focus:bg-zinc-900 rounded-md p-2 outline-none" v-model="username" @input="() => store.commit('setUsername', username)" type="text" name="username" id="username">
      </div>
    </div>
    <div class="flex flex-col gap-[10px] flex-1 overflow-auto">
      <div v-if="messages.length > 0" class="flex flex-col gap-4">
        <Message v-for="m in messages" :key="m" :content="m.content" :from="m.from" :created="m.created" :color="m.color" />
      </div>
      <div v-else class="flex flex-col gap-4">
        <div class="text-slate-200 text-opacity-40 text-center text-2xl">No messages yet...</div>
      </div>
    </div>
    <div class="flex mt-auto pb-5 gap-[20px]">
      <input class="flex-1 border border-gray-900 bg-zinc-800 focus:bg-zinc-900 rounded-md p-4 outline-none text-white" placeholder="Send a message..." v-model="message" @keyup.enter="sendMessage" type="text">
      <button class="border-gray-900 bg-zinc-800 active:bg-zinc-900 rounded-md p-4 outline-none text-white" @click="sendMessage">Send</button>
    </div>
  </div>
</template>

<style>
/* Tooltip container */
.tooltip {
  position: relative;
  display: inline-block;
}

/* Tooltip text */
.tooltip .tooltiptext {
  visibility: hidden;
  background-color: black;
  color: #fff;
  text-align: center;
  border-radius: 6px;
 
  /* Position the tooltip text - see examples below! */
  position: absolute;
  z-index: 1;
  bottom: 100%;
}

/* Show the tooltip text when you mouse over the tooltip container */
.tooltip:hover .tooltiptext {
  visibility: visible;
}
</style>