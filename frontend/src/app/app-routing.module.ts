import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AdvancedComponent } from './dashboard/advanced/advanced.component';
import { DashboardComponent } from './dashboard/dashboard.component';
// import { FileuploadComponent } from './dashboard/fileupload/fileupload.component.html';
// import { FileviewComponent } from './dashboard/fileview/fileview.component';
import { LandingComponent } from './landing/landing.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';

const routes: Routes = [
  { path: '', component: LandingComponent },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
  {
    path: 'dashboard', component: DashboardComponent,
    children: [
      // {
      // path:"fileupload",
      // component:FileuploadComponent
      // },

      // {
      // path:"fileview",
      // component: FileviewComponent
      // },
      {
        path: 'advanced',
        component: AdvancedComponent
      },
      {
        path: '',
        redirectTo: 'fileupload',
        pathMatch: 'full',
      },
    ]

  }

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
