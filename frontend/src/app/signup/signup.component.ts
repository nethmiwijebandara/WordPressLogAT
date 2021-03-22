import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AuthService } from 'src/app/services/Auth/auth.service';


@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.sass']
})
export class SignupComponent implements OnInit {
  invalidSignup: boolean;
  constructor(private  authService: AuthService) { }

  form = new FormGroup({
    firstName: new FormControl('',[
      Validators.required
    ]),
    lastName: new FormControl('',[
      Validators.required
    ]),
    email: new FormControl('',[
      Validators.required,
      Validators.email
    ]),
    password: new FormControl('',[
      Validators.required
    ]),
    repassword: new FormControl('',[
      Validators.required,
  
    ])
  
  })


  ngOnInit(): void {
  }
  signup(data){
    console.log(data);
    this.authService.register(data).subscribe(result => {
      if (result) {
        console.log(result)
      
      } else this.invalidSignup = true;
    });
  }
}
