import * as moment from 'moment'
import { Step } from '.' 

export class Task {
    constructor({ id, userId, title, tags, steps, note, order, completedAt, createdAt, updatedAt }) {
        this.id             = id
        this.userId         = userId
        this.title          = title
        this.tags           = tags
        this.steps          = steps || []
        this.note           = note
        this.order          = order
        this.completedAt    = completedAt
        this.createdAt      = createdAt
        this.updatedAt      = updatedAt
    }

    // Generate a new task from given api data
    static fromApi(t) {
        return new Task({
            id              : t["ID"],
            userId          : t["UserID"],
            title           : t["Title"],
            tags            : t["Tags"],
            steps           : Array.from(t["Steps"] || []).map(s => Step.fromApi(s)),
            note            : t["Note"],
            order           : t["Order"],
            completedAt     : t["CompletedAt"],
            createdAt       : t["CreatedAt"],
            updatedAt       : t["UpdatedAt"]
        })
    }

    // Get the completion status based on the data
    set isCompleted (value) { this.completedAt = value ? moment().toISOString() : "0001-01-01T00:00:00Z" }
    get isCompleted () {
        return this.completedAt != null && !moment(this.completedAt).isBefore('1970-01-01', 'year')
    }

}