// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract TransactionStorage {
    struct Transaction {
        uint id;
        string from;
        string to;
        uint amount;
        uint timestamp;
    }

    string public version;

    constructor() {
        version = "1";
    }

    Transaction[] private transactions;
    uint private nextId;

    event TransactionAdded(
        uint id,
        string from,
        string to,
        uint amount,
        uint timestamp
    );

    function addTransaction(
        string memory _from,
        string memory _to,
        uint _amount
    ) public {
        transactions.push(
            Transaction({
                id: nextId,
                from: _from,
                to: _to,
                amount: _amount,
                timestamp: block.timestamp
            })
        );

        emit TransactionAdded(nextId, _from, _to, _amount, block.timestamp);
        nextId++;
    }

    function getVersion() public view returns (string memory) {
        return version;
    }

    function getTransactions() public view returns (Transaction[] memory) {
        return transactions;
    }

    function getTransactionDetail(
        uint _id
    ) public view returns (Transaction memory) {
        require(_id < nextId, "Transaction ID does not exist.");
        return transactions[_id];
    }
}
