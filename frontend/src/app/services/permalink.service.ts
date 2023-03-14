import { HttpClient } from '@angular/common/http';
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
  }

  export namespace Receive {
    export interface createPermalink {
      link: string;
      code: string;
    }
  }
}

@Injectable({
  providedIn: 'root',
})
export class PermalinkService {
  constructor(private http: HttpClient) {}

  protected baseUrl = 'http://localhost:3002/permalink';

  public createPermalink(
    params: PermalinkServiceInterface.Send.createPermalink
  ): Observable<PermalinkServiceInterface.Receive.createPermalink> {
    return this.http.post<PermalinkServiceInterface.Receive.createPermalink>(
      `${this.baseUrl}`,
      params
    );
  }
}
