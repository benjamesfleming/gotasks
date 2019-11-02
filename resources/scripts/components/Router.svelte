<script>
import { Router, Route } from 'svero'

import { IsRegistered, IsAdmin } from '~/utils/auth'

import AdminScreen from '~/screens/AdminScreen'
import AuthCompleteScreen from '~/screens/AuthCompleteScreen'
import AuthLogoutScreen from '~/screens/AuthLogoutScreen'
import NotFoundScreen from '~/screens/NotFoundScreen'
import ErrorScreen from '~/screens/ErrorScreen'
import WelcomeScreen from '~/screens/WelcomeScreen'
import DashboardScreen from '~/screens/DashboardScreen'
</script>

<Router>
    <Route fallback path="*" component={NotFoundScreen} />
    
    <Route exact path="#/" component={WelcomeScreen} />
    <Route exact path="#/error" component={ErrorScreen} />

    <Route exact path="#/auth-complete" component={AuthCompleteScreen} />
    <Route exact path="#/auth-logout" component={AuthLogoutScreen} />

    <Route exact path="#/admin" component={AdminScreen} condition={$IsAdmin} redirect="#/" />

    <Router exact path="#/app" condition={$IsRegistered} redirect="#/" nofallback>
        <Route exact component={DashboardScreen} />
        <Route exact path="/tasks/create">
            <h2>Create Task Here</h2>
        </Route>
    </Router>
</Router>