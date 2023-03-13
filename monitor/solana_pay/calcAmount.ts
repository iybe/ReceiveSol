import { Connection, TransactionSignature, Finality, Transaction } from '@solana/web3.js';
import { ValidateTransferFields, ValidateTransferError } from '@solana/pay';
import { validateSystemTransfer } from './validateTransfer'
import { calcAmount } from '../utils/calc'

export async function calcAmountPayment(
    connection: Connection,
    signature: TransactionSignature,
    { recipient, amount, reference }: ValidateTransferFields,
    options?: { commitment?: Finality }
) {
    const response = await connection.getTransaction(signature, options);
    if (!response) throw new ValidateTransferError('not found');

    const { message, signatures } = response.transaction;
    const meta = response.meta;
    if (!meta) throw new ValidateTransferError('missing meta');
    if (meta.err) throw meta.err;

    if (reference && !Array.isArray(reference)) {
        reference = [reference];
    }

    // Deserialize the transaction and make a copy of the instructions we're going to validate.
    const transaction = Transaction.populate(message, signatures);
    const instructions = transaction.instructions.slice();

    // Transfer instruction must be the last instruction
    const instruction = instructions.pop();
    if (!instruction) throw new ValidateTransferError('missing transfer instruction');
    const [preAmount, postAmount] = await validateSystemTransfer(instruction, message, meta, recipient, reference);
    console.log('preAmount', preAmount.toString());
    console.log('postAmount', postAmount.toString());
    return calcAmount(preAmount, postAmount);
}