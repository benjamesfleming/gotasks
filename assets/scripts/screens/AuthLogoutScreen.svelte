<script>
import { UserObject } from '~/utils/auth'

UserObject.set(null)
window.location.href = '/auth/logout'
</script>

<p>Sorry to see you go.</p>