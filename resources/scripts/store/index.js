import { writable } from 'svelte/store'
import { Task } from '~/models'

export * from './_user'

export const TaskObject = writable(new Task({}))
export const ShowTaskModal = writable(false)

/**
 * Create Writable Store
 * wraps `svelte/store` to add  a local storage backend
 * @see https://higsch.me/2019/06/22/2019-06-21-svelte-local-storage/
 * @param {*} key 
 * @param {*} startValue 
 */
export function createWritableStore (key, startValue) {
    const { subscribe, set, update } = writable(startValue)
  
	return {
        subscribe, set, update,
        useLocalStorage (mutator=v=>v) {
            const json = localStorage.getItem(key);
            if (json) {
                set(mutator(JSON.parse(json).data))
            }
            
            subscribe(current => {
                const data = JSON.stringify({ data: current })
                current != null && current != ''
                    ? localStorage.setItem(key, data)
                    : localStorage.removeItem(key)
            })

            return { subscribe, set, update }
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