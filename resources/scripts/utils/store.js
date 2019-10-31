import { writable } from 'svelte/store'

/**
 * Create Writable Store
 * wraps `svelte/store` to add  a local storage backend
 * @see https://higsch.me/2019/06/22/2019-06-21-svelte-local-storage/
 * @param {*} key 
 * @param {*} startValue 
 */
export function createWritableStore (key, startValue) {
    const { subscribe, set } = writable(startValue)
  
	return {
        subscribe, set,
        useLocalStorage: () => {
            const json = localStorage.getItem(key);
            if (json) {
                set(JSON.parse(json).data)
            }
            
            subscribe(current => {
                const data = JSON.stringify({ data: current })
                current != null && current != ''
                    ? localStorage.setItem(key, data)
                    : localStorage.removeItem(key)
            })

            return { subscribe, set }
        }
    }
}

/**
 * Get Store Value
 * subscribe to a store and return the current value
 * @param {*} store 
 */
export async function getStoreValue (store) {
    let unsub = () => {}
    let value = await new Promise(
        done => {
            unsub = store.subscribe(
                tkn => done(tkn)
            )
        }
    )
    unsub(); return value
}