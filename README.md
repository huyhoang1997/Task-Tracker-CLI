# Task-Tracker-CLI
Its just a task tracker app

1. run go build to build the project to an executable file
Available commands:
# Adding a new task
Task-Tracker-CLI add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
Task-Tracker-CLI update 1 "Buy groceries and cook dinner"
Task-Tracker-CLI delete 1

# Marking a task as in progress or done
Task-Tracker-CLI mark-in-progress 1
Task-Tracker-CLI mark-done 1

# Listing all tasks
Task-Tracker-CLI list

# Listing tasks by status
Task-Tracker-CLI list done
Task-Tracker-CLI list todo
Task-Tracker-CLI list in-progress
