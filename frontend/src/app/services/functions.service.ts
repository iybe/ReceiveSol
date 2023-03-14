import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

interface ObjectDefined<V = any> {
  [x: string]: V;
}

export namespace FunctionsServiceInterface {
  export namespace Receive {
    export interface listAccount {
      ID: string;
      UserId: string;
      PublicKey: string;
      Nickname: string;
    }

    export interface createAccount {
      ID: string;
      UserId: string;
      PublicKey: string;
      Nickname: string;
    }

    export interface listLink {
      ID: string;
      Nickname: string;
      UserId: string;
      AccountId: string;
      Link: string;
      Reference: string;
      Recipient: string;
      Network: string;
      ExpectedAmount: number;
      AmountReceived: number;
      Status: string;
      CreatedAt: string;
      ReceivedAt: string;
      expand: boolean;
    }

    export interface createLink {
      link: string;
      reference: string;
    }
  }

  export namespace Send {
    export interface listAccount {
      id: string;
      token: string;
    }

    export interface createAccount {
      userId: string;
      publicKey: string;
      nickname: string;
    }

    export interface listLink {
      userId: string;
    }

    export interface createLink {
      userId: string;
      nickname: string;
      recipient: string;
      network: string;
      expectedAmount: number;
    }
  }
}

@Injectable({
  providedIn: 'root',
})
export class FunctionsService {
  constructor(private http: HttpClient) {}

  public userId = localStorage.getItem('id');

  public baseUrl = 'http://localhost:3002';

  public listAccount(): Observable<
    FunctionsServiceInterface.Receive.listAccount[]
  > {
    const header = new HttpHeaders({
      id: localStorage.getItem('id')!,
      token: localStorage.getItem('token')!,
    });

    return this.http.get<FunctionsServiceInterface.Receive.listAccount[]>(
      `${this.baseUrl}/account?userId=${this.userId}`,
      { headers: header }
    );
  }

  public createAccount(
    params: FunctionsServiceInterface.Send.createAccount
  ): Observable<FunctionsServiceInterface.Receive.createAccount> {
    const header = new HttpHeaders({
      id: localStorage.getItem('id')!,
      token: localStorage.getItem('token')!,
    });

    return this.http.post<FunctionsServiceInterface.Receive.createAccount>(
      `${this.baseUrl}/account`,
      params,
      { headers: header }
    );
  }

  public listLink(): Observable<FunctionsServiceInterface.Receive.listLink[]> {
    const header = new HttpHeaders({
      id: localStorage.getItem('id')!,
      token: localStorage.getItem('token')!,
    });

    return this.http.get<FunctionsServiceInterface.Receive.listLink[]>(
      `${this.baseUrl}/link?userId=${this.userId}`,
      { headers: header }
    );
  }

  public createLink(
    params: FunctionsServiceInterface.Send.createLink
  ): Observable<FunctionsServiceInterface.Receive.createLink> {
    const header = new HttpHeaders({
      id: localStorage.getItem('id')!,
      token: localStorage.getItem('token')!,
    });

    return this.http.post<FunctionsServiceInterface.Receive.createLink>(
      `${this.baseUrl}/link`,
      params,
      { headers: header }
    );
  }
}
