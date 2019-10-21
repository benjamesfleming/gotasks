import { createWritableStore } from '~/utils/store'

export const AccessToken = createWritableStore('access_token', '').useLocalStorage()
export const RefreshToken = createWritableStore('refresh_token', '').useLocalStorage()

export function login ({ access_token, refresh_token }) {
    AccessToken.set(access_token)
    RefreshToken.set(refresh_token)
}

export function logout () {
    AccessToken.set(null)
    RefreshToken.set(null)
}