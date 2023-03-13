import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, tap } from 'rxjs';

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
  }

  export namespace Send {
    export interface listAccount {
      id: string;
      token: string;
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

  public listAccount(
    params: FunctionsServiceInterface.Send.listAccount
  ): Observable<FunctionsServiceInterface.Receive.listAccount[]> {
    const header = new HttpHeaders({
      id: params.id,
      token: params.token,
    });

    return this.http.get<FunctionsServiceInterface.Receive.listAccount[]>(
      `${this.baseUrl}?userId=${this.userId}`,
      { headers: header }
    );
  }
}
