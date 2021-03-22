import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  _router: Router;
  newNotiCount: any = 0;
  noteCount: any = 0;
  tempCount: any = 0;
  tempReal: any;
  constructor(private router: Router) {
    this._router = this.router;
   }

  ngOnInit(): void {
  }


  hideBadge() {
    console.log("hide");
    this.noteCount = this.tempCount;
    this.newNotiCount = 0;
  }

  goHome() {
    localStorage.clear();
    console.log("clear data");
    localStorage.setItem("loggedIn", "0");
    this.router.navigate([""], { replaceUrl: true });
  }

}
