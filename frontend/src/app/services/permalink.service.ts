import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

interface ObjectDefined<V = any> {
  [x: string]: V;
}

export namespace PermalinkServiceInterface {
  export namespace Send {
    export interface createPermalink {
      userId: string;
      expectedAmount: number;
    }

    export interface updatePermalink {
      userId: string;
      networkPermaLink: string;
      recipientPermaLink: string;
    }
  }

  export namespace Receive {
    export interface createPermalink {
      link: string;
      code: string;
    }

    export interface getPermalink {
      recipientPermaLink: string;
      networkPermaLink: string;
      url: string;
    }
  }
}

@Injectable({
  providedIn: 'root',
})
export class PermalinkService {
  constructor(private http: HttpClient) {}

  protected baseUrl = 'http://localhost:3002/user/permalink';

  protected baseUrlCreate = 'http://localhost:3002/permalink';

  public userId = localStorage.getItem('id');

  public createPermalink(
    params: PermalinkServiceInterface.Send.createPermalink
  ): Observable<PermalinkServiceInterface.Receive.createPermalink> {
    return this.http.post<PermalinkServiceInterface.Receive.createPermalink>(
      `${this.baseUrlCreate}`,
      params
    );
  }

  public getPermalink(): Observable<
    PermalinkServiceInterface.Receive.getPermalink
  > {
    const header = new HttpHeaders({
      id: localStorage.getItem('id')!,
      token: localStorage.getItem('token')!,
    });

    return this.http.get<PermalinkServiceInterface.Receive.getPermalink>(
      `${this.baseUrl}?userId=${this.userId}`,
      { headers: header }
    );
  }

  public updatePermalink(
    params: PermalinkServiceInterface.Send.updatePermalink
  ): Observable<ObjectDefined> {
    const header = new HttpHeaders({
      id: localStorage.getItem('id')!,
      token: localStorage.getItem('token')!,
    });

    return this.http.post<ObjectDefined>(`${this.baseUrl}`, params, {
      headers: header,
    });
  }
}
