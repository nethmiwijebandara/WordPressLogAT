import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'

})
export class AuthService {

  constructor(private http: HttpClient) { }
  login(logindata: any){
    console.log("login")
    // const headers = new HttpHeaders()
    //   .append('Access-Control-Allow-Origin', '*');

    return this.http.post("http://localhost:8080/login", logindata)
   }
 
   register(regData: any){
     console.log(regData)
     return this.http.post("http://localhost:8080/register", regData)
    }
}
