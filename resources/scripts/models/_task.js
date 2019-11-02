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
}