<script>
import { Router, Route } from 'svero'

import { IsAuthenticated } from '~/utils/auth'

import AdminScreen from '~/screens/AdminScreen'
import AuthCompleteScreen from '~/screens/AuthCompleteScreen'
import AuthLogoutScreen from '~/screens/AuthLogoutScreen'
import AuthErrorScreen from '~/screens/AuthErrorScreen'
import NotFoundScreen from '~/screens/NotFoundScreen'
import WelcomeScreen from '~/screens/WelcomeScreen'
import DashboardScreen from '~/screens/DashboardScreen'
</script>

<Router>
    <Route path="*" component={NotFoundScreen} />
    <Route path="#/" component={WelcomeScreen} exact />
    <Route path="#/auth/complete" component={AuthCompleteScreen} />
    <Route path="#/auth/logout" component={AuthLogoutScreen} />
    <Route path="#/auth/error" component={AuthErrorScreen} />
    <Route path="#/admin" component={AdminScreen} condition={$IsAuthenticated} redirect="#/" />
    <Route path="#/dashboard" component={DashboardScreen} condition={$IsAuthenticated} redirect="#/" />
</Router>