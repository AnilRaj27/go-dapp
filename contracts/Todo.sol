//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

contract Todo {
    struct Task {
        string data;
        bool completed;
    }

    constructor() {
        addTask("complete assignment by 22nd May.!");
    }

    Task[] tasks;

    function addTask(string memory _data) public {
        Task memory newTask = Task(_data, false);
        tasks.push(newTask);
    }

    function getTask(uint256 _id) public view returns (Task memory) {
        return tasks[_id];
    }

    function getTaskList() public view returns (Task[] memory) {
        return tasks;
    }

    function updateTask(uint256 _id, string memory _data) public {
        tasks[_id].data = _data;
    }

    function removeTask(uint256 _id) public {
        delete tasks[_id];
    }
}
