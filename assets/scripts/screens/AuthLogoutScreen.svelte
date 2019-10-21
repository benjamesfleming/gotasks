<script>
import { replace } from 'svelte-spa-router'
import { logout } from '~/utils/auth'

logout()
replace('/')
</script>

<p>Sorry to see you go.</p>