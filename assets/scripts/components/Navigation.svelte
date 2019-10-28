<script>
import { link } from 'svelte-spa-router'
import { IsAuthenticated } from '~/utils/auth'
</script>

{#if !$IsAuthenticated}
<a href="/auth/github">Auth w/GitHub</a>
{:else}
<a href="/auth/logout" use:link>Logout</a>
{/if}