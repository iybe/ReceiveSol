import * as sharedArray from './utils/sharedArray'
import { validateTransfer } from './solana_pay/validateTransfer';
import { findReference, FindReferenceError } from './solana_pay/findReference'
import BigNumber from 'bignumber.js';
import { Connection, PublicKey, clusterApiUrl } from '@solana/web3.js';
import express from 'express';
import { MongoClient, ObjectId } from 'mongodb';
import dotenv from 'dotenv';
import { ValidateTransferError } from '@solana/pay';

const KEY_TESTNET_DATA = "KEY_TESTNET_DATA"
const KEY_TESTNET_DATA_VALIDATE = "KEY_TESTNET_DATA_VALIDATE"
const KEY_MAINNET_DATA = "KEY_MAINNET_DATA"
const KEY_MAINNET_DATA_VALIDATE = "KEY_MAINNET_DATA_VALIDATE"
const KEY_DEVNET_DATA = "KEY_DEVNET_DATA"
const KEY_DEVNET_DATA_VALIDATE = "KEY_DEVNET_DATA_VALIDATE"

interface Data {
    db_id: string;
    reference: string;
    recipient: string;
    amount: number;
    network: string
}

interface Err {
    error: string
}

interface StatusCheckTransaction {
    status: string
}

async function checker(array: Data[], connection: Connection, keyArray: string, keyArrayValidate: string) {
    const interval = setInterval(async () => {
        const now = new Date();
        const dateTimeString = now.toISOString().replace("T", " ").replace("Z", "");

        console.log(`${dateTimeString} : checker : ${keyArray} : ${array.length} links ativos`);
        let newArray = await sharedArray.copySharedArray(keyArray, array);
        let validateTransactions: Data[] = []

        const promises = newArray.map(async (element) => {
            try {
                const recipient = new PublicKey(element.recipient);
                const amount = new BigNumber(element.amount);
                const reference = new PublicKey(element.reference);

                let statusTransaction = await checkTransaction(connection, reference, amount, recipient)
                console.log(`[x] checker : statusTransaction = ${JSON.stringify(statusTransaction)}`);

                if (statusTransaction.status === "received_total" || statusTransaction.status === "received_incomplete") {
                    await Database.getInstance().updateStatus(element, statusTransaction.status);
                    await sharedArray.addToSharedArray(keyArrayValidate, validateTransactions, element)
                }
            } catch (error) {
                console.log(`[x] checker : error = ${error}`);
            }
        });

        await Promise.all(promises);

        const removePromises = validateTransactions.map(async (element) => {
            await sharedArray.removeFromSharedArray(keyArray, array, (value: any) => value === element);
        });

        await Promise.all(removePromises);
    }, 7000);
}

async function checkTransaction(connection: Connection, reference: PublicKey, amount: BigNumber, recipient: PublicKey) {
    try {
        let signatureInfo = await findReference(connection, reference, { finality: 'confirmed' });
        console.log(`[x] checkTransaction : signature found : reference = ${reference.toString()} : signature = ${signatureInfo.signature}`);
        try {
            let transactionResponse = await validateTransfer(connection, signatureInfo.signature, { recipient: recipient, amount });
            console.log(transactionResponse)
            console.log(`[x] checkTransaction : payment validate : reference = ${reference.toString()}`);
            let status: StatusCheckTransaction = { status: "received_total" }
            return status
        } catch (error: any) {
            if (error.message === "amount not transferred") {
                console.log(`[x] checkTransaction : amount not transferred : reference = ${reference.toString()} : error = ${error}`);
                let status: StatusCheckTransaction = { status: "received_incomplete" }
                return status
            }
            console.log(`[x] checkTransaction : payment validate failed : reference = ${reference.toString()} : error = ${error}`);
            let status: StatusCheckTransaction = { status: "failed" }
            return status
        }
    } catch (error: any) {
        if (!(error instanceof FindReferenceError)) {
            console.log(`[x] checkTransaction : error = ${error}`);
            let status: StatusCheckTransaction = { status: "error" }
            return status
        }
        let status: StatusCheckTransaction = { status: "not-found" }
        return status
    }
}

