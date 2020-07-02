import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SignInComponent } from 'src/app/auth/components/sign-in/sign-in.component';
import { SignUpConfirmComponent } from 'src/app/auth/components/sign-up-confirm/sign-up-confirm.component';
import { SignUpComponent } from 'src/app/auth/components/sign-up/sign-up.component';


const routes: Routes = [
  {
    path: 'sign-in',
    component: SignInComponent
  },
  {
    path: 'sign-up',
    component: SignUpComponent
  },
  {
    path: 'confirm/:token',
    component: SignUpConfirmComponent
  },
  {
    path: '**',
    redirectTo: 'sign-in'
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AuthRoutingModule { }
