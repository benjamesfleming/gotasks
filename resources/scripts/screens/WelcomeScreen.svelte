<script>
import { onMount } from 'svelte'
import fitty from 'fitty'
import { IsAuthenticated, IsRegistered } from '~/utils/auth'

onMount(() => {
    fitty(document.querySelector('#title'))
    fitty(document.querySelector('#subtitle'))
})
</script>

<style lang="postcss">
@import "../../styles/components/bg-animate";

#container {
    height: 100vh;
    width: 768px;
    max-width: 90%;
    display: flex;
    align-items: center;
    justify-content: center;
}

#subtitle,
#title,
#signin-btn {
    color: theme('colors.gray.800');
    font-weight: 900;
    letter-spacing: 0.1em;
}

#title:hover + #subtitle,
#title:hover,
#signin-btn:hover {
    @apply bg-animate-short;
    @apply bg-animate-alternate;
    @apply bg-animate-colorless;
    animation-duration: .4s;
    cursor: pointer;
}

#title:hover + #subtitle,
#title:hover {
    background: none;
}

#title {
    font-size: 4em;
    text-shadow: 0.05em 0.05em 0 theme('colors.highlight');
    padding-bottom: 0;
}

#subtitle {
    font-size: 1.25em;
    letter-spacing: 0.2em;
    text-shadow: 0.075em 0.075em 0 theme('colors.highlight');
    margin-top: -0.5em;
    padding: 0 5%;
}

#signin-btn {
    padding: 0.75em 2em;
    font-size: 1em;
    width: fit-content;
    margin: 1em auto;
    border: 5px solid theme('colors.highlight');
}
</style>

<div id="container" class="flex-col mx-auto">

    <span id="title" class="pointer-events-none">GoTasks</span>
    <span id="subtitle" class="pointer-events-none">GET STUFF DONE</span>

    {#if !$IsAuthenticated}
    <a id="signin-btn" href="/auth/github/login?from=/%23/auth-complete">
        <i class="fab fa-github mr-3" />
        <span class="mr-2">SIGN IN</span>
    </a>
    {:else}
        {#if !$IsRegistered}
        <a id="signin-btn" href="/#/auth-complete">
            COMPLETE REGISTATION
        </a>
        {:else}
        <a id="signin-btn" href="/#/app">
            OPEN DASHBOARD
        </a>
        {/if}
    {/if}

</div>