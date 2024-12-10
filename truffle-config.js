module.exports = {
    networks: {
        development: {
            host: "127.0.0.1", // 本地測試網地址
            port: 7545,        // Ganache 默認端口
            network_id: "*"    // 匹配任何網路 ID
        }
    },
    compilers: {
        solc: {
            version: "0.8.0" // 與 Voting.sol 的 Solidity 版本一致
        }
    }
};