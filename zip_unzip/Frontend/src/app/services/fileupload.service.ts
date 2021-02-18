import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class FileuploadService {

  constructor(private http: HttpClient) { }

  fileUpload(file: File){
    const zipFile = new FormData; 

    zipFile.append('zip', file);
    return this.http.post("http://localhost:8080/fileupload",zipFile)
  }



}
