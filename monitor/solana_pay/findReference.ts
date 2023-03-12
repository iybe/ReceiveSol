import { ConfirmedSignatureInfo, Finality, SignaturesForAddressOptions, Connection, PublicKey } from '@solana/web3.js';

export class FindReferenceError extends Error {
    name = 'FindReferenceError';
}

export async function findReference(
    connection: Connection,
    reference: PublicKey,
    { finality, ...options }: SignaturesForAddressOptions & { finality?: Finality } = {}
): Promise<ConfirmedSignatureInfo> {
    const signatures = await connection.getSignaturesForAddress(reference, options, finality);

    const length = signatures.length;
    if (!length) throw new FindReferenceError('not found');

    // If one or more transaction signatures are found under the limit, return the oldest one.
    const oldest = signatures[length - 1];
    if (length < (options?.limit || 1000)) return oldest;

    try {
        // In the unlikely event that signatures up to the limit are found, recursively find the oldest one.
        return await findReference(connection, reference, { finality, ...options, before: oldest.signature });
    } catch (error: any) {
        // If the signatures found were exactly at the limit, there won't be more to find, so return the oldest one.
        if (error instanceof FindReferenceError) return oldest;
        throw error;
    }
}