<script>
import { IsAuthenticated, IsRegistered } from '~/utils/auth'
</script>

<style lang="postcss">
nav > a {
    letter-spacing: .1rem;
}
</style>

<nav class="container mx-auto flex justify-between text-gray-900 px-6 py-2">

    <a class="font-bold text-xl" href="/#/">
        GoTasks
    </a>

    <div class="flex h-full px-2">
        <a class="text-gray-900 mr-4" href="/#/app">Dashboard</a>
        <a class="text-gray-900 mr-4" href="/#/account">Account</a>
        <a class="text-gray-900" href="/#/auth-logout">Logout</a>
    </div>
    
</nav>