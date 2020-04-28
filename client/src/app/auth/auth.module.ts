import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { AuthRoutingModule } from 'src/app/auth/auth-routing.module';
import { SignInComponent } from './components/sign-in/sign-in.component';

@NgModule({
  declarations: [SignInComponent],
  imports: [
    AuthRoutingModule,
    CommonModule,
    MatInputModule,
    MatButtonModule,
    ReactiveFormsModule,
  ]
})
export class AuthModule { }
