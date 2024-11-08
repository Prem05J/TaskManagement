package Test

import (
	"github.com/taskManagement/Model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMockTasks() []Model.Task {
	return []Model.Task{
		{
			Id:          primitive.NewObjectID(),
			Title:       "Complete the proposal",
			Description: "Finalize the project proposal for the new feature",
			Status:      "In Progress",
			AssignedTo:  "Alice",
			CreatedAt:   "2024-10-15T10:30:00Z",
		},
		{
			Id:          primitive.NewObjectID(),
			Title:       "Fix the bugs in the login page",
			Description: "Resolve issues related to user login functionality",
			Status:      "Pending",
			AssignedTo:  "Bob",
			CreatedAt:   "2024-10-14T14:45:00Z",
		},
		{
			Id:          primitive.NewObjectID(),
			Title:       "Write the unit tests",
			Description: "Create unit tests for the authentication module",
			Status:      "Completed",
			AssignedTo:  "Charlie",
			CreatedAt:   "2024-10-10T09:00:00Z",
		},
		{
			Id:          primitive.NewObjectID(),
			Title:       "Deploy to production",
			Description: "Deploy the latest version of the app to the production server",
			Status:      "In Progress",
			AssignedTo:  "David",
			CreatedAt:   "2024-10-12T16:15:00Z",
		},
		{
			Id:          primitive.NewObjectID(),
			Title:       "Prepare the sprint report",
			Description: "Prepare and share the sprint review report with the team",
			Status:      "On Hold",
			AssignedTo:  "Eve",
			CreatedAt:   "2024-10-11T12:30:00Z",
		},
	}
}
