<script>
import { Router, Route } from 'svero'

import { IsAuthenticated } from '~/utils/auth'

import AdminScreen from '~/screens/AdminScreen'
import AuthCompleteScreen from '~/screens/AuthCompleteScreen'
import AuthLogoutScreen from '~/screens/AuthLogoutScreen'
import NotFoundScreen from '~/screens/NotFoundScreen'
import ErrorScreen from '~/screens/ErrorScreen'
import WelcomeScreen from '~/screens/WelcomeScreen'
import DashboardScreen from '~/screens/DashboardScreen'
</script>

<Router>
    <Route path="*" component={NotFoundScreen} />
    <Route path="/error" component={ErrorScreen} />
    <Route path="/" component={WelcomeScreen} exact />
    <Route path="/auth/complete" component={AuthCompleteScreen} />
    <Route path="/auth/logout" component={AuthLogoutScreen} />
    <Route path="/admin" component={AdminScreen} condition={$IsAuthenticated} redirect="/" />
    <Route path="/dashboard" component={DashboardScreen} condition={$IsAuthenticated} redirect="/" />
</Router>