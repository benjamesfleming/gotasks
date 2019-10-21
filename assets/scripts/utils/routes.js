import AdminScreen from '~/screens/AdminScreen'
import AuthCompleteScreen from '~/screens/AuthCompleteScreen'
import AuthLogoutScreen from '~/screens/AuthLogoutScreen'
import NotFoundScreen from '~/screens/NotFoundScreen'
import WelcomeScreen from '~/screens/WelcomeScreen'

export default {
    '/'              : WelcomeScreen,
    '/auth/complete' : AuthCompleteScreen,
    '/auth/logout'   : AuthLogoutScreen,
    '/admin'         : AdminScreen,
    '*'              : NotFoundScreen,
}