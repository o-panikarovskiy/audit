import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { filter } from 'rxjs/operators';
import { environment } from 'src/environments/environment';

export enum SocketState { connected = 0, open = 1, closing = 2, closed = 3 }

export interface ISocketMessage {
  data?: any;
  eventName: string;
}

export interface OutputSocketMessage extends ISocketMessage {
  excludeMe?: boolean;
}

export interface InputSocketMessage extends ISocketMessage {
  clientId?: string;
  isMine?: boolean;
}

export interface ISocketOptions {
  pingInterval: number;
  reconnectInterval: number;
}

const DEFAULT_OPTIONS: ISocketOptions = {
  pingInterval: 3 * 1000,
  reconnectInterval: 3 * 1000,
};


@Injectable({ providedIn: 'root' })
export class WebSocketService {
  private clientId: string;
  private reconnectTimerId: number;

  private readonly isProduction: boolean;
  private readonly connectionUrl: string;
  private readonly options: ISocketOptions;
  private readonly sockets = new Map<string, WebSocket>();
  private readonly msgSub = new Subject<InputSocketMessage>();

  public readonly messages$ = this.msgSub.asObservable();
  public readonly appMessages$ = this.messages$.pipe(filter(m => !m.eventName.startsWith('socket:')));

  constructor() {
    this.isProduction = environment.production;
    this.options = { ...DEFAULT_OPTIONS, ...environment.socket };

    const protocol = window.location.protocol.replace('http', 'ws');
    this.connectionUrl = `${protocol}//${window.location.host}/${environment.apiRoot}/ws`;
  }

  public getState(socketId: string): SocketState {
    const socket = this.sockets.get(socketId);
    return socket !== void 0 ? socket.readyState : SocketState.closed;
  }

  public send(eventName: string, data: any, excludeMe = false) {
    const socketId = eventName.split(':')[0];
    const socket = this.sockets.get(socketId);

    if (socket !== void 0 && socket.readyState === SocketState.open) {
      const msg: OutputSocketMessage = { eventName, data, excludeMe };
      socket.send(JSON.stringify(msg));
    }

    return this;
  }

  public connect(socketId: string, protocols: string[] = []) {
    const socket = this.sockets.get(socketId);
    if (socket?.readyState === SocketState.open) {
      socket.close();
    }

    try {
      const ws = new WebSocket(this.connectionUrl, protocols);

      ws.onmessage = (e: MessageEvent) => this.onMessage(e);
      ws.onopen = () => this.onOpen();
      ws.onclose = (e: CloseEvent) => this.onClose(socketId, e, protocols);
      ws.onerror = (e: Event) => this.msgSub.error({ ws, event: e });

      this.sockets.set(socketId, ws);

      return true;
    } catch (error) {
      console.error(error);
      return false;
    }
  }

  public disconnect(socketId: string) {
    const socket = this.sockets.get(socketId);
    if (socket?.readyState === SocketState.open) {
      this.sockets.delete(socketId);
      socket.close();
    }
  }

  private onMessage(e: MessageEvent) {
    const message = JSON.parse(e.data) as InputSocketMessage;

    if (message.eventName === 'socket:client:id') {
      this.clientId = message.data;
    }

    message.isMine = message.clientId === this.clientId;

    this.notify(message);
  }

  private onOpen() {
    window.clearTimeout(this.reconnectTimerId);
    this.reconnectTimerId = 0;
    this.notify({ eventName: 'socket:open' });
  }

  private onClose(socketId: string, e: CloseEvent, protocols: string[] = []) {
    window.clearTimeout(this.reconnectTimerId);
    this.reconnectTimerId = 0;

    if (this.sockets.has(this.clientId)) {
      this.reconnectTimerId = window.setTimeout(() => {
        if (this.getState(socketId) === SocketState.closed) {
          this.connect(socketId, protocols);
        }
      }, this.options.reconnectInterval);
    }

    this.notify({ eventName: 'socket:close' });
  }

  private notify(message: ISocketMessage) {
    if (!this.isProduction) {
      // tslint:disable-next-line:no-console
      console.debug(message.eventName, message);
    }

    this.msgSub.next(message);
  }
}
