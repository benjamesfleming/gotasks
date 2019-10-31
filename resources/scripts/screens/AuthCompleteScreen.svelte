<script>
import { navigateTo } from 'svero'
import { UserObject } from '~/utils/auth'

const querystring = window.location.href.split('?')[1]
const userId = (new URLSearchParams(querystring)).get('user_id')

if (!userId) {
    navigateTo('/auth/error')
}

fetch('/api/users/' + userId)
    .then(res => res.ok ? res.json() : null)
    .then(user => { UserObject.set(user); console.log(user)})
    .then(() => navigateTo('/dashboard'))
    .catch(() => navigateTo('/auth/error'))
</script>

<p>Loading...</p>