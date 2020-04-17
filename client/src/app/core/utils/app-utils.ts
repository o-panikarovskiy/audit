export function tryParseJson(str: string): any {
  try {
    return JSON.parse(str);
  } catch (error) {
    return str;
  }
}

