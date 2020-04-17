export function getRandomString(length: number) {
  return ab2str(crypto.getRandomValues(new Uint16Array(length)));
}

export function ab2str(buf): string {
  return String.fromCharCode.apply(null, new Uint16Array(buf));
}

export function str2ab(str) {
  const length = str.length;
  const buf = new ArrayBuffer(length * 2); // 2 bytes for each char
  const bufView = new Uint16Array(buf);
  for (let i = 0; i < length; i++) { bufView[i] = str.charCodeAt(i); }
  return buf;
}

export function createGuid() {
  const buf = new Uint16Array(8);
  crypto.getRandomValues(buf);
  return `${s4(buf[0])}${s4(buf[1])}-${s4(buf[2])}-${s4(buf[3])}-${s4(buf[4])}-${s4(buf[5])}${s4(buf[6])}${s4(buf[7])}`;

  function s4(num) {
    let ret = num.toString(16);
    while (ret.length < 4) { ret = '0' + ret; }
    return ret;
  }
}
