import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class FileViewService {

  url = 'https://api.rooster.jobs/core/public/jobs';

  constructor(private http: HttpClient, ) { }

  getFileView() {
    return this.http.get<any>(this.url);
  }
}
