<script>
import { writable } from 'svelte/store'
import * as moment from 'moment'
import { sortBy } from 'lodash'
import { AuthObject as u } from '~/utils/auth'
import CheckBox from '~/components/CheckBox'


export let filter = tasks => tasks

let onToggle = id => u.toggleTask(id)

// On Toggle Steps
// toggles the visiblity of a given tasks steps
let showSteps = writable({})
let onToggleSteps = id => {
    showSteps.set({ ...$showSteps, [id]: !($showSteps[id] || false) })
}
</script>

<style lang="postcss">
.step-list > .step {
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
</style>

{#if $u.tasks.filter(filter).length > 0}
    <div class="w-full overflow-hidden rounded shadow-md transition-all hover:shadow-lg">
        {#each $u.tasks.filter(filter) as task (task.id)}
            <div class="flex justify-between items-center bg-gray-300 odd:bg-gray-400 cursor-pointer opacity-75 hover:opacity-80 transition-all" on:click={() => onToggleSteps(task.id)}>
                <div class="p-3 w-16">
                    <CheckBox checked={task.isCompleted} on:change={() => onToggle(task.id)}/>
                </div>
                <div class="p-3 flex-1">{task.title}</div>
                <div class="p-3 flex-1">
                    {task.isCompleted ? moment(task.completedAt).fromNow() : 'Incomplete'}
                </div>
                <div class="p-3">
                    <i class="fas fa-chevron-circle-{$showSteps[task.id] ? 'up' : 'down'} transition-all"></i>
                </div>
            </div>
            <div class="step-list {$showSteps[task.id] ? 'max-h-full py-3' : 'max-h-0 py-0'} overflow-hidden transition-all bg-gray-300">
                {#each sortBy(task.steps, ['order']) as step, idx}
                    <div class="step flex py-1">
                        <div class="flex justify-end w-16">
                            <span class="text-center text-sm self-center">{idx +1} •</span>
                        </div>
                        <div class="px-3">
                            {step.title} <br/>
                        </div> 
                    </div>
                {/each}
            </div>
        {/each}
    </div>
{:else}
    <slot></slot>
{/if}
