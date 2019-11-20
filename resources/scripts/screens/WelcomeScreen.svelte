<script>
import { IsAuthenticated, IsRegistered } from '~/utils/auth'
</script>

<style lang="postcss">
@import "../../styles/components/bg-animate";

.container {
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
}

.subtitle,
.title,
.signin-btn {
    color: theme('colors.gray.800');
    font-weight: 900;
    letter-spacing: 0.1em;
}

.title:hover + .subtitle,
.title:hover,
.signin-btn:hover {
    @apply bg-animate-short;
    @apply bg-animate-alternate;
    @apply bg-animate-colorless;
    animation-duration: .4s;
    cursor: pointer;
}

.title:hover + .subtitle,
.title:hover {
    background: none;
}

.title {
    font-size: 12vw;
    text-shadow: 5px 5px 0 theme('colors.highlight');
    padding-bottom: 1.5rem;
}

.subtitle {
    font-size: 3vw;
    text-shadow: 2px 2px 0 theme('colors.highlight');
    margin-top: -75px;
}

.signin-btn {
    padding: 1rem 3rem;
    font-size: 1.5vw;
    width: fit-content;
    margin: 3rem auto;
    border: 5px solid theme('colors.highlight');
}
</style>

<div class="container flex-col mx-auto">

    <h1 class="title w-full text-center">GoTasks</h1>
    <h2 class="subtitle w-full text-center text-4xl pointer-events-none">GET STUFF DONE</h2>

    {#if !$IsAuthenticated}
    <a class="signin-btn" href="/auth/github/login?from=/%23/auth-complete">
        <i class="fab fa-github mr-3" />
        <span class="mr-2">SIGN IN</span>
    </a>
    {:else}
        {#if !$IsRegistered}
        <a class="signin-btn" href="/#/auth-complete">
            COMPLETE REGISTATION
        </a>
        {:else}
        <a class="signin-btn" href="/#/app">
            OPEN DASHBOARD
        </a>
        {/if}
    {/if}

</div>