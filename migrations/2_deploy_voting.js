const Voting = artifacts.require("Voting");

module.exports = async function (deployer) {
    deployer.deploy(Voting);
    const voting = await Voting.deployed();
    
    // 初始化候選人
    await voting.addCandidate("Costco大熱狗");
    await voting.addCandidate("全連便當");
};