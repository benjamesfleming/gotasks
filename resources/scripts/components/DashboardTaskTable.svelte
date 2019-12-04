<script>
import { onMount } from 'svelte'
import { writable } from 'svelte/store'
import { Sortable, Plugins } from '@shopify/draggable'
import * as mudder from 'mudder' 
import * as moment from 'moment'
import { sortBy } from 'lodash'
import Swal from 'sweetalert2'
import { AuthObject as u } from '~/utils/auth'
import { TaskObject as t, ShowTaskModal } from '~/store'
import CheckBox from '~/components/CheckBox'

export let sort = (a, b) => 0
export let filter = tasks => tasks
export let sortable = false

let listEl = null

$: if (listEl != null && sortable) {
    let sortable = new Sortable(listEl, {
        draggable: '.sortable',
        handle: '.sortable-handle',
        mirror: {
            append: '.sortable-list',
            constrainDimensions: true,
        },
        swapAnimation: {
            duration: 200,
            easingFunction: 'ease-in-out'
        },
        plugins: [Plugins.SwapAnimation]
    })

    sortable.on('sortable:stop', onSortStop);
}

// On Sort Stop
// handles when the user stop sorting a item (they drop it)
// `https://shopify.github.io/draggable/`
let onSortStop = e => {
    let { oldIndex, newIndex } = e.data
    let tasks = $u.tasks.filter(filter).sort(sort)
        tasks.splice(newIndex, 0, tasks.splice(oldIndex, 1)[0])

    let task = tasks[newIndex]
    let prevTask = tasks[newIndex - 1]
    let nextTask = tasks[newIndex + 1]

    // if the item has been moved to the start of the list
    // then set its position to `a` and set the position of the previous first item
    // to in between `a` and whatever the previous second item was
    // `https://fasiha.github.io/post/mudder/`

    if (tasks.length > 1) {
        if (prevTask == null && nextTask == null) {
            task.position = 'a'
        } else {
            if (newIndex == 0) {
                task.position = 'a'
                nextTask.position = mudder.alphabet.mudder('a', (tasks[2] || {}).position || '')[0]
            } else {
                task.position = mudder.alphabet.mudder(prevTask.position, (nextTask || {}).position || '')[0]
            }
        }
    }

    // console.log('----')
    // tasks.sort(sort).forEach(t => console.log(t.title, t.position))

    u.updateTask(task)
    nextTask && u.updateTask(nextTask)
}

// On Toggle
// hides the tasks steps then toggles it completion
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

// On Task Edit
// opens the task model for task editing
let onTaskEdit = id => {
    t.set(
        $u.tasks.find(t => t.id == id)
    )
    ShowTaskModal.set(true)
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
    <div bind:this={listEl} class="sortable-list w-full overflow-hidden rounded shadow-md transition-all hover:shadow-lg">
        {#each $u.tasks.filter(filter).sort(sort) as task (task.id)}
            <div class="sortable opacity-75 hover:opacity-80" data-id={task.id}>
                <div class="sortable-content flex justify-between items-center bg-gray-400 cursor-pointer {$showSteps[task.id] && 'py-2'} transition-all" on:click={() => onToggleSteps(task.id)}>
                    <div class="p-2 sm:p-3 w-16">
                        <CheckBox checked={task.isCompleted} on:change={() => onToggle(task.id)}/>
                    </div>
                    <div class="p-2 sm:p-3 flex-1">
                        <span>{task.title}</span>
                        <div class="flex flex-wrap">
                            {#each task.tags.split(',') as tag, idx}
                                <div class="px-3 py-0 mt-1 mr-2 text-center text-xs bg-gray-300 rounded-sm shadow-md hover:shadow-lg transition-all">{tag}</div>
                            {/each}
                        </div>
                    </div>
                    <div class="p-2 sm:p-3 flex-1 hidden sm:block">
                        {task.isCompleted ? moment(task.completedAt).fromNow() : 'Incomplete'}
                    </div>
                    <div class="p-2 sm:p-3 pr-4 sm:pr-6 text-lg flex flex-col sm:flex-row"> 
                        <div class="flex-1 flex justify-around {$showSteps[task.id] && 'mb-2 sm:mb-0'}">
                            <i class="fas fa-sort {sortable ? 'text-gray-900 hover:text-gray-800 transition-all sortable-handle' : 'text-gray-500'} sm:ml-4"></i>
                            <i class="fas fa-chevron-circle-{$showSteps[task.id] ? 'up' : 'down'} ml-3 sm:ml-4 text-clip bg-gray-900 hover:bg-gray-800 active:bg-animate transition-all"></i>
                        </div>
                        <div class="flex-1 justify-around {$showSteps[task.id] ? 'flex' : 'hidden sm:flex'}">
                            <i class="fas fa-pen-square sm:ml-4 text-clip bg-gray-900 hover:bg-gray-800 active:bg-animate transition-all" on:click|stopPropagation={() => onTaskEdit(task.id)}></i>
                            <i class="fas fa-trash ml-3 sm:ml-4 text-clip bg-gray-900 hover:bg-gray-800 active:bg-animate transition-all" on:click|stopPropagation={() => onTaskDelete(task.id)}></i>
                        </div>
                    </div>
                </div>
                <div class="sortable-content step-list {$showSteps[task.id] ? 'max-h-full py-3' : 'max-h-0 py-0'} overflow-hidden transition-all bg-gray-300">
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
