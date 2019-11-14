import { kebabCase, replace, split } from 'lodash'

export class User {

    constructor ({ id, providerId, avatar, firstName, lastName, username, email, tasks, isAdmin, isRegistered, createdAt, updatedAt }) {
        this.id             = id
        this.providerId     = providerId
        this.avatar         = avatar
        this.firstName      = firstName
        this.lastName       = lastName
        this.username       = username
        this.email          = email
        this.tasks          = tasks || []
        this.isAdmin        = isAdmin
        this.isRegistered   = isRegistered
        this.createdAt      = createdAt
        this.updatedAt      = updatedAt
    }

    // Generate a new user from a given provider user
    static fromProvider({ id, name, picture, attrs }) {
        return new User({
            id              : null,
            providerId      : id,
            avatar          : picture,
            firstName       : split(name, ' ')[0],
            lastName        : split(name, ' ')[1],
            username        : replace(kebabCase(name), '-', ''),
            email           : '',
            tasks           : [],
            isAdmin         : attrs['admin'],
            isRegistered    : attrs['registered'],
            createdAt       : null,
            updatedAt       : null
        })
    }

    // Generate a new user from given api data
    static fromApi(u) {
        return new User({
            id              : u["ID"],
            providerId      : u["ProviderID"],
            avatar          : u["Avatar"],
            firstName       : u["FirstName"],
            lastName        : u["LastName"],
            username        : u["Username"],
            email           : u["Email"],
            tasks           : u["Tasks"],
            isAdmin         : u["IsAdmin"],
            isRegistered    : true,
            createdAt       : u["CreatedAt"],
            updatedAt       : u["UpdatedAt"]
        })
    }

    // Generate the full name based on the given names
    get fullName () {
        return this.firstName + ' ' + this.lastName
    }

    // Get the provider name from the provider id
    get provider () {
        return split(this.providerId, '_')[0]
    }
}