import { Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { CoreStoreService } from 'src/app/core/services/core.store.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit, OnDestroy {
  private destroy$ = new Subject<void>();
  private isSending = false;

  constructor(
    private router: Router,
    public store: CoreStoreService,
  ) { }

  public ngOnInit(): void {

  }

  public ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  public singOut(e: MouseEvent) {
    e.preventDefault();
    if (this.isSending) { return; }

    this.isSending = true;
    this.store.singOut().pipe(
      takeUntil(this.destroy$)
    ).subscribe(() => {
      this.router.navigate(['auth', 'sign-in']);
    });
  }
}
