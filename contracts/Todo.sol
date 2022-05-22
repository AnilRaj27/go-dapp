//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

contract Todo {
    struct Task {
        string data;
        bool completed;
    }

    // create a new default task
    constructor() {
        addTask("complete assignment by 22nd May.!");
    }

    // dynamic array to store tasks
    Task[] tasks;

    // creates a task and adds to the list
    function addTask(string memory _data) public {
        Task memory newTask = Task(_data, false);
        tasks.push(newTask);
    }

    // gets a task provided a valid id
    function getTask(uint32 _id) public view returns (Task memory) {
        return tasks[_id];
    }

    // returns the entire list of tasks stored
    function getTaskList() public view returns (Task[] memory) {
        return tasks;
    }

    // toggles the task state provided an id of that task
    function toggleCompleted(uint32 _id) public {
        tasks[_id].completed = !tasks[_id].completed;
    }

    // updates the data of a particular task
    function updateTaskData(uint32 _id, string memory _data) public {
        tasks[_id].data = _data;
    }

    // removes the task from the task list
    function removeTask(uint32 _id) public {
        for (uint32 i = _id; i < tasks.length - 1; i++) {
            tasks[i] = tasks[i + 1];
        }
        tasks.pop();
    }
}
