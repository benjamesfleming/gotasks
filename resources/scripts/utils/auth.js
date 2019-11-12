import { derived } from 'svelte/store'
import { navigateTo } from 'svero'
import { get } from '~/utils/api'
import { createWritableStore, getStoreValue } from '~/utils/store'

export const UserObject = createWritableStore('user', '').useLocalStorage()

/**
 * Is Authenticated
 * true / false, is the user currently logged in
 */
export const IsAuthenticated = derived(
    UserObject, $User => $User != null && $User.id != null
)

/**
 * Is Registered
 * true / fase, is the logged in user registered
 */
export const IsRegistered = derived(
    [UserObject, IsAuthenticated], ([$User, $IsAuthenticated]) => $IsAuthenticated && ($User != null && $User.isRegistered)
)

/**
 * Is Admin
 * true / fase, is the logged in user an admin
 */
export const IsAdmin = derived(
    [UserObject, IsAuthenticated], ([$User, $IsAuthenticated]) => $IsAuthenticated && ($User != null && $User.isAdmin)
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

/**
 * Reauthenticate
 * check that the user is authenticated by calling the server
 */
export async function reAuthenticate () {
    const [, err] = await get('/auth/user', {}, '')
    if (err != null) {
        UserObject.set(null)
        navigateTo('/#/')
    }
}