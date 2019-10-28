import { writable } from 'svelte/store'
import { createWritableStore, getStoreValue } from '~/utils/store'

export const UserObject = createWritableStore('user', '').useLocalStorage()

/**
 * Is Authenticated
 * true / false, is the user currently logged in
 */
export const IsAuthenticated = writable(false)
UserObject.subscribe(
    user => {
        IsAuthenticated.set(user != null && user.id != null)
    }
)

/**
 * Has Privilege
 * checks if the user has a given provilege
 * @param {string} privilege 
 */
export async function hasPrivilege (privilege='') {
    const { privileges } = await getStoreValue(UserObject)
    if (Array.from(privileges || []).indexOf(privilege) > -1) {
        return true
    }
    return false
}

/**
 * On Authorized
 * check that the user has all the given callbacks, then run
 * the corrosponding callback
 * @param {string[]} privileges 
 * @param {object} callbacks 
 */
export async function onAuthorized (privileges=[], { onSuccess, onFailure }) {
    const ok = (await Promise.all(privileges.map(hasPrivilege))).indexOf(false) == -1
    return (
        ok ? onSuccess() : onFailure()
    )
}