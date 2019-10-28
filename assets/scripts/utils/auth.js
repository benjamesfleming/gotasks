import * as jwt from 'jwt-simple'
import { writable } from 'svelte/store'
import { createWritableStore, getStoreValue } from '~/utils/store'

export const AccessToken = createWritableStore('access_token', '').useLocalStorage()
export const RefreshToken = createWritableStore('refresh_token', '').useLocalStorage()
export const UserObject = createWritableStore('user', '').useLocalStorage()

/**
 * Is Authenticated
 * true / false, is the user currently logged in
 */
export const IsAuthenticated = writable(false)
UserObject.subscribe(
    user => {
        IsAuthenticated.set(user != null && user.exp - ((new Date).getTime() / 1000 | 0) > 0)
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

/**
 * Login
 * save the current access and refresh tokens
 * @param {object} tokens 
 */
export function login ({ access_token, refresh_token }) {
    AccessToken.set(access_token)
    RefreshToken.set(refresh_token)
    UserObject.set(
        jwt.decode(access_token, '', true)
    )
}

/**
 * Logout
 * clear all the personalized data from the client
 */
export function logout () {
    AccessToken.set(null)
    RefreshToken.set(null)
    UserObject.set(null)
}