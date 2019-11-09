<script>
import CheckBox from '~/components/CheckBox'

export let tasks;
</script>

<table class="table-fixed w-full overflow-hidden rounded shadow-md transition-all hover:shadow-lg">
    <tbody>
        {#each tasks as task}
            <tr class="bg-gray-300 odd:bg-gray-400 opacity-75 hover:opacity-80 transition-all">
                <td class="p-3 w-16">
                    <CheckBox/>
                </td>
                <td class="p-3">{task.title}</td>
                <td class="p-3">{task.isCompleted ? "true" : "false"}</td>
                <td class="p-3">{task.streak}</td>
                <td class="p-3">
                    <a href="/#/app/tasks/{task.id}/edit">Edit</a>
                    <span>
                        <input type="checkbox" checked={task.isCompleted}/>
                    </span>
                </td>
            </tr>
        {:else}
            <tr class="p-3 bg-gray-300 odd:bg-gray-400">
                <td class="p-3" colspan="5">No Tasks Found... <a href="/#/app/tasks/create">Create One</a></td>
            </tr>
        {/each}
    </tbody>
</table>