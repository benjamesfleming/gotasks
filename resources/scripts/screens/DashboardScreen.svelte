<script>
import * as moment from 'moment'
import fuzzyFilter from 'fuzzy-array-filter'
import TaskModal from '~/components/DashboardTaskModal'
import TaskTable from '~/components/DashboardTaskTable'
import { AuthObject as u } from '~/utils/auth'

u.loadTasks()

let taskModal = false
let searchQuery = ''
let searchOptions = {
    keys: [
        'title',
        'tags',
        'note',
        'steps.title'
    ],
    id: 'id'
}
</script>

{#if taskModal}
<TaskModal on:close={() => taskModal = false}/>
{/if}

<div class="container mx-auto my-6">

    <div class="max-w-xs px-6">
        <img class="rounded-lg shadow-md mb-6 transition-all hover:shadow-lg" src="{$u.avatar}" alt="{$u.fullName}'s Profile Picture"/>
        <div class="user-info">
            <h1 class="text-4xl">{$u.fullName}</h1>
            <p>{$u.email}</p>
            <small>Joined {moment($u.createdAt).format('Do MMM YY')}</small>
        </div>
    </div>

    <div class="w-full px-6">

        <div class="toolbar h-12 mb-6 flex justify-between">
            <div class="toolbar-input-group">
                <span><input type="text" bind:value={searchQuery} /></span>
                <i class="fas fa-search"></i>
            </div>

            <a class="button my-3" href="#/app" on:click={() => taskModal = true}>
                <i class="fas fa-plus"></i> Add
            </a>
        </div>

        {#if searchQuery != ''}
            <TaskTable filter={fuzzyFilter(searchQuery, searchOptions)} sort={(a, b) => moment(a.completedAt).isAfter(b.completedAt) ? 1 : -1}>
                <div class="mx-auto p-12 flex items-center justify-center">
                    <img class="max-w-md pl-6 pr-6 drop-shadow-md hover:drop-shadow-lg transition-all" src="/assets/images/undraw_imagination.svg" alt="Kiwi standing on oval">
                    <p class="text-4xl max-w-xl text-center font-extrabold text-gray-800">
                        NO TASKS FOUND. <br/> TRY A DIFFERENT SEARCH TERM
                    </p>
                </div>
            </TaskTable>
        {:else}
            {#if $u.tasks.length > 0}
                <TaskTable filter={t => !t.isCompleted} sort={(a, b) => a.position > b.position ? 1 : -1} sortable>
                    <div class="max-w-lg mx-auto py-12 flex items-center justify-center">
                        <img class="max-w-xs pr-12 drop-shadow-md hover:drop-shadow-lg transition-all" src="/assets/images/undraw_completed.svg" alt="Kiwi standing on oval">
                        <p class="text-3xl text-center font-extrabold text-gray-800">
                            WOO-HOO!!!<br/>YOU ARE UP TO DATE
                        </p>
                    </div>
                </TaskTable>

                <hr class="h-1 my-6 bg-gray-300 shadow-sm rounded"/>

                <TaskTable filter={t => t.isCompleted} sort={(a, b) => moment(a.completedAt).isBefore(b.completedAt) ? 1 : -1}>
                    <div class="max-w-lg mx-auto py-12 flex items-center justify-center">
                        <img class="max-w-xs pr-12 drop-shadow-md hover:drop-shadow-lg transition-all" src="/assets/images/undraw_void.svg" alt="Kiwi standing on oval">
                        <p class="text-3xl text-center font-extrabold text-gray-800">
                            COMPLETE TASKS TO SEE THEM HERE
                        </p>
                    </div>
                </TaskTable>
            {:else}
                <div class="mx-auto p-12 flex items-center justify-center">
                    <img class="max-w-md pl-6 pr-6 drop-shadow-md hover:drop-shadow-lg transition-all" src="/assets/images/undraw_progress_tracking.svg" alt="Kiwi standing on oval">
                    <p class="text-4xl max-w-xl text-center font-extrabold text-gray-800">
                        WHAT... NO TASKS?<br/> CREATE ONE TO GET STARTED
                    </p>
                </div>
            {/if}
        {/if}
    </div>
</div>