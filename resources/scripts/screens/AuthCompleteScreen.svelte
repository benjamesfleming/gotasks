<script>
import { onMount } from 'svelte'
import { navigateTo } from 'svero'
import { UserObject } from '~/utils/auth'
import { get, post } from '~/utils/api'
import { User } from '~/models'

let currentUser
let validationErrors

let onComplete = async () => {
    let [returnedUser, error] = await post('/auth/register', currentUser, {}, '')
    let hasErrors = (error != null && error.code != null)

    if (!hasErrors) {
        UserObject.set(
            User.fromApi(returnedUser)
        )
        navigateTo('/#/app')
    } else {
        validationErrors = Object.keys(error.all).map(k => `${k}: ${error.all[k]}!`)
    }
}

onMount(async function () {
    let [providedUser] = await get('/auth/user', {}, '')
    let isRegistered = providedUser['attrs']['registered'] == true

    if (isRegistered) {
        let [returnedUser] = await get('/auth/me', {}, '')
        UserObject.set(
            User.fromApi(returnedUser)
        )
        navigateTo('/#/app')
    } else {
        currentUser = User.fromProvider(providedUser)
        UserObject.set(
            currentUser
        )
    }
})
</script>

{#if currentUser == null}
    <p>Loading...</p>  
{:else}
    <h4>Hello, {currentUser.fullName}!</h4>
    <p>Enter details to finish registration!</p>

    {#each validationErrors || [] as error}
        <p>{error}</p>
    {/each}

    <input type="text" bind:value={currentUser.firstName} placeholder="First Name"/>
    <input type="text" bind:value={currentUser.lastName} placeholder="Last Name"/>
    <input type="text" bind:value={currentUser.username} placeholder="Username"/>
    <input type="email" bind:value={currentUser.email} placeholder="Email"/>

    <input type="button" value="Complete" on:click={onComplete}/>
{/if}