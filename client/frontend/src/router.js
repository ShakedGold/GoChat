import {createRouter, createWebHistory} from 'vue-router';

import GeneralView from './views/General.vue';
import AdminView from './views/AdminView.vue';
import RandomView from './views/RandomView.vue';
import HomeView from './views/Home.vue';
import HelpView from './views/HelpView.vue';
import Chat from './components/Chat.vue';

const routes = [
    {
        path: '/:hubName',
        name: 'hub',
        component: Chat,
        props: true,
        meta: {
            isChat: true,
        },
    },
    {
        path: '/random',
        name: 'Random',
        component: RandomView,
    },
    {
        path: '/',
        name: 'Home',
        redirect: '/general',
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;