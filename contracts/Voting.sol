// SPDX-License-Identifier: MIT
// 指定合約的版權許可，MIT 是一種開放許可協議
pragma solidity ^0.8.0; // 指定 Solidity 編譯器版本，使用 0.8.x 系列

// 定義投票合約
contract Voting {
    address public owner; // 合約所有者的地址

    // 候選人結構體，用於存儲候選人的信息
    struct Candidate {
        uint id;          // 候選人 ID
        string name;      // 候選人姓名
        uint voteCount;   // 候選人獲得的票數
    }

    // 候選人映射，通過 ID 存儲候選人信息
    mapping(uint => Candidate) public candidates;
    // 投票者映射，用於記錄每個地址是否已經投票
    mapping(address => bool) public voters;
    uint public candidatesCount; // 候選人總數

    // 定義事件，用於記錄候選人添加和投票操作
    event CandidateAdded(uint id, string name); // 當候選人被添加時觸發
    event Voted(uint candidateId, address voter); // 當有人投票時觸發

    // 修飾符：限制只有合約所有者可以執行某些操作
    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can perform this action."); // 驗證調用者是否為所有者
        _; // 繼續執行修飾符所修飾的函數
    }

    // 構造函數：在合約部署時執行，設置合約所有者為部署者地址
    constructor() {
        owner = msg.sender; // 設置合約所有者為部署該合約的地址
    }

    // 添加候選人函數（僅限所有者調用）
    function addCandidate(string memory _name) public onlyOwner {
        candidatesCount++; // 增加候選人總數
        candidates[candidatesCount] = Candidate(candidatesCount, _name, 0); // 創建新候選人並存儲到映射中
        emit CandidateAdded(candidatesCount, _name); // 觸發 CandidateAdded 事件
    }

    // 投票函數（任何地址都可以調用，但每個地址只能投一次）
    function vote(uint _candidateId) public {
        require(!voters[msg.sender], "You have already voted."); // 確保當前地址尚未投票
        require(_candidateId > 0 && _candidateId <= candidatesCount, "Invalid candidate."); // 確保候選人 ID 有效

        voters[msg.sender] = true; // 記錄當前地址已經投票
        candidates[_candidateId].voteCount++; // 增加指定候選人的票數
        emit Voted(_candidateId, msg.sender); // 觸發 Voted 事件
    }

    // 獲取指定候選人的得票數（只讀函數，不消耗 Gas）
    function getVotes(uint _candidateId) public view returns (uint) {
        return candidates[_candidateId].voteCount; // 返回指定候選人的得票數
    }

    // 獲取所有候選人的信息（返回一個包含所有候選人的數組）
    function getAllCandidates() public view returns (Candidate[] memory) {
        Candidate[] memory allCandidates = new Candidate[](candidatesCount); // 創建動態數組來存儲所有候選人信息
        for (uint i = 1; i <= candidatesCount; i++) { 
            allCandidates[i - 1] = candidates[i]; // 將每個候選人的信息添加到數組中（索引從 0 開始）
        }
        return allCandidates; // 返回包含所有候選人的數組
    }
}