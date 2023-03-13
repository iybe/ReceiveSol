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
  }
}

@Injectable({
  providedIn: 'root',
})
export class FunctionsService {
  constructor(private http: HttpClient) {}

  public userId = localStorage.getItem('id');

  public baseUrl = 'http://localhost:3002/account';

  public listAccount(): Observable<
    FunctionsServiceInterface.Receive.listAccount[]
  > {
    const header = new HttpHeaders({
      id: localStorage.getItem('id')!,
      token: localStorage.getItem('token')!,
    });

    return this.http.get<FunctionsServiceInterface.Receive.listAccount[]>(
      `${this.baseUrl}?userId=${this.userId}`,
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
      `${this.baseUrl}`,
      params,
      { headers: header }
    );
  }
}
