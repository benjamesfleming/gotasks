package actions

import (
	"fmt"

	"github.com/benjamesfleming/gotasks/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// TasksResource is the resource for the Task model
type TasksResource struct {
	buffalo.Resource
}

/*
██╗     ██╗███████╗████████╗
██║     ██║██╔════╝╚══██╔══╝
██║     ██║███████╗   ██║
██║     ██║╚════██║   ██║
███████╗██║███████║   ██║
╚══════╝╚═╝╚══════╝   ╚═╝

*/

// List gets all Tasks.
// GET /api/tasks
func (v TasksResource) List(c buffalo.Context) error {
	// Grab the database connection from the current context
	// else return error and break
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Get all the tasks from the database
	// paginate with `page=?`, `per_page=?` params
	tasks := &models.Tasks{}
	err := tx.PaginateFromParams(c.Params()).All(tasks)
	if err != nil {
		return err
	}

	// Return the list of tasks
	return c.Render(200, r.JSON(tasks))
}

/*
███████╗██╗  ██╗ ██████╗ ██╗    ██╗
██╔════╝██║  ██║██╔═══██╗██║    ██║
███████╗███████║██║   ██║██║ █╗ ██║
╚════██║██╔══██║██║   ██║██║███╗██║
███████║██║  ██║╚██████╔╝╚███╔███╔╝
╚══════╝╚═╝  ╚═╝ ╚═════╝  ╚══╝╚══╝

*/

// Show gets the data for one Task.
// GET /api/tasks/{task_id}
func (v TasksResource) Show(c buffalo.Context) error {
	// Grab the database connection from the current context
	// else return error and break
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Attempt to find the task with the given task_id
	// else return 404 not found
	task := &models.Task{}
	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return c.Error(404, err)
	}

	// Return the requested task
	return c.Render(200, r.JSON(task))
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝

*/

// Create adds a Task to the DB.
// POST /api/tasks
func (v TasksResource) Create(c buffalo.Context) error {
	// Grab the database connection from the current context
	// else return error and break
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// // Bind the task to the context
	// if err := c.Bind(task); err != nil {
	// 	return err
	// }

	// Validate the created task
	// if there are any unknown errors then break
	task := &models.Task{}
	verrs, err := tx.ValidateAndCreate(task)
	if err != nil {
		return err
	}

	// Check for any validataion errors
	// if there are any return them in a 400 request
	if verrs.HasAny() {
		return c.Render(400, r.JSON(verrs))
	}

	// Return with the created task
	return c.Render(200, r.JSON(task))
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝

*/

// Update changes a Task in the DB
// PUT /api/tasks/{task_id}
func (v TasksResource) Update(c buffalo.Context) error {
	// Grab the database connection from the current context
	// else return error and break
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Attempt to find the task with the given task_id
	// else return 404 not found
	task := &models.Task{}
	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return c.Error(404, err)
	}

	// // Bind the task to the context
	// if err := c.Bind(task); err != nil {
	// 	return err
	// }

	// Validate the updated task
	// if there are any unknown errors then break
	verrs, err := tx.ValidateAndUpdate(task)
	if err != nil {
		return err
	}

	// Check for any validataion errors
	// if there are any return them in a 400 request
	if verrs.HasAny() {
		return c.Render(400, r.JSON(verrs))
	}

	// Return the updated task
	return c.Render(200, r.JSON(task))
}

/*
██████╗ ███████╗███████╗████████╗ ██████╗ ██████╗ ██╗   ██╗
██╔══██╗██╔════╝██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║  ██║█████╗  ███████╗   ██║   ██║   ██║██████╔╝ ╚████╔╝
██║  ██║██╔══╝  ╚════██║   ██║   ██║   ██║██╔══██╗  ╚██╔╝
██████╔╝███████╗███████║   ██║   ╚██████╔╝██║  ██║   ██║
╚═════╝ ╚══════╝╚══════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝   ╚═╝

*/

// Destroy deletes a Task from the DB.
// DELETE /api/tasks/{task_id}
func (v TasksResource) Destroy(c buffalo.Context) error {
	// Grab the database connection from the current context
	// else return error and break
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Attempt to find the task with the given task_id
	// else return 404 not found
	task := &models.Task{}
	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return c.Error(404, err)
	}

	// Delete the given task from the database
	// else return error and break
	if err := tx.Destroy(task); err != nil {
		return err
	}

	// Return 200 success
	return c.Render(200, r.JSON(task.ID))
}
