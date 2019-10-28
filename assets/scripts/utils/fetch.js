import { getStoreValue } from '~/utils/store'
import { AccessToken, RefreshToken } from '~/utils/auth'

/**
 * Get Headers
 * build a object with the default headers
 */
async function getHeaders () {
    const authToken = await getStoreValue(AccessToken)
    return {
        Authorization: 'Bearer ' + authToken
    }
}

/**
 * Refresh Token
 * atempt to refresh the access token
 */
export async function updateAccessToken (options) {
    const authToken = await getStoreValue(RefreshToken)
    const headers = {
        Authorization: 'Bearer ' + authToken
    }

    const response = await fetch(
        '/auth/refresh', { headers, ...options }
    )

    if (response.status == 200) {
        const json = await response.json()
        AccessToken.set(json.access_token)
    }
}

/**
 * Get Data
 * send a GET request to the server
 * @param {*} uri 
 * @param {*} options 
 */
export async function getData (uri, { retry = true, ...options } = {}) {
    const headers = await getHeaders()
    const response = await fetch(
        uri, { headers, ...options }
    )

    if (response.status == 401 && retry) {
        await updateAccessToken()
        return getData(
            uri, { retry: false, ...options }
        )
    }

    return response.json()
}