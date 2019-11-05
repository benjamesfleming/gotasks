<script>
import { IsAuthenticated, IsRegistered } from '~/utils/auth'
</script>

<style lang="scss">
nav {
    @apply flex justify-between;
    @apply bg-gray-700 text-white;
    @apply shadow-xl;
    @apply h-16;
}

nav > h2 {
    @apply font-bold;
    @apply px-4 w-32 flex;
    @apply items-center justify-center;
    letter-spacing: .1rem;
}

nav > div {
    @apply flex;
    height: 100%;
}

nav > div > a {
    @apply px-4 my-2 mr-2 w-32 flex rounded;
    @apply items-center justify-center;
    transition: all 0.2s ease-in-out;

    &:hover {
        @apply bg-gray-500 shadow-inner;
    }
}
</style>

<nav>

    <h2>GoTasks</h2>

    <div>
    {#if !$IsAuthenticated}
    <a href="/auth/github/login?from=/#/auth-complete">Auth w/GitHub</a>
    {:else}
        {#if $IsRegistered}
        <a href="/#/app">Dashboard</a>
        {:else}
        <a href="/#/auth-complete">Complete Registation</a>
        {/if}

        <a href="/#/auth-logout">Logout</a>
    {/if}
    </div>
    
</nav>