<script>
import { IsAuthenticated, IsRegistered } from '~/utils/auth'
</script>

{#if !$IsAuthenticated}
    <a href="/auth/github/login?from=/auth-complete">Auth w/GitHub</a>
{:else}
    {#if $IsRegistered}
        <a href="/dashboard">Dashboard</a>
    {:else}
        <a href="/auth-complete">Complete Registation</a>
    {/if}

    <a href="/auth-logout">Logout</a>
{/if}