<script>
import { onMount } from 'svelte'
import { navigateTo } from 'svero'
import { User } from '~/models'
import { AuthObject } from '~/utils/auth'
import { get, post } from '~/utils/api'
import Modal from '~/components/Modal'

let isLoading = true
let errors = {}

let avatar = ""
let username = ""
let firstName = ""
let lastName = ""
let email = ""

// On Continue Click
// validate the given data then submit to the api
let onContinueClick = async () => {
    let [returnedUser, error] = await post('/auth/register', { username, firstName, lastName, email }, {}, '')
    let hasErrors = (error != null && error.code != null)

    if (!hasErrors) {
        AuthObject.fromApi(returnedUser)
        navigateTo('/#/app')
    } else {
        errors = error.all
    }
}

// On Clear Error
// clear the error of a given input from the array
let onClearError = key => {
    let { [key]: e, ...errs } = errors
    errors = errs
}

// On Mount
// on component mount, get the provided user
onMount(async function () {
    let [providedUser] = await get('/auth/user', {}, '')
    let isRegistered = providedUser['attrs']['registered'] == true

    if (isRegistered) {
        AuthObject.fromApi(
            (await get('/auth/me', {}, ''))[0]
        )

        navigateTo('/#/app')
    } else {
        let u = User.fromProvider(providedUser)

        avatar      = u.avatar
        username    = u.username
        firstName   = u.firstName
        lastName    = u.lastName
        email       = u.email

        isLoading   = false
    }
})
</script>

<Modal fadeIn={false}>
{#if isLoading == true}
    <p>Loading...</p>  
{:else}
    
    <img style="margin-top: -120px;" class="w-1/2 mx-auto mb-6 rounded-lg shadow-md hover:shadow-lg transition-all" src={avatar} alt=""/>

    <!-- 
        Modal Header
     -->
    <div class="flex flex-col text-gray-800 p-2 mb-6">
        <h2 class="text-xl font-extrabold">COMPLETE REGISTRATION</h2>
        <span>Enter details to finish registration</span>
    </div>

    <div class="flex">
        <!-- 
            First Name Input
            -->
        <label class="w-1/2 flex flex-col justify-between px-2 pb-2 bg-white overflow-hidden">
            <span class="text-gray-700">First Name</span>
            {#if errors.firstName}<span class="block text-sm text-red-500">*{errors.firstName}</span>{/if}
            <input class="form-input mt-1 block w-full {errors.firstName ? 'border-red-500' : ''}" placeholder="" bind:value={firstName} on:focus={() => onClearError('firstName')}>
        </label>

        <!-- 
            Last Name Input
            -->
        <label class="w-1/2 flex flex-col justify-between px-2 pb-2 bg-white overflow-hidden">
            <span class="text-gray-700">Last Name</span>
            {#if errors.lastName}<span class="block text-sm text-red-500">*{errors.lastName}</span>{/if}
            <input class="form-input mt-1 block w-full {errors.lastName ? 'border-red-500' : ''}" placeholder="" bind:value={lastName} on:focus={() => onClearError('lastName')}>
        </label>
    </div>

    <!-- 
        Username Input
        -->
    <label class="block px-2 pb-2 bg-white overflow-hidden">
        <span class="text-gray-700">Username</span>
        {#if errors.username}<span class="block text-sm text-red-500">*{errors.username}</span>{/if}
        <input class="form-input mt-1 block w-full {errors.username ? 'border-red-500' : ''}" placeholder="" bind:value={username} on:focus={() => onClearError('username')}>
    </label>

    <!-- 
        Email Input
        -->
    <label class="block px-2 pb-2 bg-white overflow-hidden">
        <span class="text-gray-700">Email</span>
        {#if errors.email}<span class="block text-sm text-red-500">*{errors.email}</span>{/if}
        <input class="form-input mt-1 block w-full {errors.email ? 'border-red-500' : ''}" type="email" placeholder="" bind:value={email} on:focus={() => onClearError('email')}>
    </label>

    <button class="form-button m-2 mt-0" on:click={onContinueClick}>Register</button>
{/if}
</Modal>
