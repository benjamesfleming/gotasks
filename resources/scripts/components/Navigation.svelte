<script>
import { IsAuthenticated, IsRegistered } from '~/utils/auth'

let open = false
</script>

<style lang="postcss">
nav > a {
    letter-spacing: .1rem;
}

.menu-icon {
    top: 0;
    right: 0;
}
</style>

<nav class="container m-auto relative flex flex-col sm:flex-row justify-between text-gray-900 py-2 px-4 sm:px-6">

    <a class="font-bold text-xl" href="/#/">
        GoTasks
    </a>

    <div class="flex flex-col sm:flex-row h-full p-2 sm:pt-0">
        <i class="cursor-pointer menu-icon fas fa-bars absolute m-4 sm:hidden" on:click={() => open = !open}></i>
        <a class="cursor-pointer text-gray-900 sm:mr-4 {!open && 'hidden sm:inline-block'}" href="/#/app">Dashboard</a>
        <a class="cursor-pointer text-gray-900 sm:mr-4 {!open && 'hidden sm:inline-block'}" href="/#/help">Help</a>
        <a class="cursor-pointer text-gray-900 sm:mr-4 {!open && 'hidden sm:inline-block'}" href="/#/account">Account</a>
        <a class="cursor-pointer text-gray-900 {!open && 'hidden sm:inline-block'}" href="/#/auth-logout">Logout</a>
    </div>
    
</nav>