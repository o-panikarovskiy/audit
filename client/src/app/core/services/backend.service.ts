import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';



@Injectable()
export class BackendService {

  constructor(
    private http: HttpClient,
  ) { }

  public get(path: string, params?: any): Observable<any> {
    const url = this.getUrl(path);
    return this.http.get(url, { params });
  }

  public post(path: string, data?: any): Observable<any> {
    const url = this.getUrl(path);
    return this.http.post(url, data);
  }

  public put(path: string, data?: any): Observable<any> {
    const url = this.getUrl(path);
    return this.http.put(url, data);
  }

  public delete(path: string, params?: any): Observable<any> {
    const url = this.getUrl(path);
    return this.http.delete(url, { params });
  }

  private getUrl(path: string) {
    return `${environment.apiRoot}/${path}`;
  }
}
