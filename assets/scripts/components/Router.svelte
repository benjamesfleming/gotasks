<script>
import Router from 'svelte-spa-router'
import AdminScreen from '~/screens/AdminScreen'
import AuthCompleteScreen from '~/screens/AuthCompleteScreen'
import AuthLogoutScreen from '~/screens/AuthLogoutScreen'
import NotFoundScreen from '~/screens/NotFoundScreen'
import WelcomeScreen from '~/screens/WelcomeScreen'
import DashboardScreen from '~/screens/DashboardScreen'

const routes = {
    '/'              : WelcomeScreen,
    '/auth/complete' : AuthCompleteScreen,
    '/auth/logout'   : AuthLogoutScreen,
    '/admin'         : AdminScreen,
    '/dashboard'     : DashboardScreen,
    '*'              : NotFoundScreen,
}
</script>

<Router {routes}/>