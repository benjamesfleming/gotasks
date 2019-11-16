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
}