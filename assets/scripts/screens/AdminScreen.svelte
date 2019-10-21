<script>
import { getData } from '~/utils/fetch'
</script>

{#await getData('/api/users') then users}
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