<script>
import { IsAuthenticated, IsRegistered } from '~/utils/auth'
</script>

<style lang="postcss">
nav > h2 {
    letter-spacing: .1rem;
}

nav a:not(:last-child) {
    @apply mr-4;
}
</style>

<nav class="container mx-auto flex justify-between text-gray-900 px-6 py-2">

    <h2 class="font-bold text-xl flex items-center justify-center">
        GoTasks
    </h2>

    <div class="flex h-full px-2">
        <a class="text-gray-900" href="/#/app">Dashboard</a>
        <a class="text-gray-900" href="/#/auth-logout">Logout</a>
    </div>
    
</nav>