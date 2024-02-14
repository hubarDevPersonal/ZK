const fs = require('fs');

const UltraVerifier = artifacts.require("UltraVerifier");

describe("Correct data test:", function () {
    let verifier, receipt, proof, inputs, game, _return, winner;

    before(async function () {
        accounts = await web3.eth.getAccounts();

        proof = fs.readFileSync('./proofs/noirTicTacToe.proof', 'utf8')

        inputs = fs.readFileSync('./Verifier.toml', 'utf8').split('\n');
        game = inputs[0].slice(8, -1);
        winner = inputs[1].slice(10, -1);
    });

    describe("Verifier deployment:", function () {
        it("Deployed successfully", async function () {
          verifier = await UltraVerifier.new({ from: accounts[0] });
        });
    });

    describe("Proving a correct game result:", function () {
        it("Proved successfully", async function () {
            const publicInputs = [game, winner];
            const proofBytes = '0x' + proof

            receipt = await verifier.verify(proofBytes, publicInputs, { from: accounts[0] });

            assert.equal(true, receipt, "Successfully verified!");
        });
    });

    describe("Proving an incorrect game result:", function () {
        it("Proved unsuccessfully", async function () {
            const publicInputs = [winner, game];
            const proofBytes = '0x' + proof

            try {
                await verifier.verify(proofBytes, publicInputs, { from: accounts[0] });
                throw null;
            }
            catch (error) {
                assert(error, "Successfully failed!");
            }
        });
    });
});
