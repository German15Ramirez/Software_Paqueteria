import { createRouter, createWebHistory } from 'vue-router';
import HomePage from './views/Home-Page.vue';
import LoginPage from './views/Login-Page.vue';
import AboutPage from './views/About-Page.vue';
import AdministradorPrincipalPage from './views/AdministradorPrincipal-Page.vue';
import RecepcionistaPrincipalPage from './views/RecepcionistaPrincipal-Page.vue';
import OperadorPrincipalPage from './views/OperadorPrincipal-Page.vue';
import ProcesosPage from './views/ProcessPackages-Page.vue';
import AdminUsersPage from './views/AdminUsers-Page.vue';
import AdminRoutesPage from './views/AdminRoutes-Page.vue';
import AdminCheckpointsPage from './views/AdminCheckpoints-Page.vue';
import ListPackagesPage from './views/ListPackages-Page.vue';
import PackageLocationPage from './views/PackageLocation-Page.vue';
import PackageStatusPage from './views/PackageStatus-Page.vue';
import ProcessPackagesPage from './views/ProcessPackages-Page.vue';
import RouteReportPage from './views/RouteReport-Page.vue';
import ProfitReportPage from './views/ProfitReport-Page.vue';
import CustomerReportPage from './views/CustomerReport-Page.vue';
import PopularRoutesReportPage from './views/PopularRoutesReport-Page.vue';
import AdminReportsPage from "./views/AdminReports-Page.vue";
import AdminClientesPage from "./views/AdminClientes-Page.vue";
import CrearClientesPage from "./views/CrearClientes-Page.vue";
import ActualizarClientesPage from "./views/ActualizarClientes-Page.vue";
import ListarClientesPage from "./views/ListarClientes-Page.vue";
import EliminarClientesPage from "./views/EliminarClientes-Page.vue";
import ListarUsuariosPage from "./views/ListarUsuarios-Page.vue";
import CrearUsuariosPage from "./views/CrearUsuarios-Page.vue";
import ActualizarUsuariosPage from "./views/ActualizarUsuarios-Page.vue";
import EliminarUsuariosPage from "./views/EliminarUsuarios-Page.vue";
import TarifaGlobalPage from "./views/TarifaGlobal-Page.vue";
import ListarRutasPage from "./views/ListarRutas-Page.vue";
import CrearRutasPage from "./views/CrearRutas-Page.vue";
import EliminarRutasPage from "./views/EliminarRutas-Page.vue";
import ActualizarRutasPage from "./views/ActualizarRutas-Page.vue";
import ListarPuntosPage from "./views/ListarPuntos-Page.vue";
import CrearPuntosPage from "./views/CrearPuntos-Page.vue";
import EliminarPuntosPage from "./views/EliminarPuntos-Page.vue";
import ActualizarPuntosPage from "./views/ActualizarPuntos-Page.vue";

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
        path: '/oper/process',
        name: 'ProcesosOperador',
        component: ProcesosPage
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
        path: '/recep/locations',
        name: 'PackageLocation',
        component: PackageLocationPage
    },
    {
        path: '/recep/status',
        name: 'PackageStatus',
        component: PackageStatusPage
    },
    {
        path: '/recep/administrar_clientes',
        name: 'AdminClientes',
        component: AdminClientesPage
    },
    {
        path: '/recep/clientes/listar',
        name: 'ListarClientes',
        component: ListarClientesPage
    },
    {
        path: '/recep/clientes/crear',
        name: 'CrearClientes',
        component: CrearClientesPage
    },
    {
        path: '/recep/clientes/actualizar',
        name: 'ActualizarClientes',
        component: ActualizarClientesPage
    },
    {
        path: '/recep/clientes/eliminar',
        name: 'EliminarClientes',
        component: EliminarClientesPage
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
        path: '/admin/users/crear',
        name: 'CrearUsuarios',
        component: CrearUsuariosPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/users/actualizar',
        name: 'ActualizarUsuarios',
        component: ActualizarUsuariosPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/users/eliminar',
        name: 'EliminarUsuarios',
        component: EliminarUsuariosPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/routes/listar',
        name: 'ListarRutas',
        component: ListarRutasPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/routes/crear',
        name: 'CrearRutas',
        component: CrearRutasPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/routes/eliminar',
        name: 'EliminarRutas',
        component: EliminarRutasPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/routes/actualizar',
        name: 'ActualizarRutas',
        component: ActualizarRutasPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/routes/establecer_tarifa_global',
        name: 'EstablecerTarifaGlobal',
        component: TarifaGlobalPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/checkpoints/listar',
        name: 'ListarPuntosDeControl',
        component: ListarPuntosPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/checkpoints/crear',
        name: 'CrearPuntosDeControl',
        component: CrearPuntosPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/checkpoints/eliminar',
        name: 'EliminarPuntosDeControl',
        component: EliminarPuntosPage // Agrega la ruta para ListarUsuariosPage
    },
    {
        path: '/admin/checkpoints/actualizar',
        name: 'ActualizarPuntosDeControl',
        component: ActualizarPuntosPage // Agrega la ruta para ListarUsuariosPage
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;