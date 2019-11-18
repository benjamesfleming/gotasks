<script>
import { Router, Route } from 'svero'

import { IsRegistered, IsAdmin } from '~/utils/auth'
import Navigation from '~/components/Navigation'

import AuthCompleteScreen from '~/screens/AuthCompleteScreen'
import AuthLogoutScreen from '~/screens/AuthLogoutScreen'
import NotFoundScreen from '~/screens/NotFoundScreen'
import ErrorScreen from '~/screens/ErrorScreen'
import WelcomeScreen from '~/screens/WelcomeScreen'
import DashboardScreen from '~/screens/DashboardScreen'
</script>

<Router path="#">
    <Route exact component={WelcomeScreen} />
    <Route exact path="/error" component={ErrorScreen} />

    <Route exact path="/auth-complete" component={AuthCompleteScreen} />
    <Route exact path="/auth-logout" component={AuthLogoutScreen} />
    
    <Route exact path="/app" condition={() => $IsRegistered} redirect="#/">
        <Navigation/>
        <DashboardScreen/>
    </Route>
    
    <Route fallback path="*" component={NotFoundScreen} />
</Router>