<script>
import { IsAuthenticated, IsRegistered } from '~/utils/auth'
</script>

<style lang="postcss">
nav {
    background: theme('colors.primary');
}

nav > h2 {
    letter-spacing: .1rem;
}

nav > div {
    height: 100%;
}
</style>

<nav class="flex justify-between text-white shadow-xl h-16">

    <h2 class="font-bold px-4 w-32 flex items-center justify-center">
        GoTasks
    </h2>

    <div class="flex h-100">
    {#if !$IsAuthenticated}
    <a class="button" href="/auth/github/login?from=/#/auth-complete">
        <i class="fab fa-github"></i> Login
    </a>
    {:else}
        {#if $IsRegistered}
        <a class="button" href="/#/app">Dashboard</a>
        {:else}
        <a class="button" href="/#/auth-complete">Complete Registation</a>
        {/if}

        <a class="button" href="/#/auth-logout">Logout</a>
    {/if}
    </div>
    
</nav>