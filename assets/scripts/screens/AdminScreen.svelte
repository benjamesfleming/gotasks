<script>
import { replace } from 'svelte-spa-router'
import { getData } from '~/utils/fetch'
import { onAuthorized } from '~/utils/auth'

let allUsers = []

onAuthorized(
    ['iam:gotasks:users:*:list'],
    {
        onFailure () { replace('/dashboard') },
        onSuccess () {
            allUsers = getData('/api/users')
        }
    }
)
</script>

{#await allUsers then users}
    <table>
        <thead>
            <tr>
                <td>Email</td>
                <td>UserName</td>
                <td>Provider</td>
                <td>CreatedAt</td>
            </tr>
        </thead>
        <tbody>
            {#each users as user}
                <tr>
                    <td>{user.email}</td>
                    <td>{user.username}</td>
                    <td>{user.provider}</td>
                    <td>{user.created_at}</td>
                </tr>
            {/each}
        </tbody>
    </table>
{/await}