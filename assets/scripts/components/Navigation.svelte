<script>
import { IsAuthenticated } from '~/utils/auth'
</script>

{#if !$IsAuthenticated}
<a href="/auth/github">Auth w/GitHub</a>
{:else}
<a href="#/dashboard">Dashboard</a>
<a href="#/auth/logout">Logout</a>
{/if}