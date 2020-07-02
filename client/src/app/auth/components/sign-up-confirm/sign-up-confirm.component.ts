import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { CoreStoreService } from 'src/app/core/services/core.store.service';

@Component({
  selector: 'app-sign-up-confirm',
  templateUrl: './sign-up-confirm.component.html',
  styleUrls: ['./sign-up-confirm.component.scss']
})
export class SignUpConfirmComponent implements OnInit, OnDestroy {
  public isOK = false;
  public isLoading = false;

  private destroy$ = new Subject<void>();

  constructor(
    private store: CoreStoreService,
    private activatedRoute: ActivatedRoute,
  ) { }

  public ngOnInit(): void {
    const token: string = this.activatedRoute.snapshot.params.token;
    if (!token) { this.isOK = false; return; }

    this.isOK = false;
    this.isLoading = true;

    this.store.signUpConfirm(token).pipe(
      takeUntil(this.destroy$)
    ).subscribe(() => {
      this.isLoading = false;
      this.isOK = true;
    }, () => {
      this.isLoading = false;
      this.isOK = false;
    });
  }

  public ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
