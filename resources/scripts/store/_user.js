import { Task, User } from '~/models'
import { createWritableStore, getStoreValue } from '~/store'
import { get, post } from '~/utils/api'

/**
 * Create User Store
 * wraps `svelte/store` to add all user functions
 * @param {*} user 
 * @param {*} preload
 * @param {bool} saveToLocalStorage
 */
export function createUserStore (user, { tasks }, saveToLocalStorage = false) {
    const _user = new User(user || {})
    const { subscribe, set, update } = saveToLocalStorage
        ? createWritableStore('user', _user).useLocalStorage(u => new User(u))
        : createWritableStore('user', _user)

    const store = {
        subscribe, 
        
        // Set User, update the user using the given user data
        // e.g. createUserStore(null).set(user)
        set: user => set(new User(user || {})),

        // Set From API, update the user using the user data from the api
        // e.g. createUserStore(null).fromApi(apiResponse)
        fromApi: user => set(User.fromApi(user)),

        // Load Tasks, load all the given users tasks from the api
        // e.g. createUserStore(user).loadTasks()
        async loadTasks () {
            const { id } = await getStoreValue({ subscribe })
            const [tasks] = await get(`/users/${id}/tasks`)
            for (let i in tasks) {
                tasks[i] = Task.fromApi(tasks[i])
            }
            update(u => new User({ ...u, tasks }))
        },

        // Toggle Task, toggle the task with the given id
        // e.g. createUserStore(user).toggleTask(id)
        async toggleTask (id) {
            const { tasks } = await getStoreValue({ subscribe })
            const task = tasks.find(t => t.id == id)
            if (task != null) {
                task.isCompleted = !task.isCompleted
                await post(`/tasks/${task.id}`, task)
            }
            update(u => new User({ ...u, tasks }))
        },

        // Create Task, create a new task with given data
        // e.g. createUserStore(user).createTask(task)
        async createTask (task) {
            const [response, error] = await post(`/tasks`, task)
            const _t = Task.fromApi(response || {})
            if (response != null) {
                update(u => new User({ ...u, tasks: [...u.tasks, _t] }))
            }
            return [response != null, error]
        },
    }

    // Preload Requested Data
    tasks && store.loadTasks()

    return store
}