class Database {
    private static instance: Database;
    private client: MongoClient | undefined;
    private db: any;
    private collection: any;

    private constructor() { }

    public static async createInstance(uri: string, dbName: string, collectionName: string) {
        if (!Database.instance) {
            Database.instance = new Database();
            await Database.instance.connect(uri, dbName, collectionName);
        }
    }

    public static getInstance() {
        return Database.instance;
    }

    private async connect(uri: string, dbName: string, collectionName: string): Promise<void> {
        if (!this.client) {
            this.client = await MongoClient.connect(uri);
            this.db = this.client.db(dbName);
            this.collection = this.db.collection(collectionName);
        }
    }

    public async updateStatus(data: Data, status: string): Promise<void> {
        try {
            const query = { _id: new ObjectId(data.db_id) };
            const update = { $set: { status: status, amountReceived: data.amount } };

            await this.collection.updateOne(query, update);

            console.log(`[X] updateStatus : document ${data.db_id} updated`);
        } catch (error: any) {
            console.log(`[X] updateStatus : error = ${error} : data = ${JSON.stringify(data)} : status = ${status}`);
        }
    }

    public async close(): Promise<void> {
        if (this.client) {
            await this.client.close();
            console.log('[X] Database connection closed');
        }
    }
}

async function server(mainNetArray: Data[], testNetArray: Data[], devNetArray: Data[]) {
    const app = express();
    app.use(express.json());

    app.post('/link', (req, res) => {
        const data: Data = req.body;

        if (!data.db_id || !data.reference || !data.recipient || !data.amount || !data.network) {
            console.log("[X] dados invalidos: ", data);
            let err: Err = { error: "dados invalidos" }
            res.status(400).send(err);
            return
        }

        console.log("[X] mensagem recebida: ", data);

        switch (data.network) {
            case "mainnet":
                sharedArray.addToSharedArray(KEY_MAINNET_DATA, mainNetArray, data);
                break
            case "testnet":
                sharedArray.addToSharedArray(KEY_TESTNET_DATA, testNetArray, data);
                break
            case "devnet":
                sharedArray.addToSharedArray(KEY_DEVNET_DATA, devNetArray, data);
                break
            default:
                console.log("[X] network invalida: ", data.network);
                let err: Err = { error: "network invalida" }
                res.status(400).send(err);
                return
        }
        res.status(204).send();
    });

    app.listen(3000, () => {
        console.log('[X] Servidor iniciado na porta 3000');
    });
}

async function main() {
    dotenv.config();

    const mongodbUrl = process.env.MONGODB_URL;
    if (!mongodbUrl) {
        console.log(`[X] mongodbUrl not found`);
        return;
    }

    try {
        await Database.createInstance(mongodbUrl, 'mydb_backend', 'link');
        console.log('[X] Database connected');
    } catch (error: any) {
        console.log(`[X] Database connection error = ${error}`);
        return;
    }

    let linksMainNet: Data[] = []
    let linksTestNet: Data[] = []
    let linksDevNet: Data[] = []

    server(linksMainNet, linksTestNet, linksDevNet)

    console.log(`[*] conectando a rede testnet solana`);
    const connectionTestNet = new Connection(clusterApiUrl("testnet"), 'confirmed');
    console.log(`[x] conectado a rede testnet solana`);
    checker(linksTestNet, connectionTestNet, KEY_TESTNET_DATA, KEY_TESTNET_DATA_VALIDATE);

    console.log(`[*] conectando a rede mainnet solana`);
    const connectionMainNet = new Connection(clusterApiUrl("mainnet-beta"), 'confirmed');
    console.log(`[x] conectado a rede mainnet solana`);
    checker(linksMainNet, connectionMainNet, KEY_MAINNET_DATA, KEY_MAINNET_DATA_VALIDATE);

    console.log(`[*] conectando a rede devnet solana`);
    const connectionDevNet = new Connection(clusterApiUrl("devnet"), 'confirmed');
    console.log(`[x] conectado a rede devnet solana`);
    checker(linksDevNet, connectionDevNet, KEY_DEVNET_DATA, KEY_DEVNET_DATA_VALIDATE);
}

main()