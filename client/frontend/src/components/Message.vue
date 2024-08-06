<script setup>
import { computed } from 'vue';
import { useStore } from 'vuex';

const store = useStore();

defineProps({
    content: String,
    from: String,
    created: String,
    color: String,
});

const username = computed(() => store.state.username);
</script>

<template>
    <div :class="from === username ? 'bg-zinc-800' : 'bg-zinc-600'" class="flex justify-between p-4 gap-2 rounded-md">
        <div class="flex gap-[10px]">
            <p :style="{ color }">{{ from }}:</p>
            <p class="text-white">{{ content }}</p>
        </div>
        <span class="sm:block hidden w-fit text-slate-400">{{ created }}</span>
    </div>
</template>