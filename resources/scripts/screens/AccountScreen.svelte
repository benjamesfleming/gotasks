<script>
import { navigateTo } from 'svero'
import { AuthObject as u } from '~/utils/auth'
import Modal from '~/components/Modal'

let errors = {}

let avatar = $u.avatar
let username = $u.username
let firstName = $u.firstName
let lastName = $u.lastName
let email = $u.email

// On Clear Error
// clear the error of a given input from the array
let onClearError = key => {
    let { [key]: e, ...errs } = errors
    errors = errs
}

// On Save Click
// handle when the user clicks the save
let onSaveClick = () => {

}

// On Back Click
// take the user back to the dashbaord
let onBackClick = () => {
    navigateTo('/#/app')
}
</script>

<Modal fadeIn={false}>
    <img style="margin-top: -120px;" class="w-1/2 mx-auto mb-6 rounded-lg shadow-md hover:shadow-lg transition-all" src={$u.avatar} alt=""/>

    <!-- 
        Modal Header
    -->
    <div class="flex flex-col text-gray-800 p-2 mb-6">
        <h2 class="text-xl font-extrabold">MY ACCOUNT</h2>
        <span>Update and save details about your account</span>
        <span class="text-red-500">WARNING: Editing Not Implemented</span>
    </div>

    <div class="flex">
        <!-- 
            First Name Input
            -->
        <label class="w-1/2 flex flex-col justify-between px-2 pb-2 bg-white overflow-hidden">
            <span class="text-gray-700">First Name</span>
            {#if errors.firstName}<span class="block text-sm text-red-500">*{errors.firstName}</span>{/if}
            <input disabled class="form-input mt-1 block w-full {errors.firstName ? 'border-red-500' : ''}" placeholder="" bind:value={firstName} on:focus={() => onClearError('firstName')}>
        </label>

        <!-- 
            Last Name Input
            -->
        <label class="w-1/2 flex flex-col justify-between px-2 pb-2 bg-white overflow-hidden">
            <span class="text-gray-700">Last Name</span>
            {#if errors.lastName}<span class="block text-sm text-red-500">*{errors.lastName}</span>{/if}
            <input disabled class="form-input mt-1 block w-full {errors.lastName ? 'border-red-500' : ''}" placeholder="" bind:value={lastName} on:focus={() => onClearError('lastName')}>
        </label>
    </div>

    <!-- 
        Username Input
        -->
    <label class="block px-2 pb-2 bg-white overflow-hidden">
        <span class="text-gray-700">Username</span>
        {#if errors.username}<span class="block text-sm text-red-500">*{errors.username}</span>{/if}
        <input disabled class="form-input mt-1 block w-full {errors.username ? 'border-red-500' : ''}" placeholder="" bind:value={username} on:focus={() => onClearError('username')}>
    </label>

    <!-- 
        Email Input
        -->
    <label class="block px-2 pb-2 bg-white overflow-hidden">
        <span class="text-gray-700">Email</span>
        {#if errors.email}<span class="block text-sm text-red-500">*{errors.email}</span>{/if}
        <input disabled class="form-input mt-1 block w-full {errors.email ? 'border-red-500' : ''}" type="email" placeholder="" bind:value={email} on:focus={() => onClearError('email')}>
    </label>

    <button disabled class="form-button m-2 mt-0" on:click={onSaveClick}>Save</button>
    <button class="form-button-clear text-gray-800" on:click={onBackClick}>Back</button>
</Modal>