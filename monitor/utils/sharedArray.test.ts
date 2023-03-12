import { removeFromSharedArray, addToSharedArray, copySharedArray } from './sharedArray';

describe('removeFromSharedArray', () => {
    it('deve remover os elementos correspondentes do array', async () => {
        const array = [1, 2, 3, 4, 5];
        await removeFromSharedArray('key', array, (value) => value % 2 === 0);
        expect(array).toEqual([1, 3, 5]);
    });

    it('não deve afetar o array se nenhum elemento corresponder ao predicado', async () => {
        const array = [1, 2, 3, 4, 5];
        await removeFromSharedArray('key', array, (value) => value > 10);
        expect(array).toEqual([1, 2, 3, 4, 5]);
    });
});

describe('addToSharedArray', () => {
    test('should add value to array', async () => {
        let array: any = [];
        const value = 'test';
        await addToSharedArray('key', array, value);
        expect(array).toContain(value);
    });
});

describe('copySharedArray', () => {
    test('should return a copy of the array', async () => {
        const array = ['value1', 'value2'];
        const arrayCopy = await copySharedArray('key', array);
        expect(arrayCopy).toEqual(array);
        expect(arrayCopy).not.toBe(array);
    });
});