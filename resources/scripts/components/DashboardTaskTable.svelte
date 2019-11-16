<script>
import { writable } from 'svelte/store'
import * as moment from 'moment'
import { sortBy } from 'lodash'
import Swal from 'sweetalert2'
import { AuthObject as u } from '~/utils/auth'
import CheckBox from '~/components/CheckBox'

export let sort = (a, b) => 0
export let filter = tasks => tasks

let onToggle = id => {
    showSteps.set({ ...$showSteps, [id]: false })
    u.toggleTask(id)
}

// On Toggle Steps
// toggles the visiblity of a given tasks steps
let showSteps = writable({})
let onToggleSteps = id => {
    showSteps.set({ ...$showSteps, [id]: !($showSteps[id] || false) })
}

// On Step Toggle
// toggles the completion of a given step
let onStepToggle = (task, id) => {
    u.toggleStep(task.id, id)
}

// On Task Delete
// deletes the task after confermation
let onTaskDelete = async id => {
    let { value } = await Swal.fire({
        title: 'Are you sure?',
        width: '400px',
        html: `
            <p>Once a task has been deleted<br/> it can\'t be bought back.</p>
        `,
        icon: 'warning',
        showCancelButton: true,
        cancelButtonColor: '#4299e1',
        confirmButtonColor: '#e53e3e',
        reverseButtons: true,
        confirmButtonText: 'Delete'
    })

    if (value == true) {
        u.deleteTask(id)
    }
}
</script>

<style lang="postcss">
.step-list > .step,
.step-list > .step > .step-title {
    position: relative;
}

.step-list > .step::after {
    @apply bg-gray-400;
    content: "";
    position: absolute;
    top: -60%;
    clip-path: polygon(0% 0%, 0% 100%, 25% 100%, 25% 0, 100% 0, 100% 85%, 25% 85%, 25% 100%, 100% 90%, 100% 0%);
    left: 1rem;
    width: 1.5rem;
    height: 100%;
    transform: scaleY(1.5);
}

.step-list > .step > .custom-line-through::after {
    @apply bg-gray-800;
    content: "";
    position: absolute;
    left: 10px;
    right: 10px;
    top: calc(50% - 1.5px);
    height: 3px;
    opacity: 0.8;
}
</style>

{#if $u.tasks.filter(filter).length > 0}
    <div class="w-full overflow-hidden rounded shadow-md transition-all hover:shadow-lg">
        {#each $u.tasks.filter(filter).sort(sort) as task (task.id)}
            <div class="opacity-75 hover:opacity-80 transition-all">
                <div class="flex justify-between items-center bg-gray-400 cursor-pointer" on:click={() => onToggleSteps(task.id)}>
                    <div class="p-3 w-16">
                        <CheckBox checked={task.isCompleted} on:change={() => onToggle(task.id)}/>
                    </div>
                    <div class="p-3 flex-1">
                        <span>{task.title}</span>
                        <div class="flex flex-wrap">
                            {#each task.tags.split(',') as tag, idx}
                                <div class="px-3 py-0 mt-1 mr-2 text-center text-xs bg-gray-300 rounded-sm shadow-md hover:shadow-lg transition-all" on:click={() => onTagRemove(idx)}>{tag}</div>
                            {/each}
                        </div>
                    </div>
                    <div class="p-3 flex-1">
                        {task.isCompleted ? moment(task.completedAt).fromNow() : 'Incomplete'}
                    </div>
                    <div class="p-3">
                        <i class="fas fa-chevron-circle-{$showSteps[task.id] ? 'up' : 'down'} text-grey-800 hover:text-gray-700 transition-all"></i>
                        <i class="fas fa-trash ml-3 text-red-600 hover:text-red-400 transition-all" on:click|stopPropagation={() => onTaskDelete(task.id)}></i>
                    </div>
                </div>
                <div class="step-list {$showSteps[task.id] ? 'max-h-full py-3' : 'max-h-0 py-0'} overflow-hidden transition-all bg-gray-300">
                    {#each sortBy(task.steps, ['order']) as step, idx (step.id)}
                        <div class="step flex py-1">
                            <div class="flex justify-end w-16">
                                <span class="text-center text-sm self-center">{idx +1} •</span>
                            </div>
                            <div class="step-title px-3 {step.isCompleted ? 'custom-line-through' : ''} cursor-pointer" on:click={() => onStepToggle(task, step.id)}>
                                {step.title}
                            </div> 
                        </div>
                    {/each}
                </div>
            </div>
        {/each}
    </div>
{:else}
    <slot></slot>
{/if}
