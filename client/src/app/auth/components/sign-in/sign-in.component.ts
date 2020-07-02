import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { CoreStoreService } from 'src/app/core/services/core.store.service';

@Component({
  selector: 'app-login',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.scss']
})
export class SignInComponent implements OnInit, OnDestroy {
  public readonly form = new FormGroup({
    username: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required)
  });

  public isSending = false;
  public isSubmitted = false;
  public serverError = '';

  private destroy$ = new Subject<void>();

  constructor(
    private router: Router,
    private store: CoreStoreService
  ) { }

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

    this.store.signIn(
      this.form.controls.username.value,
      this.form.controls.password.value
    ).pipe(
      takeUntil(this.destroy$)
    ).subscribe(() => {
      this.router.navigate(['/home']);
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
