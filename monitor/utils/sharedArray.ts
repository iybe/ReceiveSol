import AsyncLock from 'async-lock';

const lock = new AsyncLock();

export async function addToSharedArray(key: string, array: any, value: any) {
    await lock.acquire(key, async () => {
        array.push(value);
    });
}

export async function copySharedArray(key: string, array: any) {
    const arrayCopy = await lock.acquire(key, async () => {
        return [...array];
    });
    return arrayCopy
}

export async function removeFromSharedArray(key: string, array: any[], predicate: (value: any) => boolean) {
    await lock.acquire(key, async () => {
        for (let i = array.length - 1; i >= 0; i--) {
            if (predicate(array[i])) {
                array.splice(i, 1);
            }
        }
    });
}
