import BigNumber from "bignumber.js";

export function calcAmount(preBalance: BigNumber, postBalance: BigNumber) {
    let postBalance2 = postBalance.minus(preBalance);
    return postBalance2.toFixed(2);
}