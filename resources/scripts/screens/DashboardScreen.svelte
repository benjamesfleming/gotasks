<script>
import * as moment from 'moment'
import { UserObject } from '~/utils/auth'
import { User } from '~/models'

$: u = new User($UserObject)
</script>

<style lang="scss">
.container {
    @apply flex py-6;
}

.user-sidebar {
    @apply max-w-xs px-6;
}

.user-sidebar > img {
    @apply rounded-lg shadow-md;
    @apply mb-6;
    transition: all 0.2s ease-in-out;

    &:hover {
        @apply shadow-lg;
    }
}

.user-sidebar > .user-info {
    h1 {
        @apply text-4xl;
    }
}

.content {
    @apply px-6;
    width: fill;
}

.toolbar {
    @apply flex justify-between;
    @apply mb-2;
}

.toolbar > .input-group {
    @apply bg-gray-400;
    @apply rounded shadow-md;
    @apply flex justify-center items-center;
    transition: all 0.2s ease-in-out;
    box-sizing: border-box;
    overflow: hidden;
    cursor: pointer;

    span, input {
        margin: 0;
        padding: 0;
    }

    span {
        display: inline-block;
        @apply h-full;
        overflow: hidden;
    }

    input {
        @apply h-full;
        border: none;
        padding: 8px;
    }

    i {
        @apply py-2 px-2 pr-3;
        color: #444;
    }

    &:hover {
        @apply bg-gray-300 shadow-lg;
    }
}

.toolbar > a {
    @apply bg-gray-400;
    @apply py-2 w-24;
    @apply rounded shadow-md;
    @apply flex justify-center items-center;
    transition: all 0.2s ease-in-out;

    i {
        @apply pr-3;
        color: #444;
    }

    &:hover {
        @apply bg-gray-300 shadow-lg;
    }
}
</style>

<div class="container mx-auto mt-6">

    <div class="user-sidebar">
        <img src="{u.avatar}" alt="{u.fullName}'s Profile Picture"/>
        <div class="user-info">
            <h1>{u.fullName}</h1>
            <p>{u.email}</p>
            <small>Joined {moment(u.createdAt).format('Do MMM YY')}</small>
        </div>
    </div>

    <div class="content">

        <div class="toolbar">
            <div class="input-group">
                <span><input type="text" /></span>
                <i class="fas fa-search"></i>
            </div>

            <a href="#/app/tasks/create">
                <i class="fas fa-plus"></i> Add
            </a>
        </div>

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
                {#each u.tasks || [] as task}
                <tr>
                    <td>{task.id}</td>
                    <td>{task.title}</td>
                    <td>{task.completed ? "true" : "false"}</td>
                    <td>{task.streak}</td>
                    <td>
                        <a href="/#/app/tasks/{task.id}/edit">Edit</a>
                        <span>
                            <input type="checkbox" bind:value={task.completed}/>
                        </span>
                    </td>
                </tr>
                {:else}
                <tr>
                    <td colspan="5">No Tasks Found... <a href="/#/app/tasks/create">Create One</a></td>
                </tr>
                {/each}
            </tbody>
        </table>
    </div>
</div>