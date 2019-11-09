<script>
import * as moment from 'moment'
import { UserObject } from '~/utils/auth'
import { User } from '~/models'

$: u = new User($UserObject)
</script>

<div class="container mx-auto mt-6">

    <div class="max-w-xs px-6">
        <img class="rounded-lg shadow-md mb-6 transition-all hover:shadow-lg" src="{u.avatar}" alt="{u.fullName}'s Profile Picture"/>
        <div class="user-info">
            <h1 class="text-4xl">{u.fullName}</h1>
            <p>{u.email}</p>
            <small>Joined {moment(u.createdAt).format('Do MMM YY')}</small>
        </div>
    </div>

    <div class="w-full px-6">

        <div class="toolbar flex justify-between">
            <div class="toolbar-input-group">
                <span><input type="text" /></span>
                <i class="fas fa-search"></i>
            </div>

            <a class="button" href="#/app/tasks/create">
                <i class="fas fa-plus"></i> Add
            </a>
        </div>

        <div class="toolbar flex justify-start my-6">
            <button class="mx-6 ml-0">ALL</button>
            <button class="mx-6">DUE TODAY</button>
            <button class="mx-6 mr-0">REPEATING</button>
        </div>

        <table class="w-full overflow-hidden rounded shadow-md transition-all hover:shadow-lg">
            <tbody>
            {#await u.tasks then tasks}
                {#each tasks as task}
                <tr class="bg-gray-200 odd:bg-gray-300 opacity-75 hover:opacity-100">
                    <td class="p-3">{task.id}</td>
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
                <tr class="p-3 bg-gray-200 odd:bg-gray-300">
                    <td class="p-3" colspan="5">No Tasks Found... <a href="/#/app/tasks/create">Create One</a></td>
                </tr>
                {/each}
            {/await}
            </tbody>
        </table>
    </div>
</div>