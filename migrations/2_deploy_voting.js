const Voting = artifacts.require("Voting");

module.exports = async function (deployer) {
    deployer.deploy(Voting);
    const voting = await Voting.deployed();
    
    // 初始化候選人
    await voting.addCandidate("Wayne You 肯德雞");
    await voting.addCandidate("Jacky Fan 麥當勞");
    await voting.addCandidate("Cheng Tseng Pizza");
    await voting.addCandidate("Nelson Tseng 茶餐廳");
    await voting.addCandidate("Yoyo Wang LadyM 蛋糕");
};