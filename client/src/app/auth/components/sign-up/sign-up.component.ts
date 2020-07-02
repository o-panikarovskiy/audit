import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { getPasswordValidator } from 'src/app/auth/components/sign-up/verify-password';
import { CoreStoreService } from 'src/app/core/services/core.store.service';

@Component({
  selector: 'app-login',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.scss']
})
export class SignUpComponent implements OnInit, OnDestroy {
  public readonly form = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required, Validators.minLength(8)]),
    passwordVerify: new FormControl('', [Validators.required, Validators.minLength(8)]),
  }, { validators: getPasswordValidator('password', 'passwordVerify') });

  public isSending = false;
  public isSuccess = false;
  public isSubmitted = false;
  public serverError = '';

  private destroy$ = new Subject<void>();

  constructor(private store: CoreStoreService) { }

  public ngOnInit(): void {
    this.form.valueChanges.pipe(
      takeUntil(this.destroy$)
    ).subscribe(() => {
      this.serverError = '';
    });
  }

  public ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  public onSubmit() {
    this.isSubmitted = true;
    if (!this.form.valid) { return; }

    this.form.disable();
    this.serverError = '';
    this.isSending = true;

    this.store.signUp(
      this.form.controls.email.value,
      this.form.controls.password.value
    ).pipe(
      takeUntil(this.destroy$)
    ).subscribe(() => {
      this.reset();
      this.isSuccess = true;
    }, (err: HttpErrorResponse) => {
      this.reset();
      this.serverError = err?.error?.message;
    });
  }

  private reset() {
    this.form.enable();
    this.serverError = '';
    this.isSending = false;
  }
}
