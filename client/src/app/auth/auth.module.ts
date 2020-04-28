import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { AuthRoutingModule } from 'src/app/auth/auth-routing.module';
import { SignUpComponent } from 'src/app/auth/components/sign-up/sign-up.component';
import { SignInComponent } from './components/sign-in/sign-in.component';
import { SignUpConfirmComponent } from './components/sign-up-confirm/sign-up-confirm.component';

@NgModule({
  declarations: [
    SignInComponent,
    SignUpComponent,
    SignUpConfirmComponent,
  ],
  imports: [
    AuthRoutingModule,
    CommonModule,
    MatInputModule,
    MatButtonModule,
    ReactiveFormsModule,
  ]
})
export class AuthModule { }
