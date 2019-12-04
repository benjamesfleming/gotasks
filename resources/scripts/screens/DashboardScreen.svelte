<script>
import { onMount } from 'svelte'
import * as moment from 'moment'
import fitty from 'fitty'
import fuzzyFilter from 'fuzzy-array-filter'
import TaskModal from '~/components/DashboardTaskModal'
import TaskTable from '~/components/DashboardTaskTable'
import WarningPanel from '~/components/WarningPanel'
import { AuthObject as u } from '~/utils/auth'
import { ShowTaskModal } from '~/store'

u.loadTasks()

onMount(() => {
    fitty(document.querySelector('#user-info--full-name'))
    fitty(document.querySelector('#user-info--email'))
})

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

let onTaskModalShow = () => {
    ShowTaskModal.set(true)
}
</script>

<TaskModal />

<div class="container mx-auto my-6">

    <div class="max-w-xs px-6">
        <img class="rounded-lg shadow-md mb-6 transition-all hover:shadow-lg" src="{$u.avatar+"?s=512"}" alt="{$u.fullName}'s Profile Picture"/>
        <div class="user-info">
            <h1 id="user-info--full-name" class="text-4xl">{$u.fullName}</h1>
            <p id="user-info--email">{$u.email}</p>
            <small>Joined {moment($u.createdAt).format('Do MMM YY')}</small>
        </div>
    </div>

    <div class="w-full px-6">

        <div class="toolbar h-12 mb-6 flex justify-between">
            <div class="toolbar-input-group">
                <span><input type="text" bind:value={searchQuery} /></span>
                <i class="fas fa-search"></i>
            </div>

            <a class="button my-3" href="#/app" on:click={() => onTaskModalShow()}>
                <i class="fas fa-plus text-gray-800"></i> Add
            </a>
        </div>

        {#if searchQuery != ''}
            <TaskTable filter={fuzzyFilter(searchQuery, searchOptions)} sort={(a, b) => moment(a.completedAt).isAfter(b.completedAt) ? 1 : -1}>
                <WarningPanel content="NO TASKS FOUND.<br/>TRY A DIFFERENT SEARCH TERM" image="/assets/images/undraw_imagination.svg"/>
            </TaskTable>
        {:else}
            {#if $u.tasks.length > 0}
                <TaskTable filter={t => !t.isCompleted} sort={(a, b) => a.position > b.position ? 1 : -1} sortable>
                    <WarningPanel content="WOO-HOO!!!<br/>YOU ARE UP TO DATE" image="/assets/images/undraw_completed.svg"/>
                </TaskTable>

                <hr class="h-1 my-6 bg-gray-300 shadow-sm rounded"/>

                <TaskTable filter={t => t.isCompleted} sort={(a, b) => moment(a.completedAt).isBefore(b.completedAt) ? 1 : -1}>
                    <WarningPanel content="COMPLETE TASKS TO SEE THEM HERE" image="/assets/images/undraw_void.svg"/>
                </TaskTable>
            {:else}
                <WarningPanel content="WHAT... NO TASKS?<br/>CREATE ONE TO GET STARTED" image="/assets/images/undraw_progress_tracking.svg"/>
            {/if}
        {/if}
    </div>
</div>