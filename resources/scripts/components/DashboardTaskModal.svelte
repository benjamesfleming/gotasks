<script>
import { writable } from 'svelte/store'
import { sortBy, maxBy } from 'lodash'
import { AuthObject as u } from '~/utils/auth'
import { TaskObject as t, ShowTaskModal } from '~/store'
import Modal from '~/components/Modal'

let show = {note: false, tags: true, steps: true}
let isValid = false
let errors = {}

let tags = []
let steps = []
let title = ""
let note = ""

$: $t.id, onResetFields()
$: isNewTask = $t.id ? false : true

let onResetFields = () => {
    console.log('resetting fields')
    if ($t.id != null && $t.id != '') {
        tags = $t.tags.split(',')
        steps = $t.steps
        title = $t.title
        note = $t.note
    } else {
        tags = []
        steps = []
        title = ""
        note = ""
    }
}

// On Step Enter
// handle all keyup events on the step input
// add the current step into the array, then reset
// the input text box
let onStepEnter = e => {
    isValid = /^[\x00-\x7F]+$/g.test(e.target.value) ? null : true
    errors = { ...errors, steps: isValid } 
    if (e.key == "Enter" && errors.steps == null) {
        steps = [...steps, { title: e.target.value, order: steps.length + 1 }]
        e.target.value = ""
    }
}

// On Step Remove
// remove a given step from the steps array
let onStepRemove = idx => {
    steps = steps.filter((v, i) => i != idx)
}

// On Tag Enter, handle all keyup events on the tag input
// handle all keyup events on the tag input
// add the current tag into the array, then reset
// the input text box
let onTagEnter = e => {
    isValid = /^[\x00-\x7F]+$/g.test(e.target.value) ? null : true
    errors = { ...errors, tags: isValid } 
    if (e.key == "Enter" && errors.tags == null) {
        tags = [...tags, e.target.value]
        e.target.value = ""
    }
}

// On Tag Remove
// remove a given tag from the tags array
let onTagRemove = idx => {
    tags = tags.filter((v, i) => i != idx)
}

// On Continue Click
// validate the given data then submit to the api
let onContinueClick = async () => {
    let [ok, err] = isNewTask
        ? (await createTask())
        : (await updateTask())
    
    if (err != null && Object.keys(err.all || {}).length > 0) {
        errors = err.all
    } else {
        onClose()
    }
}

// Create Task
// creates the task in the database
let createTask = async () => {
    return await u.createTask({ 
        title, note, steps,
        tags: tags.join(", "), 
    })
}

// Update Task
// updates the task in the database
let updateTask = async () => {
    return await u.updateTask({
        ...$t, title, note, steps,
        tags: tags.join(", "),
    })
}

// On Clear Click
// clear all the current changes
let onClearClick = () => {
    errors = {}
    onResetFields()
}

// On Clear Error
// clear the error of a given input from the array
let onClearError = key => {
    let { [key]: e, ...errs } = errors
    errors = errs
}

// On Close
// handle the close button
let onClose = () => {
    t.set({})
    ShowTaskModal.set(false)
}

</script>

<style lang="postcss">
.vis-selector:hover > i {
    @apply text-gray-400;
}
</style>

{#if $ShowTaskModal}
<Modal>

    <!-- 
        Modal Header
     -->
    <div class="flex justify-between items-center text-gray-800 p-2 mb-6">
        <h2 class="text-xl font-extrabold">
            {isNewTask ? 'CREATE TASK' : 'UPDATE TASK'}
        </h2>
        <i class="text-xl cursor-pointer fas fa-times" on:click={() => onClose()}></i>
    </div>

    <!-- 
        Title Input
        -->
    <label class="block px-2 pb-2 bg-white overflow-hidden">
        <span class="text-gray-700">Task Title</span>
        {#if errors.title}<span class="block text-sm text-red-500">*{errors.title}</span>{/if}
        <input class="form-input mt-1 block w-full {errors.title ? 'border-red-500' : ''}" placeholder="My New Task Title" bind:value={title} on:focus={() => onClearError('title')}>
    </label>

    <!-- 
        Note Input
        -->
    <label class="block px-2 {show.note ? 'pb-2' : ''} mt-3 bg-white overflow-hidden transition-all">
        <span class="vis-selector text-gray-700 flex items-center justify-between cursor-pointer" on:click={() => show.note = !show.note}>
            Descriptive Note <i class="fas fa-chevron-circle-{show.note ? 'up' : 'down'} transition-all"></i>
        </span>
        {#if errors.note}<span class="block text-sm text-red-500">*{errors.note}</span>{/if}
        <div class="{show.note ? 'max-h-full' : 'max-h-0'} transition-all">
            <textarea class="form-textarea mt-1 block w-full h-24 resize-none {errors.note ? 'border-red-500' : ''}" placeholder="..." disabled={!show.note} bind:value={note} on:focus={() => onClearError('note')}></textarea>
        </div>
    </label>

    <!-- 
        Tags Input
        -->
    <label class="block px-2 {show.tags ? 'pb-2' : ''} mt-3 bg-white overflow-hidden transition-all">
        <span class="vis-selector text-gray-700 flex items-center justify-between cursor-pointer" on:click={() => show.tags = !show.tags}>
            Tags <i class="fas fa-chevron-circle-{show.tags ? 'up' : 'down'} transition-all"></i>
        </span>
        <div class="{show.tags ? 'max-h-full' : 'max-h-0'} transition-all">
            {#if tags.length > 0}
                <div class="pb-3 flex flex-wrap">
                    {#each tags as tag, idx}
                        <div class="px-3 py-1 mt-2 mr-2 text-center text-sm bg-gray-300 rounded shadow-md hover:shadow-lg hover:bg-red-400 transition-all cursor-pointer" on:click={() => onTagRemove(idx)}>{tag}</div>
                    {/each}
                </div>
            {/if}
            <input class="form-input mt-1 block w-full {errors.tags ? 'border-red-500' : ''}" placeholder="Enter Some Tags" on:keyup={onTagEnter} disabled={!show.tags}>
        </div>
    </label>

    <!-- 
        Steps Input
        -->
    <label class="block px-2 {show.steps ? 'pb-2' : ''} mt-3 text-gray-700 bg-white overflow-hidden transition-all">
        <span class="vis-selector flex items-center justify-between cursor-pointer" on:click={() => show.steps = !show.steps}>
            Steps <i class="fas fa-chevron-circle-{show.steps ? 'up' : 'down'} transition-all"></i>
        </span>
        <div class="{show.steps ? 'max-h-full' : 'max-h-0'} transition-all">
            {#each sortBy(steps, ['order']) as step, idx (idx)}
                <div class="w-full mt-1 flex">
                    <span class="w-12 pr-1 text-center text-sm self-center">{idx +1} •</span>
                    <input class="form-input border-none outline-none block w-full" bind:value={step.title}>
                    <button class="button" on:click={() => onStepRemove(idx)}>
                        
                    </button>
                </div>
            {/each}
            <div class="w-full mt-1 flex">
                <input class="form-input block w-full {errors.steps ? 'border-red-500' : ''}" placeholder="Extra Step ..." on:keyup={onStepEnter} disabled={!show.steps}>
                <button class="button" on:click={onStepEnter}>
                    ⏎
                </button>
            </div>
        </div>
    </label>

    <button class="form-button m-2 mt-0" on:click={onContinueClick}>Continue</button>
    {#if !isNewTask}
    <button class="form-button-clear text-gray-800" on:click={onClearClick}>Clear</button>
    {/if}
</Modal>
{/if}