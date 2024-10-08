// This example shows that the due date of a task can be updated.
// If you do not specify a due date it will remain the same,
// and if you specify clickup.NullDate() it will be removed.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Zlatex/go-clickup/clickup"
)

func main() {
	var taskId string
	fmt.Print("Enter clickup taskId: ")
	fmt.Scanf("%s", &taskId)

	ctx := context.Background()
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	fmt.Println("\nGet due date of current task")
	getTask(ctx, client, taskId)

	fmt.Println("\nUpdate due date of the task to 2122/01/02 03:04:05:06")
	updateTask(ctx, client, taskId, &clickup.TaskUpdateRequest{
		DueDate: clickup.NewDate(
			time.Date(2122, 1, 2, 3, 4, 5, 6, time.Now().Location()),
		),
	})

	fmt.Println("\nUpdate the task with empty TaskRequest")
	updateTask(ctx, client, taskId, &clickup.TaskUpdateRequest{})

	fmt.Println("\nRemove task due date with NullDate()")
	updateTask(ctx, client, taskId, &clickup.TaskUpdateRequest{
		DueDate: clickup.NullDate(),
	})
}

func getTask(ctx context.Context, client *clickup.Client, taskID string) {
	task, _, err := client.Tasks.GetTask(ctx, taskID, &clickup.GetTaskOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(task.Name, task.DueDate)
}

func updateTask(ctx context.Context, client *clickup.Client, taskID string, tr *clickup.TaskUpdateRequest) {
	task, _, err := client.Tasks.UpdateTask(ctx, taskID, &clickup.GetTaskOptions{}, tr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(task.Name, task.DueDate)
}
