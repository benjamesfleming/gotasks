<script>
import { onMount } from 'svelte'
import { navigateTo } from 'svero'
import { UserObject } from '~/utils/auth'
import { get } from '~/utils/api'

onMount(async function () {
    await get('/auth/logout', {}, '').catch(console.error)
    UserObject.set(null)
    navigateTo('/#/')
})
</script>

<p>Sorry to see you go.</p>