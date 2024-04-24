import { createRouter, createWebHistory } from 'vue-router';
import HomePage from './views/Home-Page.vue';
import LoginPage from './views/Login-Page.vue';
import AboutPage from './views/About-Page.vue';
import AdminPage from './views/AdministratorPrincipalPage.vue';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: HomePage
    },
    {
        path: '/login',
        name: 'Login',
        component: LoginPage
    },
    {
        path: '/about',
        name: 'About',
        component: AboutPage
    },
    {
        path: '/admin',
        name: 'Admin',
        component: AdminPage // Asigna el componente AdminPage a la ruta /admin
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;