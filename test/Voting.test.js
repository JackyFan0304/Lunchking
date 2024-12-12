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

    it("should allow the owner to add candidates", async () => {
        await votingInstance.addCandidate("Wayne You 肯德雞", { from: accounts[0] });
        await votingInstance.addCandidate("Jacky Fan 麥當勞", { from: accounts[0] });
        await votingInstance.addCandidate("Cheng Tseng Pizza", { from: accounts[0] });
        await votingInstance.addCandidate("Nelson Tseng 茶餐廳", { from: accounts[0] });
        await votingInstance.addCandidate("Yoyo Wang LadyM 蛋糕", { from: accounts[0] });

        const candidate1 = await votingInstance.candidates(1);
        const candidate2 = await votingInstance.candidates(2);
        const candidate3 = await votingInstance.candidates(3);
        const candidate4 = await votingInstance.candidates(4);
        const candidate5 = await votingInstance.candidates(5);

        assert.equal(candidate1.name, "Wayne You 肯德雞", "First candidate name should be Wayne You 肯德雞");
        assert.equal(candidate2.name, "Jacky Fan 麥當勞", "Second candidate name should be Jacky Fan 麥當勞");
        assert.equal(candidate3.name, "Cheng Tseng Pizza", "Third candidate name should be Cheng Tseng Pizza");
        assert.equal(candidate4.name, "Nelson Tseng 茶餐廳", "Fourth candidate name should be Nelson Tseng 茶餐廳");
        assert.equal(candidate5.name, "Yoyo Wang LadyM 蛋糕", "Fifth candidate name should be Yoyo Wang LadyM 蛋糕");
    });

    it("should allow a user to vote for a candidate", async () => {
        await votingInstance.addCandidate("Wayne You 肯德雞", { from: accounts[0] });
        await votingInstance.vote(1, { from: accounts[1] }); // 帳戶 1 投票給候選人 ID 為 1

        const candidate = await votingInstance.candidates(1);
        assert.equal(candidate.voteCount.toNumber(), 1, "Vote count should be 1");
    });

    it("should prevent double voting by the same user", async () => {
        await votingInstance.addCandidate("Wayne You 肯德雞", { from: accounts[0] });
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
        await votingInstance.addCandidate("Wayne You 肯德雞", { from: accounts[0] });
        await votingInstance.addCandidate("Jacky Fan 麥當勞", { from: accounts[0] });
        await votingInstance.addCandidate("Cheng Tseng Pizza", { from: accounts[0] });
        await votingInstance.addCandidate("Nelson Tseng 茶餐廳", { from: accounts[0] });
        await votingInstance.addCandidate("Yoyo Wang LadyM 蛋糕", { from: accounts[0] });

        const candidates = await votingInstance.getAllCandidates();

        assert.equal(candidates.length, 5, "There should be 5 candidates");
        assert.equal(candidates[0].name, "Wayne You 肯德雞", "First candidate name should be Wayne You 肯德雞");
        assert.equal(candidates[1].name, "Jacky Fan 麥當勞", "Second candidate name should be Jacky Fan 麥當勞");
        assert.equal(candidates[2].name, "Cheng Tseng Pizza", "Third candidate name should be Cheng Tseng Pizza");
        assert.equal(candidates[3].name, "Nelson Tseng 茶餐廳", "Fourth candidate name should be Nelson Tseng 茶餐廳");
        assert.equal(candidates[4].name, "Yoyo Wang LadyM 蛋糕", "Fifth candidate name should be Yoyo Wang LadyM 蛋糕");
    });
});
