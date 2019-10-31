<script>
import { UserObject } from '~/utils/auth'

let user = $UserObject
let tasks = fetch(`/api/users/${user.id}/tasks`).then(res => res.json())
</script>

<ul>
    <li>Username: {user.username}</li>
    <li>Email: {user.email}</li>
    <li>Provider: {user.provider}</li>
    <li>CreatedAt: {user.created_at}</li>
</ul>

{#await tasks}
<code>
{JSON.stringify(tasks)}
</code>
{/await}