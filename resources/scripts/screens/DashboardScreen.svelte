<script>
import { UserObject as user } from '~/utils/auth'
</script>

<ul>
    <li>Username: {$user.username}</li>
    <li>Email: {$user.email}</li>
    <li>Provider: {$user.provider}</li>
    <li>CreatedAt: {$user.createdAt}</li>
</ul>

<table>
    <thead>
        <tr>
            <td>Id</td>
            <td>Title</td>
            <td>Completed</td>
            <td>Streak</td>
            <td>Actions</td>
        </tr>
    </thead>
    <tbody>
        {#each $user.tasks || [] as task}
        <tr>
            <td>{task.id}</td>
            <td>{task.title}</td>
            <td>{task.completed ? "true" : "false"}</td>
            <td>{task.streak}</td>
            <td>
                <a href="/dashboard/tasks/{task.id}/edit">Edit</a>
                <input type="checkbox" bind:value={task.completed}/>
            </td>
        </tr>
        {:else}
        <tr>
            <td colspan="5">No Tasks Found... <a href="/dashboard/tasks/create">Create One</a></td>
        </tr>
        {/each}
    </tbody>
</table>