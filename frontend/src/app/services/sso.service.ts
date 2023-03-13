import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

interface ObjectDefined<V = any> {
  [x: string]: V;
}

export namespace SsoServiceInterface {
  export namespace Receive {
    export interface Login {
      token: string;
      id: string;
    }
  }

  export namespace Send {
    export interface Login {
      username: string;
      password: string;
    }

    export interface Register {
      username: string;
      password: string;
    }
  }
}

@Injectable({
  providedIn: 'root',
})
export class SsoService {
  public baseUrl = 'http://localhost:3002/user';

  constructor(private http: HttpClient) {}

  public login(
    params: SsoServiceInterface.Send.Login
  ): Observable<SsoServiceInterface.Receive.Login> {
    return this.http.post<SsoServiceInterface.Receive.Login>(
      `${this.baseUrl}/login`,
      params
    );
  }

  public register(
    params: SsoServiceInterface.Send.Register
  ): Observable<ObjectDefined> {
    return this.http.post<ObjectDefined>(`${this.baseUrl}`, params);
  }
}
