import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AuthService } from 'src/app/services/Auth/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.sass']
})
export class LoginComponent implements OnInit {
  invalidLogin: boolean;

    constructor(private  authService: AuthService ) { }

  form = new FormGroup({
   
    email: new FormControl('',[
      Validators.required,
      Validators.email
    ]),
    password: new FormControl('',[
      Validators.required
    ])
  })
  
  ngOnInit(): void {
  }

  signIn(credentials) {
    console.log(credentials);
    this.authService.login(credentials).subscribe(result => {
      if (result) {
        console.log(result)
        // let returnUrl = this.route.snapshot.queryParamMap.get("returnUrl");
        // this.router.navigate([returnUrl || "/admin"]);
      } else this.invalidLogin = true;
    });
  }

}
