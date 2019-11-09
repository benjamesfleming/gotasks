import * as Cookie from 'js-cookie'
import { navigateTo } from 'svero'
import { UserObject } from '~/utils/auth'

/**
 * Post
 * sends data to the api using the post method
 * @param {string} path 
 * @param {object} body 
 * @param {object} options 
 * @param {string} prefix
 */
export async function post(path, body, { headers, ...options }, prefix='/api') {
   return __fetch(
        path, {
            method: 'POST',
            body: JSON.stringify(body),
            headers: {
                'Content-Type': 'application/json',
                ...headers
            },
            ...options
        },
        prefix
    )
}

/**
 * Get
 * sends data to the api using the get method
 * @param {string} path 
 * @param {object} options 
 * @param {string} prefix
 */
export async function get(path, options={}, prefix='/api') {
    return __fetch(
         path, {
             method: 'GET',
             ...options
         },
         prefix
     )
 }

/**
 * Fetch
 * basic wrapper araound the window.fetch api to add auth headers
 * @param {string} path 
 * @param {object} options 
 * @param {string} prefix
 * @returns {array} [data, errors]
 */
export async function __fetch(path, { headers, ...options }, prefix='') {
    const response = await fetch(
        prefix + path,
        { 
            headers: {
               'X-XSRF-TOKEN': Cookie.get('XSRF-TOKEN'),
               // add more default headers here ...
               ...headers
            },
            ...options
        }
    )

    switch (response.status) {
        // 401 Unauthorized
        case 401:
            UserObject.set(null)
            navigateTo('/')
            break

        // 400 Bad Request
        case 400:
            return [null, { code: 400, all: (await response.json()) }]
            break

        // 200 Success
        case 200:
            return [(await response.json()), null]
            break

        // Default
        default:
            throw new Error(`[GOTASKS] ERROR - Failed to fetch api \`${prefix}${path}\`. \n ${response.status} ${response.statusText}`)
            break
    }
}