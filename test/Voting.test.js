const Voting = artifacts.require("Voting"); // 引入 Voting 合約

contract("Voting", (accounts) => {
    let votingInstance;

    // 在每個測試之前部署合約
    beforeEach(async () => {
        votingInstance = await Voting.new(); // 部署新的合約實例
    });

    it("should set the owner correctly", async () => {
        const owner = await votingInstance.owner();
        assert.equal(owner, accounts[0], "Owner should be the deployer");
    });

    it("should allow the owner to add a candidate", async () => {
        await votingInstance.addCandidate("Alice", { from: accounts[0] }); // 添加候選人
        const candidate = await votingInstance.candidates(1); // 獲取候選人信息

        assert.equal(candidate.name, "Alice", "Candidate name should be Alice");
        assert.equal(candidate.voteCount.toNumber(), 0, "Initial vote count should be 0");
    });

    it("should not allow non-owner to add a candidate", async () => {
        try {
            await votingInstance.addCandidate("Bob", { from: accounts[1] }); // 非所有者嘗試添加候選人
            assert.fail("Non-owner was able to add a candidate");
        } catch (error) {
            assert(error.message.includes("Only the owner can perform this action."), "Error message should contain 'Only the owner can perform this action.'");
        }
    });

    it("should allow a user to vote for a candidate", async () => {
        await votingInstance.addCandidate("Alice", { from: accounts[0] });
        await votingInstance.vote(1, { from: accounts[1] }); // 帳戶 1 投票給候選人 ID 為 1

        const candidate = await votingInstance.candidates(1);
        assert.equal(candidate.voteCount.toNumber(), 1, "Vote count should be 1");
    });

    it("should prevent double voting by the same user", async () => {
        await votingInstance.addCandidate("Alice", { from: accounts[0] });
        await votingInstance.vote(1, { from: accounts[1] }); // 第一次投票

        try {
            await votingInstance.vote(1, { from: accounts[1] }); // 第二次投票（應該失敗）
            assert.fail("User was able to vote twice");
        } catch (error) {
            assert(error.message.includes("You have already voted."), "Error message should contain 'You have already voted.'");
        }
    });

    it("should reject votes for invalid candidate IDs", async () => {
        try {
            await votingInstance.vote(99, { from: accounts[1] }); // 投票給不存在的候選人 ID
            assert.fail("Vote for invalid candidate ID was accepted");
        } catch (error) {
            assert(error.message.includes("Invalid candidate"), "Error message should contain 'Invalid candidate'");
        }
    });

    it("should return all candidates correctly", async () => {
        await votingInstance.addCandidate("Alice", { from: accounts[0] });
        await votingInstance.addCandidate("Bob", { from: accounts[0] });

        const candidates = await votingInstance.getAllCandidates();

        assert.equal(candidates.length, 2, "There should be 2 candidates");
        assert.equal(candidates[0].name, "Alice", "First candidate name should be Alice");
        assert.equal(candidates[1].name, "Bob", "Second candidate name should be Bob");
    });
});