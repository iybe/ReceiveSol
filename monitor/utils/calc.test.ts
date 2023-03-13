import BigNumber from "bignumber.js";
import { calcAmount } from "./calc";

describe("calcAmount", () => {
    it("should calculate the correct amount with positive balance and fee", () => {
        let preBalance = new BigNumber(1.269995);
        let postBalance = new BigNumber(1.479995);

        const result = calcAmount(preBalance, postBalance);

        expect(result).toEqual(new BigNumber(0.21).toString());
    });
});
