import * as moment from 'moment'

export class Task {
    constructor({ id, userId, title, tags, steps, note, completed, createdAt, updatedAt }) {
        this.id             = id
        this.userId         = userId
        this.title          = title
        this.tags           = tags
        this.steps          = steps
        this.note           = note
        this.completed      = completed
        this.createdAt      = createdAt
        this.updatedAt      = updatedAt
    }

    // Generate a new task from given api data
    static fromApi(u) {
        return new Task({
            id              : u["ID"],
            userId          : u["UserID"],
            title           : u["Title"],
            tags            : u["Tags"],
            steps           : u["Steps"],
            note            : u["Note"],
            completed       : u["Completed"],
            createdAt       : u["CreatedAt"],
            updatedAt       : u["UpdatedAt"]
        })
    }

    // Get the completion status based on the data
    set isCompleted (value) { this.completed = value ? moment().toISOString() : "0001-01-01T00:00:00Z" }
    get isCompleted () {
        return this.completed != null && !moment(this.completed).isBefore('1970-01-01', 'year')
    }
}