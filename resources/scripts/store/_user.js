import * as mudder from 'mudder' 
import { uniqBy } from 'lodash'
import { Task, User } from '~/models'
import { createWritableStore, getStoreValue } from '~/store'
import { get, post, del } from '~/utils/api'

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
            let { id } = await getStoreValue({ subscribe })
            let [tasks] = await get(`/users/${id}/tasks`)
            for (let i in tasks) {
                tasks[i] = Task.fromApi(tasks[i])
            }
            update(u => new User({ ...u, tasks }))
        },

        // Toggle Task, toggle the task with the given id
        // e.g. createUserStore(user).toggleTask(id)
        async toggleTask (id) {
            let { tasks } = await getStoreValue({ subscribe })
            let sortedTasks = tasks
                .filter(t => t.id != id && t.position != ' ')
                .sort((a, b) => a.position > b.position ? 1 : -1)
            let task = tasks.find(t => t.id == id)

            if (task != null) {
                task.isCompleted = !task.isCompleted

                if (!task.isCompleted && sortedTasks.length > 0) {
                    sortedTasks.unshift(task)
                    sortedTasks[0].position = 'a'
                    sortedTasks[1].position = mudder.alphabet.mudder('a', (sortedTasks[2] || {}).position || '')[0]

                    tasks = uniqBy([...tasks, ...sortedTasks], 'id')
                    task = tasks.find(t => t.id == id)

                    await post(`/tasks/${sortedTasks[0].id}`, sortedTasks[0])
                    await post(`/tasks/${sortedTasks[1].id}`, sortedTasks[1])
                } else {
                    task.position = !task.isCompleted
                        ? 'a'
                        : ' '

                    await post(`/tasks/${task.id}`, task)
                }
            }

            update(u => new User({ ...u, tasks }))
        },

        // Toggle Step, toggles the step of a task with the given id
        // e.g. createUserStore(user).toggleStep(taskId, stepId)
        async toggleStep (taskId, stepId) {
            let { tasks } = await getStoreValue({ subscribe })
            let task = tasks.find(t => t.id == taskId)
            let step = task.steps.find(s => s.id == stepId)
            if (step != null) {
                step.isCompleted = !step.isCompleted
                await post(`/tasks/${task.id}/steps/${step.id}`, step)
            }
            update(u => new User({ ...u, tasks }))
        },

        // Create Task, create a new task with given data
        // e.g. createUserStore(user).createTask(task)
        async createTask (task) {
            let { tasks } = await getStoreValue({ subscribe })
            let sortedTasks = tasks
                .filter(t => t.position != ' ')
                .sort((a, b) => a.position > b.position ? 1 : -1)

            task.position = 'a'

            if (sortedTasks.length > 0) {
                let firstTask = sortedTasks[0]
                let secondTask = sortedTasks[1]

                firstTask.position = mudder.alphabet.mudder('a', (secondTask || {}).position || '')[0]
                await post(`/tasks/${firstTask.id}`, firstTask)
            }
           
            let [response, error] = await post(`/tasks`, task)
            let newTask = Task.fromApi(response || {})

            if (response != null) {
                update(u => new User({ ...u, tasks: [...u.tasks, newTask] }))
            }
            
            return [response != null, error]
        },

        // Delete Task, deletes the task with the given id
        // e.g. createUserStore(user).deleteTask(id)
        async deleteTask (id) {
            let [response, error] = await del(`/tasks/${id}`)
            if (response != null) {
                update(u => {
                    let tasks = u.tasks.filter(t => t.id != id)
                    return new User({ ...u, tasks })
                })
            }
        },

        // Order Task, reorders the tasks array to move the given indexs
        // e.g. createUserStore(user).orderTask(id, 0, 1) // move task index 0 to index 1
        async updateTask (task) {
            let { tasks } = await getStoreValue({ subscribe })
            let taskIdx = tasks.findIndex(t => t.id == task.id)
            tasks[taskIdx] = task
            await post(`/tasks/${task.id}`, task)
            update(u => new User({ ...u, tasks }))
        },
    }

    // Preload Requested Data
    tasks && store.loadTasks()

    return store
}