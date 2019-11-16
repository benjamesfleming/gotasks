import * as moment from 'moment'

export class Step {
    constructor({ id, userId, taskId, title, order, completedAt, createdAt, updatedAt }) {
        this.id             = id
        this.userId         = userId
        this.taskId         = taskId
        this.title          = title
        this.order          = order
        this.completedAt    = completedAt
        this.createdAt      = createdAt
        this.updatedAt      = updatedAt
    }

    // Generate a new step from given api data
    static fromApi(s) {
        return new Step({
            id              : s["ID"],
            userId          : s["UserID"],
            taskId          : s["TaskID"],
            title           : s["Title"],
            order           : s["Order"],
            completedAt     : s["CompletedAt"],
            createdAt       : s["CreatedAt"],
            updatedAt       : s["UpdatedAt"]
        })
    }

    // Get the completion status based on the data
    set isCompleted (value) { this.completedAt = value ? moment().toISOString() : "0001-01-01T00:00:00Z" }
    get isCompleted () {
        return this.completedAt != null && !moment(this.completedAt).isBefore('1970-01-01', 'year')
    }
}