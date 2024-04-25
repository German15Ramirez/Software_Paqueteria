import { createRouter, createWebHistory } from 'vue-router';
import HomePage from './views/Home-Page.vue';
import LoginPage from './views/Login-Page.vue';
import AboutPage from './views/About-Page.vue';
import AdministradorPrincipalPage from './views/AdministradorPrincipal-Page.vue';
import RecepcionistaPrincipalPage from './views/RecepcionistaPrincipal-Page.vue';
import OperadorPrincipalPage from './views/OperadorPrincipal-Page.vue';
import AdminUsersPage from './views/AdminUsers-Page.vue';
import AdminRoutesPage from './views/AdminRoutes-Page.vue';
import AdminCheckpointsPage from './views/AdminCheckpoints-Page.vue';
import ListPackagesPage from './views/ListPackages-Page.vue';
import PackageLocationPage from './views/PackageLocation-Page.vue';
import ProcessPackagesPage from './views/ProcessPackages-Page.vue';
import RouteReportPage from './views/RouteReport-Page.vue';
import ProfitReportPage from './views/ProfitReport-Page.vue';
import CustomerReportPage from './views/CustomerReport-Page.vue';
import PopularRoutesReportPage from './views/PopularRoutesReport-Page.vue';
import AdminReportsPage from "@/views/AdminReports-Page.vue";
import ListarUsuariosPage from "@/views/ListarUsuarios-Page.vue"; // Importa el componente ListarUsuariosPage
import ListarRutasPage from "@/views/ListarRutas-Page.vue";
import ListarPuntosPage from "@/views/ListarPuntos-Page.vue";

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
        path: '/OperadorPrincipal-Page',
        name: 'OperadorPrincipal',
        component: OperadorPrincipalPage
    },
    {
        path: '/RecepcionistaPrincipal-Page',
        name: 'RecepcionistaPrincipal',
        component: RecepcionistaPrincipalPage
    },
    {
        path: '/AdministradorPrincipal-Page',
        name: 'AdministradorPrincipal',
        component: AdministradorPrincipalPage
    },
    {
        path: '/admin/users',
        name: 'AdminUsers',
        component: AdminUsersPage
    },
    {
        path: '/admin/routes',
        name: 'AdminRoutes',
        component: AdminRoutesPage
    },
    {
        path: '/admin/reports',
        name: 'AdminReports',
        component: AdminReportsPage
    },
    {
        path: '/admin/checkpoints',
        name: 'AdminCheckpoints',
        component: AdminCheckpointsPage
    },
    {
        path: '/recep/packages',
        name: 'ListPackages',
        component: ListPackagesPage
    },
    {
        path: '/recep/package-location',
        name: 'PackageLocation',
        component: PackageLocationPage
    },
    {
        path: '/operator/process-packages',
        name: 'ProcessPackages',
        component: ProcessPackagesPage
    },
    {
        path: '/admin/reports/route-report',
        name: 'RouteReport',
        component: RouteReportPage
    },
    {
        path: '/admin/reports/profit-report',
        name: 'ProfitReport',
        component: ProfitReportPage
    },
    {
        path: '/admin/reports/customer-report',
        name: 'CustomerReport',
        component: CustomerReportPage
    },
    {
        path: '/admin/reports/popular-routes-report',
        name: 'PopularRoutesReport',
        component: PopularRoutesReportPage
    },
    {
        path: '/admin/users/listar',
        name: 'ListarUsuarios',
        component: ListarUsuariosPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/routes/listar',
        name: 'ListarRutas',
        component: ListarRutasPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/checkpoints/listar',
        name: 'ListarPuntosDeControl',
        component: ListarPuntosPage // Agrega la ruta para ListarUsuariosPage
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;