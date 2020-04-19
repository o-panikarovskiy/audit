export const environment = {
  production: true,
  apiRoot: 'api',
  socket: {
    url: 'api/ws',
    pingInterval: 3 * 1000,
    reconnectInterval: 3 * 1000,
  }
};
