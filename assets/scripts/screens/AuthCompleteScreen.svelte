<script>
import { parse } from 'qs'
import { querystring, replace } from 'svelte-spa-router'
import { login } from '~/utils/auth'

$: {
    const { access_token, refresh_token } = parse($querystring)
    if (!access_token || !refresh_token) {
        replace('/auth/error')
    }

    login({ access_token, refresh_token })
    replace('/admin')
}
</script>

<p>Loading...</p>