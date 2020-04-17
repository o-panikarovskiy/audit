import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { ISocketMessage, WebSocketService } from 'src/app/core/services/socket.service';

@Component({
  selector: 'app-socket',
  templateUrl: './socket.component.html',
  styleUrls: ['./socket.component.scss']
})
export class SocketComponent implements OnInit, OnDestroy {
  public messages: ISocketMessage[] = [];
  public readonly form: FormGroup;
  public readonly echoMsg: FormControl;

  private destroy$ = new Subject<void>();

  constructor(public ws: WebSocketService) {
    this.echoMsg = new FormControl('');

    this.form = new FormGroup({
      echoMsg: this.echoMsg
    });
  }

  public ngOnInit(): void {
    this.ws.connect('app');

    this.ws.appMessages$.pipe(
      takeUntil(this.destroy$)
    ).subscribe(msg => {
      this.messages = [...this.messages, msg];
    });
  }

  public ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }

  public onSubmit() {
    this.ws.send('app:prime', this.echoMsg.value);
  }

  public broadcast() {
    this.ws.send('app:prime:broadcast', this.echoMsg.value);
  }

  public excludeMe() {
    this.ws.send('app:prime:broadcast', this.echoMsg.value, true);
  }

  public disconnect() {
    this.ws.disconnect('app');
  }
}
