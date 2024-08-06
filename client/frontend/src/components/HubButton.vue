<script setup>
import { useRoute } from 'vue-router';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faAngleRight } from '@fortawesome/free-solid-svg-icons';
import { computed } from 'vue';

const route = useRoute();

const props = defineProps({
    hubName: String,
});

const currentRouteName = computed(() => route.params.hubName);

const capitalize = (str) => {
    return str.charAt(0).toUpperCase() + str.slice(1);
};
</script>

<template>
    <router-link :to="hubName">
        <div :class="currentRouteName === hubName ? 'bg-slate-400' : 'hover:bg-zinc-900'" class="sm:p-6 text-white flex items-center justify-between rounded-md">
            <p :class="currentRouteName === hubName ? 'text-zinc-900' : 'text-slate-400'">{{ capitalize(hubName) }}</p>
            <font-awesome-icon :icon="faAngleRight" :class="currentRouteName === hubName ? 'text-zinc-900' : 'text-slate-400'" />
        </div>
    </router-link>
</template>