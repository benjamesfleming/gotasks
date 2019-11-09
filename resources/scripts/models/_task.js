import * as moment from 'moment'

export class Task {
    constructor({ id, userId, parentId, title, tags, note, completed, streak, createdAt, updatedAt }) {
        this.id             = id
        this.userId         = userId
        this.parentId       = parentId
        this.title          = title
        this.tags           = tags
        this.note           = note
        this.completed      = completed
        this.streak         = streak
        this.createdAt      = createdAt
        this.updatedAt      = updatedAt
    }

    // Generate a new task from given api data
    static fromApi(u) {
        return new Task({
            id              : u["ID"],
            userId          : u["UserID"],
            parentId        : u["ParentID"],
            title           : u["Title"],
            tags            : u["Tags"],
            note            : u["Note"],
            completed       : u["Completed"],
            streak          : u["Streak"],
            createdAt       : u["CreatedAt"],
            updatedAt       : u["UpdatedAt"]
        })
    }

    // Get the completion status based on the data
    get isCompleted () {
        return !moment(this.completed).isBefore('1970-01-01', 'year')
    }
}