<script>
import * as moment from 'moment'
import { AuthObject as u } from '~/utils/auth'
import CheckBox from '~/components/CheckBox'

export let filter = tasks => tasks

const onToggle = id => u.toggleTask(id)
</script>


{#if $u.tasks.filter(filter).length > 0}
    <table class="table-fixed w-full overflow-hidden rounded shadow-md transition-all hover:shadow-lg">
        <tbody>
            {#each $u.tasks.filter(filter) as task}
                <tr class="bg-gray-300 odd:bg-gray-400 opacity-75 hover:opacity-80 transition-all">
                    <td class="p-3 w-16">
                        <CheckBox checked={task.isCompleted} on:change={() => onToggle(task.id)}/>
                    </td>
                    <td class="p-3">{task.title}</td>
                    {#if task.isCompleted}
                        <td class="p-3">{moment(task.completed).fromNow()}</td>
                    {:else}
                        <td class="p-3">Incomplete</td>
                    {/if}
                    <td class="p-3">{task.streak}</td>
                    <td class="p-3">
                        <a href="/#/app/tasks/{task.id}/edit">Edit</a>
                        <span>
                            <input type="checkbox" checked={task.isCompleted}/>
                        </span>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
{:else}
    <slot></slot>
{/if}
