<script>
import { parse } from 'qs'
import { querystring, push, replace } from 'svelte-spa-router'
import { UserObject } from '~/utils/auth'

$: {
    const { user_id } = parse($querystring)
    if (!user_id) {
        push('/auth/error')
    }

    fetch('/api/users/' + user_id)
        .then(res => res.ok ? res.json() : null)
        .then(user => { UserObject.set(user); console.log(user)})
        .then(() => replace('/dashboard'))
        .catch(() => push('/auth/error'))
}
</script>

<p>Loading...</p>