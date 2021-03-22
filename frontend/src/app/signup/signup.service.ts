import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class SignupService {

  constructor(private http: HttpClient) { }

  signup(signupdata: any){
    return this.http.post("", signupdata)
   }
 
   register(regData: any){
     return this.http.post("http://localhost:8080/register", regData)
    }
}
